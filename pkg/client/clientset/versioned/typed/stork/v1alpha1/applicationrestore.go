/*
Copyright 2018 Openstorage.org

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
	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	scheme "github.com/libopenstorage/stork/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApplicationRestoresGetter has a method to return a ApplicationRestoreInterface.
// A group's client should implement this interface.
type ApplicationRestoresGetter interface {
	ApplicationRestores(namespace string) ApplicationRestoreInterface
}

// ApplicationRestoreInterface has methods to work with ApplicationRestore resources.
type ApplicationRestoreInterface interface {
	Create(*v1alpha1.ApplicationRestore) (*v1alpha1.ApplicationRestore, error)
	Update(*v1alpha1.ApplicationRestore) (*v1alpha1.ApplicationRestore, error)
	UpdateStatus(*v1alpha1.ApplicationRestore) (*v1alpha1.ApplicationRestore, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApplicationRestore, error)
	List(opts v1.ListOptions) (*v1alpha1.ApplicationRestoreList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationRestore, err error)
	ApplicationRestoreExpansion
}

// applicationRestores implements ApplicationRestoreInterface
type applicationRestores struct {
	client rest.Interface
	ns     string
}

// newApplicationRestores returns a ApplicationRestores
func newApplicationRestores(c *StorkV1alpha1Client, namespace string) *applicationRestores {
	return &applicationRestores{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the applicationRestore, and returns the corresponding applicationRestore object, and an error if there is any.
func (c *applicationRestores) Get(name string, options v1.GetOptions) (result *v1alpha1.ApplicationRestore, err error) {
	result = &v1alpha1.ApplicationRestore{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationrestores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApplicationRestores that match those selectors.
func (c *applicationRestores) List(opts v1.ListOptions) (result *v1alpha1.ApplicationRestoreList, err error) {
	result = &v1alpha1.ApplicationRestoreList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested applicationRestores.
func (c *applicationRestores) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("applicationrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a applicationRestore and creates it.  Returns the server's representation of the applicationRestore, and an error, if there is any.
func (c *applicationRestores) Create(applicationRestore *v1alpha1.ApplicationRestore) (result *v1alpha1.ApplicationRestore, err error) {
	result = &v1alpha1.ApplicationRestore{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("applicationrestores").
		Body(applicationRestore).
		Do().
		Into(result)
	return
}

// Update takes the representation of a applicationRestore and updates it. Returns the server's representation of the applicationRestore, and an error, if there is any.
func (c *applicationRestores) Update(applicationRestore *v1alpha1.ApplicationRestore) (result *v1alpha1.ApplicationRestore, err error) {
	result = &v1alpha1.ApplicationRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationrestores").
		Name(applicationRestore.Name).
		Body(applicationRestore).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *applicationRestores) UpdateStatus(applicationRestore *v1alpha1.ApplicationRestore) (result *v1alpha1.ApplicationRestore, err error) {
	result = &v1alpha1.ApplicationRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationrestores").
		Name(applicationRestore.Name).
		SubResource("status").
		Body(applicationRestore).
		Do().
		Into(result)
	return
}

// Delete takes name of the applicationRestore and deletes it. Returns an error if one occurs.
func (c *applicationRestores) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationrestores").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *applicationRestores) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationrestores").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched applicationRestore.
func (c *applicationRestores) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationRestore, err error) {
	result = &v1alpha1.ApplicationRestore{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("applicationrestores").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
