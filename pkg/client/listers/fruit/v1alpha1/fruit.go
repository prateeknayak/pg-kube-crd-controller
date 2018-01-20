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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/prateeknayak/pg-kube-crd-controller/pkg/apis/fruit/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FruitLister helps list Fruits.
type FruitLister interface {
	// List lists all Fruits in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Fruit, err error)
	// Fruits returns an object that can list and get Fruits.
	Fruits(namespace string) FruitNamespaceLister
	FruitListerExpansion
}

// fruitLister implements the FruitLister interface.
type fruitLister struct {
	indexer cache.Indexer
}

// NewFruitLister returns a new FruitLister.
func NewFruitLister(indexer cache.Indexer) FruitLister {
	return &fruitLister{indexer: indexer}
}

// List lists all Fruits in the indexer.
func (s *fruitLister) List(selector labels.Selector) (ret []*v1alpha1.Fruit, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Fruit))
	})
	return ret, err
}

// Fruits returns an object that can list and get Fruits.
func (s *fruitLister) Fruits(namespace string) FruitNamespaceLister {
	return fruitNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FruitNamespaceLister helps list and get Fruits.
type FruitNamespaceLister interface {
	// List lists all Fruits in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Fruit, err error)
	// Get retrieves the Fruit from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Fruit, error)
	FruitNamespaceListerExpansion
}

// fruitNamespaceLister implements the FruitNamespaceLister
// interface.
type fruitNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Fruits in the indexer for a given namespace.
func (s fruitNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Fruit, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Fruit))
	})
	return ret, err
}

// Get retrieves the Fruit from the indexer for a given namespace and name.
func (s fruitNamespaceLister) Get(name string) (*v1alpha1.Fruit, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fruit"), name)
	}
	return obj.(*v1alpha1.Fruit), nil
}
