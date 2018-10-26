/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	glog "github.com/golang/glog"
	corev1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/client/clientset/versioned/typed/core/v1alpha1"
	multiclusterdnsv1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/client/clientset/versioned/typed/multiclusterdns/v1alpha1"
	schedulingv1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/client/clientset/versioned/typed/scheduling/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	CoreV1alpha1() corev1alpha1.CoreV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Core() corev1alpha1.CoreV1alpha1Interface
	MulticlusterdnsV1alpha1() multiclusterdnsv1alpha1.MulticlusterdnsV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Multiclusterdns() multiclusterdnsv1alpha1.MulticlusterdnsV1alpha1Interface
	SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Scheduling() schedulingv1alpha1.SchedulingV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	coreV1alpha1            *corev1alpha1.CoreV1alpha1Client
	multiclusterdnsV1alpha1 *multiclusterdnsv1alpha1.MulticlusterdnsV1alpha1Client
	schedulingV1alpha1      *schedulingv1alpha1.SchedulingV1alpha1Client
}

// CoreV1alpha1 retrieves the CoreV1alpha1Client
func (c *Clientset) CoreV1alpha1() corev1alpha1.CoreV1alpha1Interface {
	return c.coreV1alpha1
}

// Deprecated: Core retrieves the default version of CoreClient.
// Please explicitly pick a version.
func (c *Clientset) Core() corev1alpha1.CoreV1alpha1Interface {
	return c.coreV1alpha1
}

// MulticlusterdnsV1alpha1 retrieves the MulticlusterdnsV1alpha1Client
func (c *Clientset) MulticlusterdnsV1alpha1() multiclusterdnsv1alpha1.MulticlusterdnsV1alpha1Interface {
	return c.multiclusterdnsV1alpha1
}

// Deprecated: Multiclusterdns retrieves the default version of MulticlusterdnsClient.
// Please explicitly pick a version.
func (c *Clientset) Multiclusterdns() multiclusterdnsv1alpha1.MulticlusterdnsV1alpha1Interface {
	return c.multiclusterdnsV1alpha1
}

// SchedulingV1alpha1 retrieves the SchedulingV1alpha1Client
func (c *Clientset) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return c.schedulingV1alpha1
}

// Deprecated: Scheduling retrieves the default version of SchedulingClient.
// Please explicitly pick a version.
func (c *Clientset) Scheduling() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return c.schedulingV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.coreV1alpha1, err = corev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.multiclusterdnsV1alpha1, err = multiclusterdnsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.schedulingV1alpha1, err = schedulingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		glog.Errorf("failed to create the DiscoveryClient: %v", err)
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.coreV1alpha1 = corev1alpha1.NewForConfigOrDie(c)
	cs.multiclusterdnsV1alpha1 = multiclusterdnsv1alpha1.NewForConfigOrDie(c)
	cs.schedulingV1alpha1 = schedulingv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.coreV1alpha1 = corev1alpha1.New(c)
	cs.multiclusterdnsV1alpha1 = multiclusterdnsv1alpha1.New(c)
	cs.schedulingV1alpha1 = schedulingv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
