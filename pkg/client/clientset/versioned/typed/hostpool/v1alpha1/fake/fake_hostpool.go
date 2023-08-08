/*
Copyright 2021-2022 Red Hat, Inc.

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

package fake

import (
	"context"

	v1alpha1 "github.com/stuartwdouglas/multi-arch-host-resolver/pkg/apis/hostpool/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHostPools implements HostPoolInterface
type FakeHostPools struct {
	Fake *FakeMultiarchV1alpha1
	ns   string
}

var hostpoolsResource = schema.GroupVersionResource{Group: "multiarch.appstudio.redhat.com", Version: "v1alpha1", Resource: "hostpools"}

var hostpoolsKind = schema.GroupVersionKind{Group: "multiarch.appstudio.redhat.com", Version: "v1alpha1", Kind: "HostPool"}

// Get takes name of the hostPool, and returns the corresponding hostPool object, and an error if there is any.
func (c *FakeHostPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HostPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hostpoolsResource, c.ns, name), &v1alpha1.HostPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostPool), err
}

// List takes label and field selectors, and returns the list of HostPools that match those selectors.
func (c *FakeHostPools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HostPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hostpoolsResource, hostpoolsKind, c.ns, opts), &v1alpha1.HostPoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HostPoolList{ListMeta: obj.(*v1alpha1.HostPoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.HostPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hostPools.
func (c *FakeHostPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hostpoolsResource, c.ns, opts))

}

// Create takes the representation of a hostPool and creates it.  Returns the server's representation of the hostPool, and an error, if there is any.
func (c *FakeHostPools) Create(ctx context.Context, hostPool *v1alpha1.HostPool, opts v1.CreateOptions) (result *v1alpha1.HostPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hostpoolsResource, c.ns, hostPool), &v1alpha1.HostPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostPool), err
}

// Update takes the representation of a hostPool and updates it. Returns the server's representation of the hostPool, and an error, if there is any.
func (c *FakeHostPools) Update(ctx context.Context, hostPool *v1alpha1.HostPool, opts v1.UpdateOptions) (result *v1alpha1.HostPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hostpoolsResource, c.ns, hostPool), &v1alpha1.HostPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostPool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHostPools) UpdateStatus(ctx context.Context, hostPool *v1alpha1.HostPool, opts v1.UpdateOptions) (*v1alpha1.HostPool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hostpoolsResource, "status", c.ns, hostPool), &v1alpha1.HostPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostPool), err
}

// Delete takes name of the hostPool and deletes it. Returns an error if one occurs.
func (c *FakeHostPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(hostpoolsResource, c.ns, name, opts), &v1alpha1.HostPool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHostPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hostpoolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HostPoolList{})
	return err
}

// Patch applies the patch and returns the patched hostPool.
func (c *FakeHostPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HostPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hostpoolsResource, c.ns, name, pt, data, subresources...), &v1alpha1.HostPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostPool), err
}