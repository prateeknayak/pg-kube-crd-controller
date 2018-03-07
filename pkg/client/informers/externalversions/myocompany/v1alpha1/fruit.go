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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	time "time"

	fruit_v1alpha1 "github.com/prateeknayak/pg-kube-crd-controller/pkg/apis/sharedconfig/v1alpha1"
	versioned "github.com/prateeknayak/pg-kube-crd-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/prateeknayak/pg-kube-crd-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/prateeknayak/pg-kube-crd-controller/pkg/client/listers/myocompany/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FruitInformer provides access to a shared informer and lister for
// Fruits.
type FruitInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FruitLister
}

type fruitInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewFruitInformer constructs a new informer for Fruit type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFruitInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.MyocompanyV1alpha1().Fruits(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.MyocompanyV1alpha1().Fruits(namespace).Watch(options)
			},
		},
		&fruit_v1alpha1.Fruit{},
		resyncPeriod,
		indexers,
	)
}

func defaultFruitInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFruitInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *fruitInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&fruit_v1alpha1.Fruit{}, defaultFruitInformer)
}

func (f *fruitInformer) Lister() v1alpha1.FruitLister {
	return v1alpha1.NewFruitLister(f.Informer().GetIndexer())
}
