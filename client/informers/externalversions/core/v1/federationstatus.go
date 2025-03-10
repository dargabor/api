// Copyright Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	internalinterfaces "maistra.io/api/client/informers/externalversions/internalinterfaces"
	v1 "maistra.io/api/client/listers/core/v1"
	versioned "maistra.io/api/client/versioned"
	corev1 "maistra.io/api/core/v1"
)

// FederationStatusInformer provides access to a shared informer and lister for
// FederationStatuses.
type FederationStatusInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.FederationStatusLister
}

type federationStatusInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFederationStatusInformer constructs a new informer for FederationStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederationStatusInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederationStatusInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFederationStatusInformer constructs a new informer for FederationStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederationStatusInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().FederationStatuses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().FederationStatuses(namespace).Watch(context.TODO(), options)
			},
		},
		&corev1.FederationStatus{},
		resyncPeriod,
		indexers,
	)
}

func (f *federationStatusInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederationStatusInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federationStatusInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.FederationStatus{}, f.defaultInformer)
}

func (f *federationStatusInformer) Lister() v1.FederationStatusLister {
	return v1.NewFederationStatusLister(f.Informer().GetIndexer())
}
