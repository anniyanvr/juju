// Copyright 2019 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package caas

import (
	"encoding/hex"
	"math/rand"

	"github.com/juju/errors"

	"github.com/juju/juju/caas/kubernetes/clientconfig"
	"github.com/juju/juju/caas/kubernetes/provider"
	jujucloud "github.com/juju/juju/cloud"
	"github.com/juju/juju/environs"
)

const rbacLabelKeyName = provider.RBACLabelKeyName

func ensureCredentialUID(
	credentialName, credentialUID string,
	credential jujucloud.Credential,
) (cred jujucloud.Credential, _ error) {

	newAttr := credential.Attributes()
	if newAttr == nil {
		return cred, errors.NotValidf("empty credential %q", credentialName)
	}
	newAttr[rbacLabelKeyName] = credentialUID
	return jujucloud.NewNamedCredential(
		credentialName, credential.AuthType(), newAttr, credential.Revoked,
	), nil
}

type credentialGetter interface {
	// CredentialForCloud gets credentials for the named cloud.
	CredentialForCloud(string) (*jujucloud.CloudCredential, error)
}

func getExistingCredential(store credentialGetter, cloudName, credentialName string) (credential jujucloud.Credential, err error) {
	cloudCredential, err := store.CredentialForCloud(cloudName)
	if err != nil {
		return credential, errors.Trace(err)
	}
	var ok bool
	if credential, ok = cloudCredential.AuthCredentials[credentialName]; !ok {
		return credential, errors.NotFoundf("credential %q for cloud %q", credentialName, cloudName)
	}
	return credential, nil
}

func decideCredentialUID(store credentialGetter, cloudName, credentialName string) (string, error) {
	var credUID string
	existingCredential, err := getExistingCredential(store, cloudName, credentialName)
	if err != nil && !errors.IsNotFound(err) {
		return "", errors.Trace(err)
	}
	if err == nil && existingCredential.Attributes() != nil {
		credUID = existingCredential.Attributes()[rbacLabelKeyName]
	}

	if credUID == "" {
		b := make([]byte, 4)
		if _, err := rand.Read(b); err != nil {
			return credUID, errors.Trace(err)
		}
		credUID = hex.EncodeToString(b)
	}
	return credUID, nil
}

func cleanUpCredentialRBAC(cloud jujucloud.Cloud, credential jujucloud.Credential) error {
	attr := credential.Attributes()
	if attr == nil {
		return nil
	}
	credUID := attr[rbacLabelKeyName]
	if credUID == "" {
		return nil
	}

	cloudSpec, err := environs.MakeCloudSpec(cloud, "", &credential)
	if err != nil {
		return errors.Trace(err)
	}
	restConfig, err := provider.CloudSpecToK8sRestConfig(cloudSpec)
	if err != nil {
		return errors.Trace(err)
	}
	err = clientconfig.RemoveCredentialRBACResources(restConfig, credUID)
	return errors.Trace(err)
}
