/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"
	"time"

	v1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/crd_generated/apis/servicegroup.tsm.tanzu.vmware.com/v1"
	scheme "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/crd_generated/client/clientset/versioned/scheme"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SvcGroupsGetter has a method to return a SvcGroupInterface.
// A group's client should implement this interface.
type SvcGroupsGetter interface {
	SvcGroups() SvcGroupInterface
}

// SvcGroupInterface has methods to work with SvcGroup resources.
type SvcGroupInterface interface {
	Create(ctx context.Context, svcGroup *v1.SvcGroup, opts metav1.CreateOptions) (*v1.SvcGroup, error)
	Update(ctx context.Context, svcGroup *v1.SvcGroup, opts metav1.UpdateOptions) (*v1.SvcGroup, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SvcGroup, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SvcGroupList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SvcGroup, err error)
	SvcGroupExpansion
}

// svcGroups implements SvcGroupInterface
type svcGroups struct {
	client rest.Interface
}

// newSvcGroups returns a SvcGroups
func newSvcGroups(c *ServicegroupTsmV1Client) *svcGroups {
	return &svcGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the svcGroup, and returns the corresponding svcGroup object, and an error if there is any.
func (c *svcGroups) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SvcGroup, err error) {
	result = &v1.SvcGroup{}
	err = c.client.Get().
		Resource("svcgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SvcGroups that match those selectors.
func (c *svcGroups) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SvcGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SvcGroupList{}
	err = c.client.Get().
		Resource("svcgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested svcGroups.
func (c *svcGroups) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("svcgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a svcGroup and creates it.  Returns the server's representation of the svcGroup, and an error, if there is any.
func (c *svcGroups) Create(ctx context.Context, svcGroup *v1.SvcGroup, opts metav1.CreateOptions) (result *v1.SvcGroup, err error) {
	result = &v1.SvcGroup{}
	err = c.client.Post().
		Resource("svcgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(svcGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a svcGroup and updates it. Returns the server's representation of the svcGroup, and an error, if there is any.
func (c *svcGroups) Update(ctx context.Context, svcGroup *v1.SvcGroup, opts metav1.UpdateOptions) (result *v1.SvcGroup, err error) {
	result = &v1.SvcGroup{}
	err = c.client.Put().
		Resource("svcgroups").
		Name(svcGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(svcGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the svcGroup and deletes it. Returns an error if one occurs.
func (c *svcGroups) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("svcgroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *svcGroups) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("svcgroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched svcGroup.
func (c *svcGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SvcGroup, err error) {
	result = &v1.SvcGroup{}
	err = c.client.Patch(pt).
		Resource("svcgroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
