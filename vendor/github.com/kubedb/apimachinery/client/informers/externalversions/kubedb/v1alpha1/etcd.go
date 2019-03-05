/*
Copyright 2019 The KubeDB Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	kubedbv1alpha1 "github.com/kubedb/apimachinery/apis/kubedb/v1alpha1"
	versioned "github.com/kubedb/apimachinery/client/clientset/versioned"
	internalinterfaces "github.com/kubedb/apimachinery/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubedb/apimachinery/client/listers/kubedb/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EtcdInformer provides access to a shared informer and lister for
// Etcds.
type EtcdInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.EtcdLister
}

type etcdInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEtcdInformer constructs a new informer for Etcd type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEtcdInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEtcdInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEtcdInformer constructs a new informer for Etcd type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEtcdInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubedbV1alpha1().Etcds(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubedbV1alpha1().Etcds(namespace).Watch(options)
			},
		},
		&kubedbv1alpha1.Etcd{},
		resyncPeriod,
		indexers,
	)
}

func (f *etcdInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEtcdInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *etcdInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubedbv1alpha1.Etcd{}, f.defaultInformer)
}

func (f *etcdInformer) Lister() v1alpha1.EtcdLister {
	return v1alpha1.NewEtcdLister(f.Informer().GetIndexer())
}
