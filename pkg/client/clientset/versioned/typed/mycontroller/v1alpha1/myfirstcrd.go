package v1alpha1

import (
	v1alpha1 "github.com/prateeknayak/pg-kube-crd-controller/pkg/apis/mycontroller/v1alpha1"
	scheme "github.com/prateeknayak/pg-kube-crd-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MyFirstCRDsGetter has a method to return a MyFirstCRDInterface.
// A group's client should implement this interface.
type MyFirstCRDsGetter interface {
	MyFirstCRDs(namespace string) MyFirstCRDInterface
}

// MyFirstCRDInterface has methods to work with MyFirstCRD resources.
type MyFirstCRDInterface interface {
	Create(*v1alpha1.MyFirstCRD) (*v1alpha1.MyFirstCRD, error)
	Update(*v1alpha1.MyFirstCRD) (*v1alpha1.MyFirstCRD, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MyFirstCRD, error)
	List(opts v1.ListOptions) (*v1alpha1.MyFirstCRDList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MyFirstCRD, err error)
	MyFirstCRDExpansion
}

// myFirstCRDs implements MyFirstCRDInterface
type myFirstCRDs struct {
	client rest.Interface
	ns     string
}

// newMyFirstCRDs returns a MyFirstCRDs
func newMyFirstCRDs(c *MycontrollerV1alpha1Client, namespace string) *myFirstCRDs {
	return &myFirstCRDs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the myFirstCRD, and returns the corresponding myFirstCRD object, and an error if there is any.
func (c *myFirstCRDs) Get(name string, options v1.GetOptions) (result *v1alpha1.MyFirstCRD, err error) {
	result = &v1alpha1.MyFirstCRD{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("myfirstcrds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MyFirstCRDs that match those selectors.
func (c *myFirstCRDs) List(opts v1.ListOptions) (result *v1alpha1.MyFirstCRDList, err error) {
	result = &v1alpha1.MyFirstCRDList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("myfirstcrds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested myFirstCRDs.
func (c *myFirstCRDs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("myfirstcrds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a myFirstCRD and creates it.  Returns the server's representation of the myFirstCRD, and an error, if there is any.
func (c *myFirstCRDs) Create(myFirstCRD *v1alpha1.MyFirstCRD) (result *v1alpha1.MyFirstCRD, err error) {
	result = &v1alpha1.MyFirstCRD{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("myfirstcrds").
		Body(myFirstCRD).
		Do().
		Into(result)
	return
}

// Update takes the representation of a myFirstCRD and updates it. Returns the server's representation of the myFirstCRD, and an error, if there is any.
func (c *myFirstCRDs) Update(myFirstCRD *v1alpha1.MyFirstCRD) (result *v1alpha1.MyFirstCRD, err error) {
	result = &v1alpha1.MyFirstCRD{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("myfirstcrds").
		Name(myFirstCRD.Name).
		Body(myFirstCRD).
		Do().
		Into(result)
	return
}

// Delete takes name of the myFirstCRD and deletes it. Returns an error if one occurs.
func (c *myFirstCRDs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("myfirstcrds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *myFirstCRDs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("myfirstcrds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched myFirstCRD.
func (c *myFirstCRDs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MyFirstCRD, err error) {
	result = &v1alpha1.MyFirstCRD{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("myfirstcrds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
