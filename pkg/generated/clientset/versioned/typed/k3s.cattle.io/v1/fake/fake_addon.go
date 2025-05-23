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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/k3s-io/api/k3s.cattle.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAddons implements AddonInterface
type FakeAddons struct {
	Fake *FakeK3sV1
	ns   string
}

var addonsResource = v1.SchemeGroupVersion.WithResource("addons")

var addonsKind = v1.SchemeGroupVersion.WithKind("Addon")

// Get takes name of the addon, and returns the corresponding addon object, and an error if there is any.
func (c *FakeAddons) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Addon, err error) {
	emptyResult := &v1.Addon{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(addonsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Addon), err
}

// List takes label and field selectors, and returns the list of Addons that match those selectors.
func (c *FakeAddons) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AddonList, err error) {
	emptyResult := &v1.AddonList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(addonsResource, addonsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.AddonList{ListMeta: obj.(*v1.AddonList).ListMeta}
	for _, item := range obj.(*v1.AddonList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested addons.
func (c *FakeAddons) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(addonsResource, c.ns, opts))

}

// Create takes the representation of a addon and creates it.  Returns the server's representation of the addon, and an error, if there is any.
func (c *FakeAddons) Create(ctx context.Context, addon *v1.Addon, opts metav1.CreateOptions) (result *v1.Addon, err error) {
	emptyResult := &v1.Addon{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(addonsResource, c.ns, addon, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Addon), err
}

// Update takes the representation of a addon and updates it. Returns the server's representation of the addon, and an error, if there is any.
func (c *FakeAddons) Update(ctx context.Context, addon *v1.Addon, opts metav1.UpdateOptions) (result *v1.Addon, err error) {
	emptyResult := &v1.Addon{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(addonsResource, c.ns, addon, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Addon), err
}

// Delete takes name of the addon and deletes it. Returns an error if one occurs.
func (c *FakeAddons) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(addonsResource, c.ns, name, opts), &v1.Addon{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAddons) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(addonsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.AddonList{})
	return err
}

// Patch applies the patch and returns the patched addon.
func (c *FakeAddons) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Addon, err error) {
	emptyResult := &v1.Addon{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(addonsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Addon), err
}
