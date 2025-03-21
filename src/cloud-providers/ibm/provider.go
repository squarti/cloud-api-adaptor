// (C) Copyright Confidential Containers Contributors
// SPDX-License-Identifier: Apache-2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"strings"

	provider "github.com/confidential-containers/cloud-api-adaptor/src/cloud-providers"
	"github.com/confidential-containers/cloud-api-adaptor/src/cloud-providers/util/cloudinit"
)

var logger = log.New(log.Writer(), "[adaptor/cloud/ibm] ", log.LstdFlags|log.Lmsgprefix)

var POWERVS_MACHINE_TYPES = []string{"s922", "s1022", "e1080", "e980"}

type ibmProvider struct {
	vpc     provider.Provider
	powervs provider.Provider
}

func NewProvider(config *Config) (provider.Provider, error) {
	vpc, err := provider.Get("ibmcloud").NewProvider()
	if err != nil {
		return nil, err
	}

	powervs, err := provider.Get("ibmcloud-powervs").NewProvider()
	if err != nil {
		return nil, err
	}

	provider := &ibmProvider{
		vpc,
		powervs,
	}

	logger.Printf("ibm config: %#v", config.Redact())

	return provider, nil
}

func (p *ibmProvider) CreateInstance(ctx context.Context, podName, sandboxID string, cloudConfig cloudinit.CloudConfigGenerator, spec provider.InstanceTypeSpec) (*provider.Instance, error) {
	for _, machineType := range POWERVS_MACHINE_TYPES {
		if strings.HasPrefix(spec.InstanceType, machineType) {
			instance, err := p.powervs.CreateInstance(ctx, podName, sandboxID, cloudConfig, spec)
			return p.wrapID("ibmcloud-powervs", instance), err
		}
	}

	instance, err := p.vpc.CreateInstance(ctx, podName, sandboxID, cloudConfig, spec)
	return p.wrapID("ibmcloud", instance), err
}

func (p *ibmProvider) DeleteInstance(ctx context.Context, instanceID string) error {
	provider, id := p.unwrapID(instanceID)
	if provider == "ibmcloud-powervs" {
		return p.powervs.DeleteInstance(ctx, id)
	} else if provider == "ibmcloud" {
		return p.vpc.DeleteInstance(ctx, id)
	}
	return fmt.Errorf("unknown provider: %s", provider)
}

func (p *ibmProvider) wrapID(provider string, instance *provider.Instance) *provider.Instance {
	if instance != nil {
		instance.ID = provider + "@" + instance.ID
	}
	return instance
}

func (p *ibmProvider) unwrapID(instanceID string) (string, string) {
	providerAndID := strings.Split(instanceID, "@")
	return providerAndID[0], providerAndID[1]
}

func (p *ibmProvider) Teardown() error {
	p.vpc.Teardown()
	p.powervs.Teardown()
	return nil
}

func (p *ibmProvider) ConfigVerifier() error {
	p.vpc.ConfigVerifier()
	p.powervs.ConfigVerifier()
	return nil
}
