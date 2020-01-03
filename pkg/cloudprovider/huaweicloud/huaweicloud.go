/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package huaweicloud

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"k8s.io/cloud-provider"
	"k8s.io/klog"
)

const (
	providerName = "huaweicloud"
)

// CloudConfig is used to read and store information from the cloud configuration file
type CloudConfig struct {
	// TODO(RainbowMango): Auth options shall be split from LoadBalancer.
	LoadBalancer LoadBalancerOpts
}

func init() {
	cloudprovider.RegisterCloudProvider(providerName, newCloud)
}

func readConfig(config io.Reader) (*CloudConfig, error) {
	if config == nil {
		return nil, fmt.Errorf("no cloud provider config given")
	}

	configBytes, err := ioutil.ReadAll(config)
	if err != nil {
		return nil, fmt.Errorf("can not read cloud provider config: %v", err)
	}

	var cloudConfig CloudConfig
	err = json.Unmarshal(configBytes, &cloudConfig)
	if err != nil {
		return nil, fmt.Errorf("cloud config format is not expected: %v", err)
	}

	return &cloudConfig, nil
}

func newCloud(config io.Reader) (cloudprovider.Interface, error) {
	cloudConfig, err := readConfig(config)
	if err != nil {
		klog.Errorf("Create cloud provider failed: %v.", err)
		return nil, err
	}

	return NewHuaweiCloud(cloudConfig)
}

func NewHuaweiCloud(config *CloudConfig) (*HuaweiCloud, error) {
	// TODO(RainbowMango): Create Huawei Cloud Instance here
	return nil, nil
}

// HuaweiCloud is an implementation of cloud provider Interface for Huawei Cloud.
type HuaweiCloud struct {
}

var _ cloudprovider.Interface = &HuaweiCloud{}

// Initialize provides the cloud with a kubernetes client builder and may spawn goroutines
// to perform housekeeping or run custom controllers specific to the cloud provider.
// Any tasks started here should be cleaned up when the stop channel closes.
func (h *HuaweiCloud) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {

}

// LoadBalancer returns a balancer interface. Also returns true if the interface is supported, false otherwise.
func (h *HuaweiCloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	// TODO(RainbowMango): waiting a solution about how to share openstack implementation and do minimum changes here.
	return nil, false
}

// Instances returns an instances interface. Also returns true if the interface is supported, false otherwise.
func (h *HuaweiCloud) Instances() (cloudprovider.Instances, bool) {
	// TODO(RainbowMango): waiting a solution about how to share openstack implementation and do minimum changes here.
	return nil, false
}

// Zones returns a zones interface. Also returns true if the interface is supported, false otherwise.
func (h *HuaweiCloud) Zones() (cloudprovider.Zones, bool) {
	// TODO(RainbowMango): waiting a solution about how to share openstack implementation and do minimum changes here.
	return nil, false
}

// Clusters returns a clusters interface.  Also returns true if the interface is supported, false otherwise.
func (h *HuaweiCloud) Clusters() (cloudprovider.Clusters, bool) {
	// TODO(RainbowMango): waiting a solution about how to share openstack implementation and do minimum changes here.
	return nil, false
}

// Routes returns a routes interface along with whether the interface is supported.
func (h *HuaweiCloud) Routes() (cloudprovider.Routes, bool) {
	// TODO(RainbowMango): waiting a solution about how to share openstack implementation and do minimum changes here.
	return nil, false
}

// HuaweiCloudProviderName returns the cloud provider ID.
func (h *HuaweiCloud) ProviderName() string {
	return providerName
}

// HasClusterID returns true if a ClusterID is required and set
func (h *HuaweiCloud) HasClusterID() bool {
	return true
}