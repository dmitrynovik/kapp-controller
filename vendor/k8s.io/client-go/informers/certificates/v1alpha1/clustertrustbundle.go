/*
Copyright The Kubernetes Authors.

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
	"context"
	time "time"

	certificatesv1alpha1 "k8s.io/api/certificates/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1alpha1 "k8s.io/client-go/listers/certificates/v1alpha1"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterTrustBundleInformer provides access to a shared informer and lister for
// ClusterTrustBundles.
type ClusterTrustBundleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterTrustBundleLister
}

type clusterTrustBundleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterTrustBundleInformer constructs a new informer for ClusterTrustBundle type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterTrustBundleInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterTrustBundleInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterTrustBundleInformer constructs a new informer for ClusterTrustBundle type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterTrustBundleInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertificatesV1alpha1().ClusterTrustBundles().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertificatesV1alpha1().ClusterTrustBundles().Watch(context.TODO(), options)
			},
		},
		&certificatesv1alpha1.ClusterTrustBundle{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterTrustBundleInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterTrustBundleInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterTrustBundleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&certificatesv1alpha1.ClusterTrustBundle{}, f.defaultInformer)
}

func (f *clusterTrustBundleInformer) Lister() v1alpha1.ClusterTrustBundleLister {
	return v1alpha1.NewClusterTrustBundleLister(f.Informer().GetIndexer())
}
