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

package fake

import (
	"context"

	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApplicationBackupSchedules implements ApplicationBackupScheduleInterface
type FakeApplicationBackupSchedules struct {
	Fake *FakeStorkV1alpha1
	ns   string
}

var applicationbackupschedulesResource = schema.GroupVersionResource{Group: "stork.libopenstorage.org", Version: "v1alpha1", Resource: "applicationbackupschedules"}

var applicationbackupschedulesKind = schema.GroupVersionKind{Group: "stork.libopenstorage.org", Version: "v1alpha1", Kind: "ApplicationBackupSchedule"}

// Get takes name of the applicationBackupSchedule, and returns the corresponding applicationBackupSchedule object, and an error if there is any.
func (c *FakeApplicationBackupSchedules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApplicationBackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationbackupschedulesResource, c.ns, name), &v1alpha1.ApplicationBackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationBackupSchedule), err
}

// List takes label and field selectors, and returns the list of ApplicationBackupSchedules that match those selectors.
func (c *FakeApplicationBackupSchedules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApplicationBackupScheduleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationbackupschedulesResource, applicationbackupschedulesKind, c.ns, opts), &v1alpha1.ApplicationBackupScheduleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApplicationBackupScheduleList{ListMeta: obj.(*v1alpha1.ApplicationBackupScheduleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApplicationBackupScheduleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested applicationBackupSchedules.
func (c *FakeApplicationBackupSchedules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationbackupschedulesResource, c.ns, opts))

}

// Create takes the representation of a applicationBackupSchedule and creates it.  Returns the server's representation of the applicationBackupSchedule, and an error, if there is any.
func (c *FakeApplicationBackupSchedules) Create(ctx context.Context, applicationBackupSchedule *v1alpha1.ApplicationBackupSchedule, opts v1.CreateOptions) (result *v1alpha1.ApplicationBackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationbackupschedulesResource, c.ns, applicationBackupSchedule), &v1alpha1.ApplicationBackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationBackupSchedule), err
}

// Update takes the representation of a applicationBackupSchedule and updates it. Returns the server's representation of the applicationBackupSchedule, and an error, if there is any.
func (c *FakeApplicationBackupSchedules) Update(ctx context.Context, applicationBackupSchedule *v1alpha1.ApplicationBackupSchedule, opts v1.UpdateOptions) (result *v1alpha1.ApplicationBackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationbackupschedulesResource, c.ns, applicationBackupSchedule), &v1alpha1.ApplicationBackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationBackupSchedule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApplicationBackupSchedules) UpdateStatus(ctx context.Context, applicationBackupSchedule *v1alpha1.ApplicationBackupSchedule, opts v1.UpdateOptions) (*v1alpha1.ApplicationBackupSchedule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(applicationbackupschedulesResource, "status", c.ns, applicationBackupSchedule), &v1alpha1.ApplicationBackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationBackupSchedule), err
}

// Delete takes name of the applicationBackupSchedule and deletes it. Returns an error if one occurs.
func (c *FakeApplicationBackupSchedules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(applicationbackupschedulesResource, c.ns, name, opts), &v1alpha1.ApplicationBackupSchedule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplicationBackupSchedules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationbackupschedulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApplicationBackupScheduleList{})
	return err
}

// Patch applies the patch and returns the patched applicationBackupSchedule.
func (c *FakeApplicationBackupSchedules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApplicationBackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationbackupschedulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApplicationBackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationBackupSchedule), err
}
