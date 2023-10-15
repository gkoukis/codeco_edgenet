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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/EdgeNet-project/edgenet/pkg/apis/core/v1alpha1"
	scheme "github.com/EdgeNet-project/edgenet/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SliceClaimsGetter has a method to return a SliceClaimInterface.
// A group's client should implement this interface.
type SliceClaimsGetter interface {
	SliceClaims(namespace string) SliceClaimInterface
}

// SliceClaimInterface has methods to work with SliceClaim resources.
type SliceClaimInterface interface {
	Create(ctx context.Context, sliceClaim *v1alpha1.SliceClaim, opts v1.CreateOptions) (*v1alpha1.SliceClaim, error)
	Update(ctx context.Context, sliceClaim *v1alpha1.SliceClaim, opts v1.UpdateOptions) (*v1alpha1.SliceClaim, error)
	UpdateStatus(ctx context.Context, sliceClaim *v1alpha1.SliceClaim, opts v1.UpdateOptions) (*v1alpha1.SliceClaim, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SliceClaim, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SliceClaimList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SliceClaim, err error)
	SliceClaimExpansion
}

// sliceClaims implements SliceClaimInterface
type sliceClaims struct {
	client rest.Interface
	ns     string
}

// newSliceClaims returns a SliceClaims
func newSliceClaims(c *CoreV1alpha1Client, namespace string) *sliceClaims {
	return &sliceClaims{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sliceClaim, and returns the corresponding sliceClaim object, and an error if there is any.
func (c *sliceClaims) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SliceClaim, err error) {
	result = &v1alpha1.SliceClaim{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sliceclaims").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SliceClaims that match those selectors.
func (c *sliceClaims) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SliceClaimList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SliceClaimList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sliceclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sliceClaims.
func (c *sliceClaims) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sliceclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sliceClaim and creates it.  Returns the server's representation of the sliceClaim, and an error, if there is any.
func (c *sliceClaims) Create(ctx context.Context, sliceClaim *v1alpha1.SliceClaim, opts v1.CreateOptions) (result *v1alpha1.SliceClaim, err error) {
	result = &v1alpha1.SliceClaim{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sliceclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sliceClaim).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sliceClaim and updates it. Returns the server's representation of the sliceClaim, and an error, if there is any.
func (c *sliceClaims) Update(ctx context.Context, sliceClaim *v1alpha1.SliceClaim, opts v1.UpdateOptions) (result *v1alpha1.SliceClaim, err error) {
	result = &v1alpha1.SliceClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sliceclaims").
		Name(sliceClaim.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sliceClaim).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *sliceClaims) UpdateStatus(ctx context.Context, sliceClaim *v1alpha1.SliceClaim, opts v1.UpdateOptions) (result *v1alpha1.SliceClaim, err error) {
	result = &v1alpha1.SliceClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sliceclaims").
		Name(sliceClaim.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sliceClaim).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sliceClaim and deletes it. Returns an error if one occurs.
func (c *sliceClaims) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sliceclaims").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sliceClaims) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sliceclaims").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sliceClaim.
func (c *sliceClaims) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SliceClaim, err error) {
	result = &v1alpha1.SliceClaim{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sliceclaims").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
