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

// SharedConfigsGetter has a method to return a SharedConfigInterface.
// A group's client should implement this interface.
type SharedConfigsGetter interface {
	SharedConfigs(namespace string) SharedConfigInterface
}

// SharedConfigInterface has methods to work with SharedConfig resources.
type SharedConfigInterface interface {
	Create(*v1alpha1.SharedConfig) (*v1alpha1.SharedConfig, error)
	Update(*v1alpha1.SharedConfig) (*v1alpha1.SharedConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SharedConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.SharedConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SharedConfig, err error)
	SharedConfigExpansion
}

// sharedConfigs implements SharedConfigInterface
type sharedConfigs struct {
	client rest.Interface
	ns     string
}

// newSharedConfigs returns a SharedConfigs
func newSharedConfigs(c *SharedconfigV1alpha1Client, namespace string) *sharedConfigs {
	return &sharedConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sharedConfig, and returns the corresponding sharedConfig object, and an error if there is any.
func (c *sharedConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.SharedConfig, err error) {
	result = &v1alpha1.SharedConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sharedconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SharedConfigs that match those selectors.
func (c *sharedConfigs) List(opts v1.ListOptions) (result *v1alpha1.SharedConfigList, err error) {
	result = &v1alpha1.SharedConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sharedconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sharedConfigs.
func (c *sharedConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sharedconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a sharedConfig and creates it.  Returns the server's representation of the sharedConfig, and an error, if there is any.
func (c *sharedConfigs) Create(sharedConfig *v1alpha1.SharedConfig) (result *v1alpha1.SharedConfig, err error) {
	result = &v1alpha1.SharedConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sharedconfigs").
		Body(sharedConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sharedConfig and updates it. Returns the server's representation of the sharedConfig, and an error, if there is any.
func (c *sharedConfigs) Update(sharedConfig *v1alpha1.SharedConfig) (result *v1alpha1.SharedConfig, err error) {
	result = &v1alpha1.SharedConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sharedconfigs").
		Name(sharedConfig.Name).
		Body(sharedConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the sharedConfig and deletes it. Returns an error if one occurs.
func (c *sharedConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sharedconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sharedConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sharedconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sharedConfig.
func (c *sharedConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SharedConfig, err error) {
	result = &v1alpha1.SharedConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sharedconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
