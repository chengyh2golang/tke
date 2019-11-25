/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/versioned/scheme"
	v1 "tkestack.io/tke/api/platform/v1"
)

// PersistentEventsGetter has a method to return a PersistentEventInterface.
// A group's client should implement this interface.
type PersistentEventsGetter interface {
	PersistentEvents() PersistentEventInterface
}

// PersistentEventInterface has methods to work with PersistentEvent resources.
type PersistentEventInterface interface {
	Create(*v1.PersistentEvent) (*v1.PersistentEvent, error)
	Update(*v1.PersistentEvent) (*v1.PersistentEvent, error)
	UpdateStatus(*v1.PersistentEvent) (*v1.PersistentEvent, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v1.PersistentEvent, error)
	List(opts metav1.ListOptions) (*v1.PersistentEventList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PersistentEvent, err error)
	PersistentEventExpansion
}

// persistentEvents implements PersistentEventInterface
type persistentEvents struct {
	client rest.Interface
}

// newPersistentEvents returns a PersistentEvents
func newPersistentEvents(c *PlatformV1Client) *persistentEvents {
	return &persistentEvents{
		client: c.RESTClient(),
	}
}

// Get takes name of the persistentEvent, and returns the corresponding persistentEvent object, and an error if there is any.
func (c *persistentEvents) Get(name string, options metav1.GetOptions) (result *v1.PersistentEvent, err error) {
	result = &v1.PersistentEvent{}
	err = c.client.Get().
		Resource("persistentevents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PersistentEvents that match those selectors.
func (c *persistentEvents) List(opts metav1.ListOptions) (result *v1.PersistentEventList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PersistentEventList{}
	err = c.client.Get().
		Resource("persistentevents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested persistentEvents.
func (c *persistentEvents) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("persistentevents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a persistentEvent and creates it.  Returns the server's representation of the persistentEvent, and an error, if there is any.
func (c *persistentEvents) Create(persistentEvent *v1.PersistentEvent) (result *v1.PersistentEvent, err error) {
	result = &v1.PersistentEvent{}
	err = c.client.Post().
		Resource("persistentevents").
		Body(persistentEvent).
		Do().
		Into(result)
	return
}

// Update takes the representation of a persistentEvent and updates it. Returns the server's representation of the persistentEvent, and an error, if there is any.
func (c *persistentEvents) Update(persistentEvent *v1.PersistentEvent) (result *v1.PersistentEvent, err error) {
	result = &v1.PersistentEvent{}
	err = c.client.Put().
		Resource("persistentevents").
		Name(persistentEvent.Name).
		Body(persistentEvent).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *persistentEvents) UpdateStatus(persistentEvent *v1.PersistentEvent) (result *v1.PersistentEvent, err error) {
	result = &v1.PersistentEvent{}
	err = c.client.Put().
		Resource("persistentevents").
		Name(persistentEvent.Name).
		SubResource("status").
		Body(persistentEvent).
		Do().
		Into(result)
	return
}

// Delete takes name of the persistentEvent and deletes it. Returns an error if one occurs.
func (c *persistentEvents) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("persistentevents").
		Name(name).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched persistentEvent.
func (c *persistentEvents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PersistentEvent, err error) {
	result = &v1.PersistentEvent{}
	err = c.client.Patch(pt).
		Resource("persistentevents").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
