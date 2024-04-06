// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by lister-gen. DO NOT EDIT.

package v2alpha1

import (
	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CiliumBGPNodeConfigOverrideLister helps list CiliumBGPNodeConfigOverrides.
// All objects returned here must be treated as read-only.
type CiliumBGPNodeConfigOverrideLister interface {
	// List lists all CiliumBGPNodeConfigOverrides in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2alpha1.CiliumBGPNodeConfigOverride, err error)
	// Get retrieves the CiliumBGPNodeConfigOverride from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2alpha1.CiliumBGPNodeConfigOverride, error)
	CiliumBGPNodeConfigOverrideListerExpansion
}

// ciliumBGPNodeConfigOverrideLister implements the CiliumBGPNodeConfigOverrideLister interface.
type ciliumBGPNodeConfigOverrideLister struct {
	indexer cache.Indexer
}

// NewCiliumBGPNodeConfigOverrideLister returns a new CiliumBGPNodeConfigOverrideLister.
func NewCiliumBGPNodeConfigOverrideLister(indexer cache.Indexer) CiliumBGPNodeConfigOverrideLister {
	return &ciliumBGPNodeConfigOverrideLister{indexer: indexer}
}

// List lists all CiliumBGPNodeConfigOverrides in the indexer.
func (s *ciliumBGPNodeConfigOverrideLister) List(selector labels.Selector) (ret []*v2alpha1.CiliumBGPNodeConfigOverride, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2alpha1.CiliumBGPNodeConfigOverride))
	})
	return ret, err
}

// Get retrieves the CiliumBGPNodeConfigOverride from the index for a given name.
func (s *ciliumBGPNodeConfigOverrideLister) Get(name string) (*v2alpha1.CiliumBGPNodeConfigOverride, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2alpha1.Resource("ciliumbgpnodeconfigoverride"), name)
	}
	return obj.(*v2alpha1.CiliumBGPNodeConfigOverride), nil
}
