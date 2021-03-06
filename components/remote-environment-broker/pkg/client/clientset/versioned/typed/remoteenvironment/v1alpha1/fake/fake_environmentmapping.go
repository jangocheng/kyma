// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kyma-project/kyma/components/remote-environment-broker/pkg/apis/remoteenvironment/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEnvironmentMappings implements EnvironmentMappingInterface
type FakeEnvironmentMappings struct {
	Fake *FakeRemoteenvironmentV1alpha1
	ns   string
}

var environmentmappingsResource = schema.GroupVersionResource{Group: "remoteenvironment.kyma.cx", Version: "v1alpha1", Resource: "environmentmappings"}

var environmentmappingsKind = schema.GroupVersionKind{Group: "remoteenvironment.kyma.cx", Version: "v1alpha1", Kind: "EnvironmentMapping"}

// Get takes name of the environmentMapping, and returns the corresponding environmentMapping object, and an error if there is any.
func (c *FakeEnvironmentMappings) Get(name string, options v1.GetOptions) (result *v1alpha1.EnvironmentMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(environmentmappingsResource, c.ns, name), &v1alpha1.EnvironmentMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EnvironmentMapping), err
}

// List takes label and field selectors, and returns the list of EnvironmentMappings that match those selectors.
func (c *FakeEnvironmentMappings) List(opts v1.ListOptions) (result *v1alpha1.EnvironmentMappingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(environmentmappingsResource, environmentmappingsKind, c.ns, opts), &v1alpha1.EnvironmentMappingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EnvironmentMappingList{}
	for _, item := range obj.(*v1alpha1.EnvironmentMappingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested environmentMappings.
func (c *FakeEnvironmentMappings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(environmentmappingsResource, c.ns, opts))

}

// Create takes the representation of a environmentMapping and creates it.  Returns the server's representation of the environmentMapping, and an error, if there is any.
func (c *FakeEnvironmentMappings) Create(environmentMapping *v1alpha1.EnvironmentMapping) (result *v1alpha1.EnvironmentMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(environmentmappingsResource, c.ns, environmentMapping), &v1alpha1.EnvironmentMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EnvironmentMapping), err
}

// Update takes the representation of a environmentMapping and updates it. Returns the server's representation of the environmentMapping, and an error, if there is any.
func (c *FakeEnvironmentMappings) Update(environmentMapping *v1alpha1.EnvironmentMapping) (result *v1alpha1.EnvironmentMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(environmentmappingsResource, c.ns, environmentMapping), &v1alpha1.EnvironmentMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EnvironmentMapping), err
}

// Delete takes name of the environmentMapping and deletes it. Returns an error if one occurs.
func (c *FakeEnvironmentMappings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(environmentmappingsResource, c.ns, name), &v1alpha1.EnvironmentMapping{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEnvironmentMappings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(environmentmappingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EnvironmentMappingList{})
	return err
}

// Patch applies the patch and returns the patched environmentMapping.
func (c *FakeEnvironmentMappings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EnvironmentMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(environmentmappingsResource, c.ns, name, data, subresources...), &v1alpha1.EnvironmentMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EnvironmentMapping), err
}
