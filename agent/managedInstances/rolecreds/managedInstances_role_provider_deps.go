// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the AWS Customer Agreement (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/agreement/

// package rolecreds contains functions that help procure the managed instance auth credentials
// dependencies
package rolecreds

import (
	"github.com/aws/amazon-ssm-agent/agent/managedInstances/registration"
)

// dependency for managed instance registration
var managedInstance instanceRegistration = instanceInfo{}

type instanceRegistration interface {
	InstanceID() string
	Region() string
	PrivateKey() string
	Fingerprint() (string, error)
	GenerateKeyPair() (string, string, string, error)
	UpdatePrivateKey(string, string) error
}

type instanceInfo struct{}

// ServerID returns the managed instance ID
func (instanceInfo) InstanceID() string { return registration.InstanceID() }

// Region returns the managed instance region
func (instanceInfo) Region() string { return registration.Region() }

// PrivateKey returns the managed instance PrivateKey
func (instanceInfo) PrivateKey() string { return registration.PrivateKey() }

// Fingerprint returns the managed instance fingerprint
func (instanceInfo) Fingerprint() (string, error) { return registration.Fingerprint() }

// GenerateKeyPair generate a new keypair
func (instanceInfo) GenerateKeyPair() (publicKey, privateKey, keyType string, err error) {
	return registration.GenerateKeyPair()
}

// UpdatePrivateKey saves the private key into the registration persistance store
func (instanceInfo) UpdatePrivateKey(privateKey, privateKeyType string) (err error) {
	return registration.UpdatePrivateKey(privateKey, privateKeyType)
}
