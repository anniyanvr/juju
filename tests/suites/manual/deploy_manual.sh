test_deploy_manual() {
	if [ "$(skip 'test_deploy_manual')" ]; then
		echo "==> TEST SKIPPED: deploy manual"
		return
	fi

	(
		set_verbosity

		cd .. || exit

		case "${BOOTSTRAP_PROVIDER:-}" in
		"lxd" | "lxd-remote" | "localhost")
			export BOOTSTRAP_PROVIDER="manual"
			run "run_deploy_manual_lxd"
			;;
		"aws" | "ec2")
			export BOOTSTRAP_PROVIDER="manual"
			run "run_deploy_manual_aws"
			;;
		*)
			echo "==> TEST SKIPPED: deploy manual - tests for LXD and AWS only"
			;;
		esac
	)
}

manual_deploy() {
	local cloud_name name addr_m1 addr_m2

	cloud_name=${1}
	name=${2}
	addr_m1=${3}
	addr_m2=${4}

	juju add-cloud --client "${cloud_name}" "${TEST_DIR}/cloud_name.yaml" >"${TEST_DIR}/add-cloud.log" 2>&1

	file="${TEST_DIR}/test-${name}.log"

	bootstrap "${cloud_name}" "test-${name}" "${file}"

	juju add-machine ssh:ubuntu@"${addr_m1}" >"${TEST_DIR}/add-machine-1.log" 2>&1
	juju add-machine ssh:ubuntu@"${addr_m2}" >"${TEST_DIR}/add-machine-2.log" 2>&1

	juju enable-ha >"${TEST_DIR}/enable-ha.log" 2>&1

	machine_series=$(juju machines --format=json | jq -r '.machines | .["0"] | .series')

	juju deploy ubuntu --to=0 --series="${machine_series}"

	wait_for "ubuntu" "$(idle_condition "ubuntu" 0 0)"

	juju remove-application ubuntu

	destroy_controller "test-${name}"

	delete_user_profile "${name}"
}
