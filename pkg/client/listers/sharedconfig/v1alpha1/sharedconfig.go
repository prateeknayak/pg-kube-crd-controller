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
	v1alpha1 "github.com/prateeknayak/pg-kube-crd-controller/pkg/apis/sharedconfig/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SharedConfigLister helps list SharedConfigs.
type SharedConfigLister interface {
	// List lists all SharedConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SharedConfig, err error)
	// SharedConfigs returns an object that can list and get SharedConfigs.
	SharedConfigs(namespace string) SharedConfigNamespaceLister
	SharedConfigListerExpansion
}

// sharedConfigLister implements the SharedConfigLister interface.
type sharedConfigLister struct {
	indexer cache.Indexer
}

// NewSharedConfigLister returns a new SharedConfigLister.
func NewSharedConfigLister(indexer cache.Indexer) SharedConfigLister {
	return &sharedConfigLister{indexer: indexer}
}

// List lists all SharedConfigs in the indexer.
func (s *sharedConfigLister) List(selector labels.Selector) (ret []*v1alpha1.SharedConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SharedConfig))
	})
	return ret, err
}

// SharedConfigs returns an object that can list and get SharedConfigs.
func (s *sharedConfigLister) SharedConfigs(namespace string) SharedConfigNamespaceLister {
	return sharedConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SharedConfigNamespaceLister helps list and get SharedConfigs.
type SharedConfigNamespaceLister interface {
	// List lists all SharedConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SharedConfig, err error)
	// Get retrieves the SharedConfig from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SharedConfig, error)
	SharedConfigNamespaceListerExpansion
}

// sharedConfigNamespaceLister implements the SharedConfigNamespaceLister
// interface.
type sharedConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SharedConfigs in the indexer for a given namespace.
func (s sharedConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SharedConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SharedConfig))
	})
	return ret, err
}

// Get retrieves the SharedConfig from the indexer for a given namespace and name.
func (s sharedConfigNamespaceLister) Get(name string) (*v1alpha1.SharedConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sharedconfig"), name)
	}
	return obj.(*v1alpha1.SharedConfig), nil
}
