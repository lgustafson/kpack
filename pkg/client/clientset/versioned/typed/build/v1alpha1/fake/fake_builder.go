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
package fake

import (
	v1alpha1 "github.com/pivotal/build-service-system/pkg/apis/build/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBuilders implements BuilderInterface
type FakeBuilders struct {
	Fake *FakeBuildV1alpha1
	ns   string
}

var buildersResource = schema.GroupVersionResource{Group: "build.pivotal.io", Version: "v1alpha1", Resource: "builders"}

var buildersKind = schema.GroupVersionKind{Group: "build.pivotal.io", Version: "v1alpha1", Kind: "Builder"}

// Get takes name of the builder, and returns the corresponding builder object, and an error if there is any.
func (c *FakeBuilders) Get(name string, options v1.GetOptions) (result *v1alpha1.Builder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(buildersResource, c.ns, name), &v1alpha1.Builder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Builder), err
}

// List takes label and field selectors, and returns the list of Builders that match those selectors.
func (c *FakeBuilders) List(opts v1.ListOptions) (result *v1alpha1.BuilderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(buildersResource, buildersKind, c.ns, opts), &v1alpha1.BuilderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BuilderList{ListMeta: obj.(*v1alpha1.BuilderList).ListMeta}
	for _, item := range obj.(*v1alpha1.BuilderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested builders.
func (c *FakeBuilders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(buildersResource, c.ns, opts))

}

// Create takes the representation of a builder and creates it.  Returns the server's representation of the builder, and an error, if there is any.
func (c *FakeBuilders) Create(builder *v1alpha1.Builder) (result *v1alpha1.Builder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(buildersResource, c.ns, builder), &v1alpha1.Builder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Builder), err
}

// Update takes the representation of a builder and updates it. Returns the server's representation of the builder, and an error, if there is any.
func (c *FakeBuilders) Update(builder *v1alpha1.Builder) (result *v1alpha1.Builder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(buildersResource, c.ns, builder), &v1alpha1.Builder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Builder), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBuilders) UpdateStatus(builder *v1alpha1.Builder) (*v1alpha1.Builder, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(buildersResource, "status", c.ns, builder), &v1alpha1.Builder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Builder), err
}

// Delete takes name of the builder and deletes it. Returns an error if one occurs.
func (c *FakeBuilders) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(buildersResource, c.ns, name), &v1alpha1.Builder{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBuilders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(buildersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BuilderList{})
	return err
}

// Patch applies the patch and returns the patched builder.
func (c *FakeBuilders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Builder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(buildersResource, c.ns, name, data, subresources...), &v1alpha1.Builder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Builder), err
}
