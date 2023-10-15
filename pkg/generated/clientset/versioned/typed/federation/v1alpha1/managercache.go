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

	v1alpha1 "github.com/EdgeNet-project/edgenet/pkg/apis/federation/v1alpha1"
	scheme "github.com/EdgeNet-project/edgenet/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ManagerCachesGetter has a method to return a ManagerCacheInterface.
// A group's client should implement this interface.
type ManagerCachesGetter interface {
	ManagerCaches() ManagerCacheInterface
}

// ManagerCacheInterface has methods to work with ManagerCache resources.
type ManagerCacheInterface interface {
	Create(ctx context.Context, managerCache *v1alpha1.ManagerCache, opts v1.CreateOptions) (*v1alpha1.ManagerCache, error)
	Update(ctx context.Context, managerCache *v1alpha1.ManagerCache, opts v1.UpdateOptions) (*v1alpha1.ManagerCache, error)
	UpdateStatus(ctx context.Context, managerCache *v1alpha1.ManagerCache, opts v1.UpdateOptions) (*v1alpha1.ManagerCache, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ManagerCache, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ManagerCacheList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagerCache, err error)
	ManagerCacheExpansion
}

// managerCaches implements ManagerCacheInterface
type managerCaches struct {
	client rest.Interface
}

// newManagerCaches returns a ManagerCaches
func newManagerCaches(c *FederationV1alpha1Client) *managerCaches {
	return &managerCaches{
		client: c.RESTClient(),
	}
}

// Get takes name of the managerCache, and returns the corresponding managerCache object, and an error if there is any.
func (c *managerCaches) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagerCache, err error) {
	result = &v1alpha1.ManagerCache{}
	err = c.client.Get().
		Resource("managercaches").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ManagerCaches that match those selectors.
func (c *managerCaches) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagerCacheList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ManagerCacheList{}
	err = c.client.Get().
		Resource("managercaches").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested managerCaches.
func (c *managerCaches) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("managercaches").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a managerCache and creates it.  Returns the server's representation of the managerCache, and an error, if there is any.
func (c *managerCaches) Create(ctx context.Context, managerCache *v1alpha1.ManagerCache, opts v1.CreateOptions) (result *v1alpha1.ManagerCache, err error) {
	result = &v1alpha1.ManagerCache{}
	err = c.client.Post().
		Resource("managercaches").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managerCache).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a managerCache and updates it. Returns the server's representation of the managerCache, and an error, if there is any.
func (c *managerCaches) Update(ctx context.Context, managerCache *v1alpha1.ManagerCache, opts v1.UpdateOptions) (result *v1alpha1.ManagerCache, err error) {
	result = &v1alpha1.ManagerCache{}
	err = c.client.Put().
		Resource("managercaches").
		Name(managerCache.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managerCache).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *managerCaches) UpdateStatus(ctx context.Context, managerCache *v1alpha1.ManagerCache, opts v1.UpdateOptions) (result *v1alpha1.ManagerCache, err error) {
	result = &v1alpha1.ManagerCache{}
	err = c.client.Put().
		Resource("managercaches").
		Name(managerCache.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managerCache).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the managerCache and deletes it. Returns an error if one occurs.
func (c *managerCaches) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("managercaches").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *managerCaches) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("managercaches").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched managerCache.
func (c *managerCaches) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagerCache, err error) {
	result = &v1alpha1.ManagerCache{}
	err = c.client.Patch(pt).
		Resource("managercaches").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
