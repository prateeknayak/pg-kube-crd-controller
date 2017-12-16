package fake

import (
	v1alpha1 "github.com/prateeknayak/pg-kube-crd-controller/pkg/client/clientset/versioned/typed/mycontroller/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMycontrollerV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMycontrollerV1alpha1) MyFirstCRDs(namespace string) v1alpha1.MyFirstCRDInterface {
	return &FakeMyFirstCRDs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMycontrollerV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
