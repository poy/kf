// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	versioned "github.com/poy/kf/pkg/client/servicecatalog/clientset/versioned"
	internalinterfaces "github.com/poy/kf/pkg/client/servicecatalog/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/poy/kf/pkg/client/servicecatalog/listers/servicecatalog/v1beta1"
	servicecatalogv1beta1 "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterServiceBrokerInformer provides access to a shared informer and lister for
// ClusterServiceBrokers.
type ClusterServiceBrokerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ClusterServiceBrokerLister
}

type clusterServiceBrokerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterServiceBrokerInformer constructs a new informer for ClusterServiceBroker type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterServiceBrokerInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterServiceBrokerInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterServiceBrokerInformer constructs a new informer for ClusterServiceBroker type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterServiceBrokerInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServicecatalogV1beta1().ClusterServiceBrokers().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServicecatalogV1beta1().ClusterServiceBrokers().Watch(options)
			},
		},
		&servicecatalogv1beta1.ClusterServiceBroker{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterServiceBrokerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterServiceBrokerInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterServiceBrokerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&servicecatalogv1beta1.ClusterServiceBroker{}, f.defaultInformer)
}

func (f *clusterServiceBrokerInformer) Lister() v1beta1.ClusterServiceBrokerLister {
	return v1beta1.NewClusterServiceBrokerLister(f.Informer().GetIndexer())
}
