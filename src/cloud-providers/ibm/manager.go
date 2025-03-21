// (C) Copyright Confidential Containers Contributors
// SPDX-License-Identifier: Apache-2.0

package ibm

import (
	"flag"

	provider "github.com/confidential-containers/cloud-api-adaptor/src/cloud-providers"
)

var ibmConfig Config

type Manager struct{}

func init() {
	provider.AddCloudProvider("ibm", &Manager{})
}

func (m *Manager) ParseCmd(flags *flag.FlagSet) {
	provider.Get("ibmcloud").ParseCmd(flags)
	provider.Get("ibmcloud-powervs").ParseCmd(flags)
}

func (m *Manager) LoadEnv() {
	provider.Get("ibmcloud").LoadEnv()
	provider.Get("ibmcloud-powervs").LoadEnv()
}

func (m *Manager) NewProvider() (provider.Provider, error) {
	return NewProvider(&ibmConfig)
}

func (m *Manager) GetConfig() (config *Config) {
	return &ibmConfig
}
