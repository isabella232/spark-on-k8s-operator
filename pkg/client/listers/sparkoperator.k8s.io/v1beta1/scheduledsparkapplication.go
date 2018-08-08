// Code generated by k8s code-generator DO NOT EDIT.

/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "k8s.io/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1beta1"
)

// ScheduledSparkApplicationLister helps list ScheduledSparkApplications.
type ScheduledSparkApplicationLister interface {
	// List lists all ScheduledSparkApplications in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.ScheduledSparkApplication, err error)
	// ScheduledSparkApplications returns an object that can list and get ScheduledSparkApplications.
	ScheduledSparkApplications(namespace string) ScheduledSparkApplicationNamespaceLister
	ScheduledSparkApplicationListerExpansion
}

// scheduledSparkApplicationLister implements the ScheduledSparkApplicationLister interface.
type scheduledSparkApplicationLister struct {
	indexer cache.Indexer
}

// NewScheduledSparkApplicationLister returns a new ScheduledSparkApplicationLister.
func NewScheduledSparkApplicationLister(indexer cache.Indexer) ScheduledSparkApplicationLister {
	return &scheduledSparkApplicationLister{indexer: indexer}
}

// List lists all ScheduledSparkApplications in the indexer.
func (s *scheduledSparkApplicationLister) List(selector labels.Selector) (ret []*v1beta1.ScheduledSparkApplication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ScheduledSparkApplication))
	})
	return ret, err
}

// ScheduledSparkApplications returns an object that can list and get ScheduledSparkApplications.
func (s *scheduledSparkApplicationLister) ScheduledSparkApplications(namespace string) ScheduledSparkApplicationNamespaceLister {
	return scheduledSparkApplicationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ScheduledSparkApplicationNamespaceLister helps list and get ScheduledSparkApplications.
type ScheduledSparkApplicationNamespaceLister interface {
	// List lists all ScheduledSparkApplications in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.ScheduledSparkApplication, err error)
	// Get retrieves the ScheduledSparkApplication from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.ScheduledSparkApplication, error)
	ScheduledSparkApplicationNamespaceListerExpansion
}

// scheduledSparkApplicationNamespaceLister implements the ScheduledSparkApplicationNamespaceLister
// interface.
type scheduledSparkApplicationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ScheduledSparkApplications in the indexer for a given namespace.
func (s scheduledSparkApplicationNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.ScheduledSparkApplication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ScheduledSparkApplication))
	})
	return ret, err
}

// Get retrieves the ScheduledSparkApplication from the indexer for a given namespace and name.
func (s scheduledSparkApplicationNamespaceLister) Get(name string) (*v1beta1.ScheduledSparkApplication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("scheduledsparkapplication"), name)
	}
	return obj.(*v1beta1.ScheduledSparkApplication), nil
}
