/*
MIT License

Copyright (c) 2018 Prateek Nayak

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package fake

import (
	v1alpha1 "github.com/prateeknayak/pg-kube-crd-controller/pkg/apis/sharedconfig/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSharedConfigs implements SharedConfigInterface
type FakeSharedConfigs struct {
	Fake *FakeSharedconfigV1alpha1
	ns   string
}

var sharedconfigsResource = schema.GroupVersionResource{Group: "sharedconfig.mycompany.com", Version: "v1alpha1", Resource: "sharedconfigs"}

var sharedconfigsKind = schema.GroupVersionKind{Group: "sharedconfig.mycompany.com", Version: "v1alpha1", Kind: "SharedConfig"}

// Get takes name of the sharedConfig, and returns the corresponding sharedConfig object, and an error if there is any.
func (c *FakeSharedConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.SharedConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sharedconfigsResource, c.ns, name), &v1alpha1.SharedConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedConfig), err
}

// List takes label and field selectors, and returns the list of SharedConfigs that match those selectors.
func (c *FakeSharedConfigs) List(opts v1.ListOptions) (result *v1alpha1.SharedConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sharedconfigsResource, sharedconfigsKind, c.ns, opts), &v1alpha1.SharedConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SharedConfigList{}
	for _, item := range obj.(*v1alpha1.SharedConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sharedConfigs.
func (c *FakeSharedConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sharedconfigsResource, c.ns, opts))

}

// Create takes the representation of a sharedConfig and creates it.  Returns the server's representation of the sharedConfig, and an error, if there is any.
func (c *FakeSharedConfigs) Create(sharedConfig *v1alpha1.SharedConfig) (result *v1alpha1.SharedConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sharedconfigsResource, c.ns, sharedConfig), &v1alpha1.SharedConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedConfig), err
}

// Update takes the representation of a sharedConfig and updates it. Returns the server's representation of the sharedConfig, and an error, if there is any.
func (c *FakeSharedConfigs) Update(sharedConfig *v1alpha1.SharedConfig) (result *v1alpha1.SharedConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sharedconfigsResource, c.ns, sharedConfig), &v1alpha1.SharedConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedConfig), err
}

// Delete takes name of the sharedConfig and deletes it. Returns an error if one occurs.
func (c *FakeSharedConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sharedconfigsResource, c.ns, name), &v1alpha1.SharedConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSharedConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sharedconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SharedConfigList{})
	return err
}

// Patch applies the patch and returns the patched sharedConfig.
func (c *FakeSharedConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SharedConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sharedconfigsResource, c.ns, name, data, subresources...), &v1alpha1.SharedConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedConfig), err
}
