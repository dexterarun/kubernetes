/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen with arguments: --input-dirs=[k8s.io/kubernetes/pkg/api,k8s.io/kubernetes/pkg/api/v1,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/certificates/v1alpha1,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/meta/v1,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/policy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/apis/storage/v1beta1]

package v1

import (
	"k8s.io/apimachinery/pkg/labels"
	api "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/errors"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/client/cache"
)

// LimitRangeLister helps list LimitRanges.
type LimitRangeLister interface {
	// List lists all LimitRanges in the indexer.
	List(selector labels.Selector) (ret []*v1.LimitRange, err error)
	// LimitRanges returns an object that can list and get LimitRanges.
	LimitRanges(namespace string) LimitRangeNamespaceLister
	LimitRangeListerExpansion
}

// limitRangeLister implements the LimitRangeLister interface.
type limitRangeLister struct {
	indexer cache.Indexer
}

// NewLimitRangeLister returns a new LimitRangeLister.
func NewLimitRangeLister(indexer cache.Indexer) LimitRangeLister {
	return &limitRangeLister{indexer: indexer}
}

// List lists all LimitRanges in the indexer.
func (s *limitRangeLister) List(selector labels.Selector) (ret []*v1.LimitRange, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.LimitRange))
	})
	return ret, err
}

// LimitRanges returns an object that can list and get LimitRanges.
func (s *limitRangeLister) LimitRanges(namespace string) LimitRangeNamespaceLister {
	return limitRangeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LimitRangeNamespaceLister helps list and get LimitRanges.
type LimitRangeNamespaceLister interface {
	// List lists all LimitRanges in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.LimitRange, err error)
	// Get retrieves the LimitRange from the indexer for a given namespace and name.
	Get(name string) (*v1.LimitRange, error)
	LimitRangeNamespaceListerExpansion
}

// limitRangeNamespaceLister implements the LimitRangeNamespaceLister
// interface.
type limitRangeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LimitRanges in the indexer for a given namespace.
func (s limitRangeNamespaceLister) List(selector labels.Selector) (ret []*v1.LimitRange, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.LimitRange))
	})
	return ret, err
}

// Get retrieves the LimitRange from the indexer for a given namespace and name.
func (s limitRangeNamespaceLister) Get(name string) (*v1.LimitRange, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(api.Resource("limitrange"), name)
	}
	return obj.(*v1.LimitRange), nil
}
