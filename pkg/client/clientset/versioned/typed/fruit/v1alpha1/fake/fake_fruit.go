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

// FakeFruits implements FruitInterface
type FakeFruits struct {
	Fake *FakeFruitV1alpha1
	ns   string
}

var fruitsResource = schema.GroupVersionResource{Group: "sharedconfig.mycompany.com", Version: "v1alpha1", Resource: "fruits"}

var fruitsKind = schema.GroupVersionKind{Group: "sharedconfig.mycompany.com", Version: "v1alpha1", Kind: "Fruit"}

// Get takes name of the sharedconfig, and returns the corresponding sharedconfig object, and an error if there is any.
func (c *FakeFruits) Get(name string, options v1.GetOptions) (result *v1alpha1.Fruit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fruitsResource, c.ns, name), &v1alpha1.Fruit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Fruit), err
}

// List takes label and field selectors, and returns the list of Fruits that match those selectors.
func (c *FakeFruits) List(opts v1.ListOptions) (result *v1alpha1.FruitList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fruitsResource, fruitsKind, c.ns, opts), &v1alpha1.FruitList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FruitList{}
	for _, item := range obj.(*v1alpha1.FruitList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fruits.
func (c *FakeFruits) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fruitsResource, c.ns, opts))

}

// Create takes the representation of a sharedconfig and creates it.  Returns the server's representation of the sharedconfig, and an error, if there is any.
func (c *FakeFruits) Create(fruit *v1alpha1.Fruit) (result *v1alpha1.Fruit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fruitsResource, c.ns, fruit), &v1alpha1.Fruit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Fruit), err
}

// Update takes the representation of a sharedconfig and updates it. Returns the server's representation of the sharedconfig, and an error, if there is any.
func (c *FakeFruits) Update(fruit *v1alpha1.Fruit) (result *v1alpha1.Fruit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fruitsResource, c.ns, fruit), &v1alpha1.Fruit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Fruit), err
}

// Delete takes name of the sharedconfig and deletes it. Returns an error if one occurs.
func (c *FakeFruits) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(fruitsResource, c.ns, name), &v1alpha1.Fruit{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFruits) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fruitsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FruitList{})
	return err
}

// Patch applies the patch and returns the patched sharedconfig.
func (c *FakeFruits) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Fruit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fruitsResource, c.ns, name, data, subresources...), &v1alpha1.Fruit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Fruit), err
}
