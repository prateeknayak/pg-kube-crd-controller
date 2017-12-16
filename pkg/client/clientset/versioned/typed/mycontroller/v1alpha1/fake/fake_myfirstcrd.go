package fake

import (
	v1alpha1 "github.com/prateeknayak/pg-kube-crd-controller/pkg/apis/mycontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMyFirstCRDs implements MyFirstCRDInterface
type FakeMyFirstCRDs struct {
	Fake *FakeMycontrollerV1alpha1
	ns   string
}

var myfirstcrdsResource = schema.GroupVersionResource{Group: "mycontroller.k8s.io", Version: "v1alpha1", Resource: "myfirstcrds"}

var myfirstcrdsKind = schema.GroupVersionKind{Group: "mycontroller.k8s.io", Version: "v1alpha1", Kind: "MyFirstCRD"}

// Get takes name of the myFirstCRD, and returns the corresponding myFirstCRD object, and an error if there is any.
func (c *FakeMyFirstCRDs) Get(name string, options v1.GetOptions) (result *v1alpha1.MyFirstCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(myfirstcrdsResource, c.ns, name), &v1alpha1.MyFirstCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyFirstCRD), err
}

// List takes label and field selectors, and returns the list of MyFirstCRDs that match those selectors.
func (c *FakeMyFirstCRDs) List(opts v1.ListOptions) (result *v1alpha1.MyFirstCRDList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(myfirstcrdsResource, myfirstcrdsKind, c.ns, opts), &v1alpha1.MyFirstCRDList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MyFirstCRDList{}
	for _, item := range obj.(*v1alpha1.MyFirstCRDList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested myFirstCRDs.
func (c *FakeMyFirstCRDs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(myfirstcrdsResource, c.ns, opts))

}

// Create takes the representation of a myFirstCRD and creates it.  Returns the server's representation of the myFirstCRD, and an error, if there is any.
func (c *FakeMyFirstCRDs) Create(myFirstCRD *v1alpha1.MyFirstCRD) (result *v1alpha1.MyFirstCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(myfirstcrdsResource, c.ns, myFirstCRD), &v1alpha1.MyFirstCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyFirstCRD), err
}

// Update takes the representation of a myFirstCRD and updates it. Returns the server's representation of the myFirstCRD, and an error, if there is any.
func (c *FakeMyFirstCRDs) Update(myFirstCRD *v1alpha1.MyFirstCRD) (result *v1alpha1.MyFirstCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(myfirstcrdsResource, c.ns, myFirstCRD), &v1alpha1.MyFirstCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyFirstCRD), err
}

// Delete takes name of the myFirstCRD and deletes it. Returns an error if one occurs.
func (c *FakeMyFirstCRDs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(myfirstcrdsResource, c.ns, name), &v1alpha1.MyFirstCRD{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMyFirstCRDs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(myfirstcrdsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MyFirstCRDList{})
	return err
}

// Patch applies the patch and returns the patched myFirstCRD.
func (c *FakeMyFirstCRDs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MyFirstCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(myfirstcrdsResource, c.ns, name, data, subresources...), &v1alpha1.MyFirstCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyFirstCRD), err
}
