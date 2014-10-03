#!/usr/bin/env bash
# Release public tools.
#
# Publish to Canonistack, HP, AWS, and Azure.
# This script requires that the user has credentials to upload the tools
# to Canonistack, HP Cloud, AWS, and Azure

set -e

SCRIPT_DIR=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd )

AWS_SITE="http://juju-dist.s3.amazonaws.com"
HP_SITE="https://region-a.geo-1.objects.hpcloudsvc.com/v1/60502529753910/juju-dist"
CAN_SITE="https://swift.canonistack.canonical.com/v1/AUTH_526ad877f3e3464589dc1145dfeaac60/juju-dist"
AZURE_SITE="https://jujutools.blob.core.windows.net/juju-tools"
JOYENT_SITE="https://us-east.manta.joyent.com/cpcjoyentsupport/public/juju-dist"


usage() {
    echo "usage: $0 PURPOSE DIST_DIRECTORY DESTINATIONS"
    echo "  PURPOSE: 'release', 'proposed', 'devel', or  'testing'"
    echo "    release installs tools/ at the top of juju-dist/tools."
    echo "    proposed installs tools/ at the top of juju-dist/proposed/tools."
    echo "    devel installs tools/ at the top of juju-dist/devel/tools."
    echo "    testing installs tools/ at juju-dist/testing/tools."
    echo "  DIST_DIRECTORY: The directory to the assembled tools."
    echo "    This is the juju-dist dir created by assemble-public-tools.bash."
    echo "  DESTINATIONS: cpc or streams"
    echo "    cpc publishes tools to the certified public clouds."
    echo "    streams publishes tools just to streams.canonical.com."
    exit 1
}


