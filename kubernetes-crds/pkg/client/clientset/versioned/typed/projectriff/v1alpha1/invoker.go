/*
 * Copyright 2018 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1alpha1

import (
	v1alpha1 "github.com/projectriff/riff/kubernetes-crds/pkg/apis/projectriff.io/v1alpha1"
	scheme "github.com/projectriff/riff/kubernetes-crds/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// InvokersGetter has a method to return a InvokerInterface.
// A group's client should implement this interface.
type InvokersGetter interface {
	Invokers(namespace string) InvokerInterface
}

// InvokerInterface has methods to work with Invoker resources.
type InvokerInterface interface {
	Create(*v1alpha1.Invoker) (*v1alpha1.Invoker, error)
	Update(*v1alpha1.Invoker) (*v1alpha1.Invoker, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Invoker, error)
	List(opts v1.ListOptions) (*v1alpha1.InvokerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Invoker, err error)
	InvokerExpansion
}

// invokers implements InvokerInterface
type invokers struct {
	client rest.Interface
	ns     string
}

// newInvokers returns a Invokers
func newInvokers(c *ProjectriffV1alpha1Client, namespace string) *invokers {
	return &invokers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the invoker, and returns the corresponding invoker object, and an error if there is any.
func (c *invokers) Get(name string, options v1.GetOptions) (result *v1alpha1.Invoker, err error) {
	result = &v1alpha1.Invoker{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("invokers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Invokers that match those selectors.
func (c *invokers) List(opts v1.ListOptions) (result *v1alpha1.InvokerList, err error) {
	result = &v1alpha1.InvokerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("invokers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested invokers.
func (c *invokers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("invokers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a invoker and creates it.  Returns the server's representation of the invoker, and an error, if there is any.
func (c *invokers) Create(invoker *v1alpha1.Invoker) (result *v1alpha1.Invoker, err error) {
	result = &v1alpha1.Invoker{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("invokers").
		Body(invoker).
		Do().
		Into(result)
	return
}

// Update takes the representation of a invoker and updates it. Returns the server's representation of the invoker, and an error, if there is any.
func (c *invokers) Update(invoker *v1alpha1.Invoker) (result *v1alpha1.Invoker, err error) {
	result = &v1alpha1.Invoker{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("invokers").
		Name(invoker.Name).
		Body(invoker).
		Do().
		Into(result)
	return
}

// Delete takes name of the invoker and deletes it. Returns an error if one occurs.
func (c *invokers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("invokers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *invokers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("invokers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched invoker.
func (c *invokers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Invoker, err error) {
	result = &v1alpha1.Invoker{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("invokers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
