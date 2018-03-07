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
package v1alpha1

import (
	v1alpha1 "github.com/prateeknayak/pg-kube-crd-controller/pkg/apis/sharedconfig/v1alpha1"
	scheme "github.com/prateeknayak/pg-kube-crd-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FruitsGetter has a method to return a FruitInterface.
// A group's client should implement this interface.
type FruitsGetter interface {
	Fruits(namespace string) FruitInterface
}

// FruitInterface has methods to work with Fruit resources.
type FruitInterface interface {
	Create(*v1alpha1.Fruit) (*v1alpha1.Fruit, error)
	Update(*v1alpha1.Fruit) (*v1alpha1.Fruit, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Fruit, error)
	List(opts v1.ListOptions) (*v1alpha1.FruitList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Fruit, err error)
	FruitExpansion
}

// fruits implements FruitInterface
type fruits struct {
	client rest.Interface
	ns     string
}

// newFruits returns a Fruits
func newFruits(c *FruitV1alpha1Client, namespace string) *fruits {
	return &fruits{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sharedconfig, and returns the corresponding sharedconfig object, and an error if there is any.
func (c *fruits) Get(name string, options v1.GetOptions) (result *v1alpha1.Fruit, err error) {
	result = &v1alpha1.Fruit{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fruits").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Fruits that match those selectors.
func (c *fruits) List(opts v1.ListOptions) (result *v1alpha1.FruitList, err error) {
	result = &v1alpha1.FruitList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fruits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested fruits.
func (c *fruits) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("fruits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a sharedconfig and creates it.  Returns the server's representation of the sharedconfig, and an error, if there is any.
func (c *fruits) Create(fruit *v1alpha1.Fruit) (result *v1alpha1.Fruit, err error) {
	result = &v1alpha1.Fruit{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("fruits").
		Body(fruit).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sharedconfig and updates it. Returns the server's representation of the sharedconfig, and an error, if there is any.
func (c *fruits) Update(fruit *v1alpha1.Fruit) (result *v1alpha1.Fruit, err error) {
	result = &v1alpha1.Fruit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("fruits").
		Name(fruit.Name).
		Body(fruit).
		Do().
		Into(result)
	return
}

// Delete takes name of the sharedconfig and deletes it. Returns an error if one occurs.
func (c *fruits) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fruits").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *fruits) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fruits").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sharedconfig.
func (c *fruits) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Fruit, err error) {
	result = &v1alpha1.Fruit{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("fruits").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
