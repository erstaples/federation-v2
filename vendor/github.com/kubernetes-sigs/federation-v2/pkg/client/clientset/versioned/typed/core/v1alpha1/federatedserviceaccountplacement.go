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

package v1alpha1

import (
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/core/v1alpha1"
	scheme "github.com/kubernetes-sigs/federation-v2/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedServiceAccountPlacementsGetter has a method to return a FederatedServiceAccountPlacementInterface.
// A group's client should implement this interface.
type FederatedServiceAccountPlacementsGetter interface {
	FederatedServiceAccountPlacements(namespace string) FederatedServiceAccountPlacementInterface
}

// FederatedServiceAccountPlacementInterface has methods to work with FederatedServiceAccountPlacement resources.
type FederatedServiceAccountPlacementInterface interface {
	Create(*v1alpha1.FederatedServiceAccountPlacement) (*v1alpha1.FederatedServiceAccountPlacement, error)
	Update(*v1alpha1.FederatedServiceAccountPlacement) (*v1alpha1.FederatedServiceAccountPlacement, error)
	UpdateStatus(*v1alpha1.FederatedServiceAccountPlacement) (*v1alpha1.FederatedServiceAccountPlacement, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FederatedServiceAccountPlacement, error)
	List(opts v1.ListOptions) (*v1alpha1.FederatedServiceAccountPlacementList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedServiceAccountPlacement, err error)
	FederatedServiceAccountPlacementExpansion
}

// federatedServiceAccountPlacements implements FederatedServiceAccountPlacementInterface
type federatedServiceAccountPlacements struct {
	client rest.Interface
	ns     string
}

// newFederatedServiceAccountPlacements returns a FederatedServiceAccountPlacements
func newFederatedServiceAccountPlacements(c *CoreV1alpha1Client, namespace string) *federatedServiceAccountPlacements {
	return &federatedServiceAccountPlacements{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedServiceAccountPlacement, and returns the corresponding federatedServiceAccountPlacement object, and an error if there is any.
func (c *federatedServiceAccountPlacements) Get(name string, options v1.GetOptions) (result *v1alpha1.FederatedServiceAccountPlacement, err error) {
	result = &v1alpha1.FederatedServiceAccountPlacement{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedserviceaccountplacements").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedServiceAccountPlacements that match those selectors.
func (c *federatedServiceAccountPlacements) List(opts v1.ListOptions) (result *v1alpha1.FederatedServiceAccountPlacementList, err error) {
	result = &v1alpha1.FederatedServiceAccountPlacementList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedserviceaccountplacements").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedServiceAccountPlacements.
func (c *federatedServiceAccountPlacements) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedserviceaccountplacements").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a federatedServiceAccountPlacement and creates it.  Returns the server's representation of the federatedServiceAccountPlacement, and an error, if there is any.
func (c *federatedServiceAccountPlacements) Create(federatedServiceAccountPlacement *v1alpha1.FederatedServiceAccountPlacement) (result *v1alpha1.FederatedServiceAccountPlacement, err error) {
	result = &v1alpha1.FederatedServiceAccountPlacement{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedserviceaccountplacements").
		Body(federatedServiceAccountPlacement).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedServiceAccountPlacement and updates it. Returns the server's representation of the federatedServiceAccountPlacement, and an error, if there is any.
func (c *federatedServiceAccountPlacements) Update(federatedServiceAccountPlacement *v1alpha1.FederatedServiceAccountPlacement) (result *v1alpha1.FederatedServiceAccountPlacement, err error) {
	result = &v1alpha1.FederatedServiceAccountPlacement{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedserviceaccountplacements").
		Name(federatedServiceAccountPlacement.Name).
		Body(federatedServiceAccountPlacement).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedServiceAccountPlacements) UpdateStatus(federatedServiceAccountPlacement *v1alpha1.FederatedServiceAccountPlacement) (result *v1alpha1.FederatedServiceAccountPlacement, err error) {
	result = &v1alpha1.FederatedServiceAccountPlacement{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedserviceaccountplacements").
		Name(federatedServiceAccountPlacement.Name).
		SubResource("status").
		Body(federatedServiceAccountPlacement).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedServiceAccountPlacement and deletes it. Returns an error if one occurs.
func (c *federatedServiceAccountPlacements) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedserviceaccountplacements").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedServiceAccountPlacements) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedserviceaccountplacements").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedServiceAccountPlacement.
func (c *federatedServiceAccountPlacements) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedServiceAccountPlacement, err error) {
	result = &v1alpha1.FederatedServiceAccountPlacement{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedserviceaccountplacements").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
