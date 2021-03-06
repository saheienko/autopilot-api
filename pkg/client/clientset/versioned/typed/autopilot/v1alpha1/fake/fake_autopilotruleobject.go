/*
Copyright 2019 Openstorage.org

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

package fake

import (
	v1alpha1 "github.com/libopenstorage/autopilot-api/pkg/apis/autopilot/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAutopilotRuleObjects implements AutopilotRuleObjectInterface
type FakeAutopilotRuleObjects struct {
	Fake *FakeAutopilotV1alpha1
}

var autopilotruleobjectsResource = schema.GroupVersionResource{Group: "autopilot.libopenstorage.org", Version: "v1alpha1", Resource: "autopilotruleobjects"}

var autopilotruleobjectsKind = schema.GroupVersionKind{Group: "autopilot.libopenstorage.org", Version: "v1alpha1", Kind: "AutopilotRuleObject"}

// Get takes name of the autopilotRuleObject, and returns the corresponding autopilotRuleObject object, and an error if there is any.
func (c *FakeAutopilotRuleObjects) Get(name string, options v1.GetOptions) (result *v1alpha1.AutopilotRuleObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(autopilotruleobjectsResource, name), &v1alpha1.AutopilotRuleObject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutopilotRuleObject), err
}

// List takes label and field selectors, and returns the list of AutopilotRuleObjects that match those selectors.
func (c *FakeAutopilotRuleObjects) List(opts v1.ListOptions) (result *v1alpha1.AutopilotRuleObjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(autopilotruleobjectsResource, autopilotruleobjectsKind, opts), &v1alpha1.AutopilotRuleObjectList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AutopilotRuleObjectList{ListMeta: obj.(*v1alpha1.AutopilotRuleObjectList).ListMeta}
	for _, item := range obj.(*v1alpha1.AutopilotRuleObjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested autopilotRuleObjects.
func (c *FakeAutopilotRuleObjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(autopilotruleobjectsResource, opts))
}

// Create takes the representation of a autopilotRuleObject and creates it.  Returns the server's representation of the autopilotRuleObject, and an error, if there is any.
func (c *FakeAutopilotRuleObjects) Create(autopilotRuleObject *v1alpha1.AutopilotRuleObject) (result *v1alpha1.AutopilotRuleObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(autopilotruleobjectsResource, autopilotRuleObject), &v1alpha1.AutopilotRuleObject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutopilotRuleObject), err
}

// Update takes the representation of a autopilotRuleObject and updates it. Returns the server's representation of the autopilotRuleObject, and an error, if there is any.
func (c *FakeAutopilotRuleObjects) Update(autopilotRuleObject *v1alpha1.AutopilotRuleObject) (result *v1alpha1.AutopilotRuleObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(autopilotruleobjectsResource, autopilotRuleObject), &v1alpha1.AutopilotRuleObject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutopilotRuleObject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAutopilotRuleObjects) UpdateStatus(autopilotRuleObject *v1alpha1.AutopilotRuleObject) (*v1alpha1.AutopilotRuleObject, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(autopilotruleobjectsResource, "status", autopilotRuleObject), &v1alpha1.AutopilotRuleObject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutopilotRuleObject), err
}

// Delete takes name of the autopilotRuleObject and deletes it. Returns an error if one occurs.
func (c *FakeAutopilotRuleObjects) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(autopilotruleobjectsResource, name), &v1alpha1.AutopilotRuleObject{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAutopilotRuleObjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(autopilotruleobjectsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AutopilotRuleObjectList{})
	return err
}

// Patch applies the patch and returns the patched autopilotRuleObject.
func (c *FakeAutopilotRuleObjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutopilotRuleObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(autopilotruleobjectsResource, name, pt, data, subresources...), &v1alpha1.AutopilotRuleObject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutopilotRuleObject), err
}
