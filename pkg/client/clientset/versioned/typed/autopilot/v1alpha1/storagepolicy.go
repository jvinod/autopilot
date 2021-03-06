/*
Copyright 2019 Openstorage.org

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
	"time"

	v1alpha1 "github.com/libopenstorage/autopilot/pkg/apis/autopilot/v1alpha1"
	scheme "github.com/libopenstorage/autopilot/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// StoragePoliciesGetter has a method to return a StoragePolicyInterface.
// A group's client should implement this interface.
type StoragePoliciesGetter interface {
	StoragePolicies() StoragePolicyInterface
}

// StoragePolicyInterface has methods to work with StoragePolicy resources.
type StoragePolicyInterface interface {
	Create(*v1alpha1.StoragePolicy) (*v1alpha1.StoragePolicy, error)
	Update(*v1alpha1.StoragePolicy) (*v1alpha1.StoragePolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.StoragePolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.StoragePolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StoragePolicy, err error)
	StoragePolicyExpansion
}

// storagePolicies implements StoragePolicyInterface
type storagePolicies struct {
	client rest.Interface
}

// newStoragePolicies returns a StoragePolicies
func newStoragePolicies(c *AutopilotV1alpha1Client) *storagePolicies {
	return &storagePolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the storagePolicy, and returns the corresponding storagePolicy object, and an error if there is any.
func (c *storagePolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.StoragePolicy, err error) {
	result = &v1alpha1.StoragePolicy{}
	err = c.client.Get().
		Resource("storagepolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StoragePolicies that match those selectors.
func (c *storagePolicies) List(opts v1.ListOptions) (result *v1alpha1.StoragePolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.StoragePolicyList{}
	err = c.client.Get().
		Resource("storagepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storagePolicies.
func (c *storagePolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("storagepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a storagePolicy and creates it.  Returns the server's representation of the storagePolicy, and an error, if there is any.
func (c *storagePolicies) Create(storagePolicy *v1alpha1.StoragePolicy) (result *v1alpha1.StoragePolicy, err error) {
	result = &v1alpha1.StoragePolicy{}
	err = c.client.Post().
		Resource("storagepolicies").
		Body(storagePolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a storagePolicy and updates it. Returns the server's representation of the storagePolicy, and an error, if there is any.
func (c *storagePolicies) Update(storagePolicy *v1alpha1.StoragePolicy) (result *v1alpha1.StoragePolicy, err error) {
	result = &v1alpha1.StoragePolicy{}
	err = c.client.Put().
		Resource("storagepolicies").
		Name(storagePolicy.Name).
		Body(storagePolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the storagePolicy and deletes it. Returns an error if one occurs.
func (c *storagePolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("storagepolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storagePolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("storagepolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched storagePolicy.
func (c *storagePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StoragePolicy, err error) {
	result = &v1alpha1.StoragePolicy{}
	err = c.client.Patch(pt).
		Resource("storagepolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