verify_stream() {
    [[ -z "$DRY_RUN" ]] || return 0
    local location="$1"
    if [[ $PURPOSE == "release" ]]; then
        local root="tools"
    else
        local root="$PURPOSE/tools"
    fi
    echo "Verifying the streams at $location/$root"
    echo "are public and are identical to the source"
    curl -s $location/$root/streams/v1/index.json > $WORK/index.json
    diff $STREAM_PATH/streams/v1/index.json $WORK/index.json
    curl -s $location/$root/streams/v1/index.json > \
        $WORK/com.ubuntu.juju:released:tools.json
    diff $STREAM_PATH/streams/v1/index.json \
        $WORK/com.ubuntu.juju:released:tools.json
    rm $WORK/*
}


check_deps() {
    echo "Phase 0: Checking requirements."
    has_deps=1
    which swift || has_deps=0
    which s3cmd || has_deps=0
    test -f $JUJU_DIR/canonistacktoolsrc || has_deps=0
    test -f $JUJU_DIR/hptoolsrc || has_deps=0
    test -f $JUJU_DIR/s3cfg || has_deps=0
    test -f $JUJU_DIR/azuretoolsrc || has_deps=0
    if [[ $has_deps == 0 ]]; then
        echo "Install python-swiftclient, and s3cmd"
        echo "Your $JUJU_DIR dir must contain rc files to publish:"
        echo "  canonistacktoolsrc, hptoolsrc, s3cfg, azuretoolsrc"
        exit 2
    fi
}


publish_to_aws() {
    [[ $DESTINATIONS == 'cpc' ]] || return 0
    if [[ $PURPOSE == "release" ]]; then
        local destination="s3://juju-dist/"
    else
        local destination="s3://juju-dist/$PURPOSE/"
    fi
    echo "Phase 1: Publishing $PURPOSE to AWS."
    s3cmd -c $JUJU_DIR/s3cfg $DRY_RUN sync --exclude '*mirror*' \
        $STREAM_PATH $destination
    verify_stream $AWS_SITE
}


publish_to_canonistack() {
    [[ $DESTINATIONS == 'cpc' ]] || return 0
    [[ "${IGNORE_CANONISTACK-}" == 'true' ]] && return 0
    if [[ $PURPOSE == "release" ]]; then
        local destination="tools"
    else
        local destination="$PURPOSE/tools"
    fi
    echo "Phase 2: Publishing $PURPOSE to canonistack."
    source $JUJU_DIR/canonistacktoolsrc
    cd $STREAM_PATH/releases/
    ${SCRIPT_DIR}/swift_sync.py $DRY_RUN $destination/releases/ *.tgz
    cd $STREAM_PATH/streams/v1
    ${SCRIPT_DIR}/swift_sync.py $DRY_RUN $destination/streams/v1/ {index,com}*
    verify_stream $CAN_SITE
}


publish_to_hp() {
    [[ $DESTINATIONS == 'cpc' ]] || return 0
    if [[ $PURPOSE == "release" ]]; then
        local destination="tools"
    else
        local destination="$PURPOSE/tools"
    fi
    echo "Phase 3: Publishing $PURPOSE to HP Cloud."
    source $JUJU_DIR/hptoolsrc
    cd $STREAM_PATH/releases/
    ${SCRIPT_DIR}/swift_sync.py $DRY_RUN $destination/releases/ *.tgz
    cd $STREAM_PATH/streams/v1
    ${SCRIPT_DIR}/swift_sync.py $DRY_RUN $destination/streams/v1/ {index,com}*
    verify_stream $HP_SITE
}


publish_to_azure() {
    [[ $DESTINATIONS == 'cpc' ]] || return 0
    echo "Phase 4: Publishing $PURPOSE to Azure."
    source $JUJU_DIR/azuretoolsrc
    ${SCRIPT_DIR}/azure_publish_tools.py $DRY_RUN publish $PURPOSE $JUJU_DIST
    verify_stream $AZURE_SITE
}


publish_to_joyent() {
    [[ $DESTINATIONS == 'cpc' ]] || return 0
    [[ "${IGNORE_JOYENT-}" == 'true' ]] && return 0
    if [[ $PURPOSE == "release" ]]; then
        local destination="tools"
    else
        local destination="$PURPOSE/tools"
    fi
    echo "Phase 5: Publishing $PURPOSE to Joyent."
    source $JUJU_DIR/joyentrc
    cd $STREAM_PATH/releases/
    ${SCRIPT_DIR}/manta_sync.py $DRY_RUN $destination/releases/ *.tgz
    cd $STREAM_PATH/streams/v1
    ${SCRIPT_DIR}/manta_sync.py $DRY_RUN $destination/streams/v1/ {index,com}*
    verify_stream $JOYENT_SITE
}


publish_to_streams() {
    [[ $DESTINATIONS == 'streams' ]] ||  return 0
    echo "Phase 6: Publishing $PURPOSE to streams.canonical.com."
    source $JUJU_DIR/streamsrc
    destination=$STREAMS_OFFICIAL_DEST
    rsync $DRY_RUN -avzh $JUJU_DIST/ $destination
}


# The location of environments.yaml and rc files.
JUJU_DIR=${JUJU_HOME:-$HOME/.juju}

DRY_RUN=""
if  [[ "$1" == "--dry-run" ]]; then
    DRY_RUN="--dry-run"
    echo "No changes will be made."
    shift
fi

test $# -eq 3 || usage

PURPOSE=$1
if [[ ! $PURPOSE =~ ^(release|proposed|devel|testing)$ ]]; then
    echo "Invalid PURPOSE."
    usage
fi

JUJU_DIST=$(cd $2; pwd)
if [[ $PURPOSE == "release" ]]; then
    STREAM_PATH="$JUJU_DIST/tools"
else
    STREAM_PATH="$JUJU_DIST/$PURPOSE/tools"
fi
if [[ ! -d $STREAM_PATH/releases && ! -d $STREAM_PATH/streams ]]; then
    echo "Invalid JUJU-DIST: $STREAM_PATH"
    usage
fi

DESTINATIONS=$3
if [[ $DESTINATIONS != "cpc" && $DESTINATIONS != "streams" ]]; then
    echo "Invalid DESTINATIONS."
    usage
fi


check_deps
WORK=$(mktemp -d)
publish_to_aws
publish_to_canonistack
publish_to_hp
publish_to_azure
publish_to_joyent
publish_to_streams
rm -r $WORK
echo "Published $PURPOSE data to all CPCs."
