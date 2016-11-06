/*
Copyright 2016 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen with arguments: --input-dirs=[k8s.io/kubernetes/pkg/api,k8s.io/kubernetes/pkg/api/v1,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/certificates/v1alpha1,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/policy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/apis/storage/v1beta1]

package internalversion

import (
	"k8s.io/kubernetes/pkg/api/errors"
	policy "k8s.io/kubernetes/pkg/apis/policy"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/labels"
)

// PodDisruptionBudgetLister helps list PodDisruptionBudgets.
type PodDisruptionBudgetLister interface {
	// List lists all PodDisruptionBudgets in the indexer.
	List(selector labels.Selector) (ret []*policy.PodDisruptionBudget, err error)
	// PodDisruptionBudgets returns an object that can list and get PodDisruptionBudgets.
	PodDisruptionBudgets(namespace string) PodDisruptionBudgetNamespaceLister
	PodDisruptionBudgetListerExpansion
}

// podDisruptionBudgetLister implements the PodDisruptionBudgetLister interface.
type podDisruptionBudgetLister struct {
	indexer cache.Indexer
}

// NewPodDisruptionBudgetLister returns a new PodDisruptionBudgetLister.
func NewPodDisruptionBudgetLister(indexer cache.Indexer) PodDisruptionBudgetLister {
	return &podDisruptionBudgetLister{indexer: indexer}
}

// List lists all PodDisruptionBudgets in the indexer.
func (s *podDisruptionBudgetLister) List(selector labels.Selector) (ret []*policy.PodDisruptionBudget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*policy.PodDisruptionBudget))
	})
	return ret, err
}

// PodDisruptionBudgets returns an object that can list and get PodDisruptionBudgets.
func (s *podDisruptionBudgetLister) PodDisruptionBudgets(namespace string) PodDisruptionBudgetNamespaceLister {
	return podDisruptionBudgetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PodDisruptionBudgetNamespaceLister helps list and get PodDisruptionBudgets.
type PodDisruptionBudgetNamespaceLister interface {
	// List lists all PodDisruptionBudgets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*policy.PodDisruptionBudget, err error)
	// Get retrieves the PodDisruptionBudget from the indexer for a given namespace and name.
	Get(name string) (*policy.PodDisruptionBudget, error)
	PodDisruptionBudgetNamespaceListerExpansion
}

// podDisruptionBudgetNamespaceLister implements the PodDisruptionBudgetNamespaceLister
// interface.
type podDisruptionBudgetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PodDisruptionBudgets in the indexer for a given namespace.
func (s podDisruptionBudgetNamespaceLister) List(selector labels.Selector) (ret []*policy.PodDisruptionBudget, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*policy.PodDisruptionBudget))
	})
	return ret, err
}

// Get retrieves the PodDisruptionBudget from the indexer for a given namespace and name.
func (s podDisruptionBudgetNamespaceLister) Get(name string) (*policy.PodDisruptionBudget, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(policy.Resource("poddisruptionbudget"), name)
	}
	return obj.(*policy.PodDisruptionBudget), nil
}
