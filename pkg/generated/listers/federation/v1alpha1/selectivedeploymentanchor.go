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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/EdgeNet-project/edgenet/pkg/apis/federation/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SelectiveDeploymentAnchorLister helps list SelectiveDeploymentAnchors.
// All objects returned here must be treated as read-only.
type SelectiveDeploymentAnchorLister interface {
	// List lists all SelectiveDeploymentAnchors in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SelectiveDeploymentAnchor, err error)
	// SelectiveDeploymentAnchors returns an object that can list and get SelectiveDeploymentAnchors.
	SelectiveDeploymentAnchors(namespace string) SelectiveDeploymentAnchorNamespaceLister
	SelectiveDeploymentAnchorListerExpansion
}

// selectiveDeploymentAnchorLister implements the SelectiveDeploymentAnchorLister interface.
type selectiveDeploymentAnchorLister struct {
	indexer cache.Indexer
}

// NewSelectiveDeploymentAnchorLister returns a new SelectiveDeploymentAnchorLister.
func NewSelectiveDeploymentAnchorLister(indexer cache.Indexer) SelectiveDeploymentAnchorLister {
	return &selectiveDeploymentAnchorLister{indexer: indexer}
}

// List lists all SelectiveDeploymentAnchors in the indexer.
func (s *selectiveDeploymentAnchorLister) List(selector labels.Selector) (ret []*v1alpha1.SelectiveDeploymentAnchor, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SelectiveDeploymentAnchor))
	})
	return ret, err
}

// SelectiveDeploymentAnchors returns an object that can list and get SelectiveDeploymentAnchors.
func (s *selectiveDeploymentAnchorLister) SelectiveDeploymentAnchors(namespace string) SelectiveDeploymentAnchorNamespaceLister {
	return selectiveDeploymentAnchorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SelectiveDeploymentAnchorNamespaceLister helps list and get SelectiveDeploymentAnchors.
// All objects returned here must be treated as read-only.
type SelectiveDeploymentAnchorNamespaceLister interface {
	// List lists all SelectiveDeploymentAnchors in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SelectiveDeploymentAnchor, err error)
	// Get retrieves the SelectiveDeploymentAnchor from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SelectiveDeploymentAnchor, error)
	SelectiveDeploymentAnchorNamespaceListerExpansion
}

// selectiveDeploymentAnchorNamespaceLister implements the SelectiveDeploymentAnchorNamespaceLister
// interface.
type selectiveDeploymentAnchorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SelectiveDeploymentAnchors in the indexer for a given namespace.
func (s selectiveDeploymentAnchorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SelectiveDeploymentAnchor, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SelectiveDeploymentAnchor))
	})
	return ret, err
}

// Get retrieves the SelectiveDeploymentAnchor from the indexer for a given namespace and name.
func (s selectiveDeploymentAnchorNamespaceLister) Get(name string) (*v1alpha1.SelectiveDeploymentAnchor, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("selectivedeploymentanchor"), name)
	}
	return obj.(*v1alpha1.SelectiveDeploymentAnchor), nil
}
