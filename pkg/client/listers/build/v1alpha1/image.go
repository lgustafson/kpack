/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package v1alpha1

import (
	v1alpha1 "github.com/pivotal/build-service-system/pkg/apis/build/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ImageLister helps list Images.
type ImageLister interface {
	// List lists all Images in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Image, err error)
	// Images returns an object that can list and get Images.
	Images(namespace string) ImageNamespaceLister
	ImageListerExpansion
}

// imageLister implements the ImageLister interface.
type imageLister struct {
	indexer cache.Indexer
}

// NewImageLister returns a new ImageLister.
func NewImageLister(indexer cache.Indexer) ImageLister {
	return &imageLister{indexer: indexer}
}

// List lists all Images in the indexer.
func (s *imageLister) List(selector labels.Selector) (ret []*v1alpha1.Image, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Image))
	})
	return ret, err
}

// Images returns an object that can list and get Images.
func (s *imageLister) Images(namespace string) ImageNamespaceLister {
	return imageNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ImageNamespaceLister helps list and get Images.
type ImageNamespaceLister interface {
	// List lists all Images in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Image, err error)
	// Get retrieves the Image from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Image, error)
	ImageNamespaceListerExpansion
}

// imageNamespaceLister implements the ImageNamespaceLister
// interface.
type imageNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Images in the indexer for a given namespace.
func (s imageNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Image, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Image))
	})
	return ret, err
}

// Get retrieves the Image from the indexer for a given namespace and name.
func (s imageNamespaceLister) Get(name string) (*v1alpha1.Image, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("image"), name)
	}
	return obj.(*v1alpha1.Image), nil
}
