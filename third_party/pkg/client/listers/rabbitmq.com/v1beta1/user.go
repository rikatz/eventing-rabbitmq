/*
Copyright 2020 The Knative Authors

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
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "knative.dev/eventing-rabbitmq/third_party/pkg/apis/rabbitmq.com/v1beta1"
)

// UserLister helps list Users.
// All objects returned here must be treated as read-only.
type UserLister interface {
	// List lists all Users in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.User, err error)
	// Users returns an object that can list and get Users.
	Users(namespace string) UserNamespaceLister
	UserListerExpansion
}

// userLister implements the UserLister interface.
type userLister struct {
	indexer cache.Indexer
}

// NewUserLister returns a new UserLister.
func NewUserLister(indexer cache.Indexer) UserLister {
	return &userLister{indexer: indexer}
}

// List lists all Users in the indexer.
func (s *userLister) List(selector labels.Selector) (ret []*v1beta1.User, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.User))
	})
	return ret, err
}

// Users returns an object that can list and get Users.
func (s *userLister) Users(namespace string) UserNamespaceLister {
	return userNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// UserNamespaceLister helps list and get Users.
// All objects returned here must be treated as read-only.
type UserNamespaceLister interface {
	// List lists all Users in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.User, err error)
	// Get retrieves the User from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.User, error)
	UserNamespaceListerExpansion
}

// userNamespaceLister implements the UserNamespaceLister
// interface.
type userNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Users in the indexer for a given namespace.
func (s userNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.User, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.User))
	})
	return ret, err
}

// Get retrieves the User from the indexer for a given namespace and name.
func (s userNamespaceLister) Get(name string) (*v1beta1.User, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("user"), name)
	}
	return obj.(*v1beta1.User), nil
}
