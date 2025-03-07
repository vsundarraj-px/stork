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

// FakeMigrations implements MigrationInterface
type FakeMigrations struct {
	Fake *FakeStorkV1alpha1
	ns   string
}

var migrationsResource = schema.GroupVersionResource{Group: "stork.libopenstorage.org", Version: "v1alpha1", Resource: "migrations"}

var migrationsKind = schema.GroupVersionKind{Group: "stork.libopenstorage.org", Version: "v1alpha1", Kind: "Migration"}

// Get takes name of the migration, and returns the corresponding migration object, and an error if there is any.
func (c *FakeMigrations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Migration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(migrationsResource, c.ns, name), &v1alpha1.Migration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Migration), err
}

// List takes label and field selectors, and returns the list of Migrations that match those selectors.
func (c *FakeMigrations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MigrationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(migrationsResource, migrationsKind, c.ns, opts), &v1alpha1.MigrationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MigrationList{ListMeta: obj.(*v1alpha1.MigrationList).ListMeta}
	for _, item := range obj.(*v1alpha1.MigrationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested migrations.
func (c *FakeMigrations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(migrationsResource, c.ns, opts))

}

// Create takes the representation of a migration and creates it.  Returns the server's representation of the migration, and an error, if there is any.
func (c *FakeMigrations) Create(ctx context.Context, migration *v1alpha1.Migration, opts v1.CreateOptions) (result *v1alpha1.Migration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(migrationsResource, c.ns, migration), &v1alpha1.Migration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Migration), err
}

// Update takes the representation of a migration and updates it. Returns the server's representation of the migration, and an error, if there is any.
func (c *FakeMigrations) Update(ctx context.Context, migration *v1alpha1.Migration, opts v1.UpdateOptions) (result *v1alpha1.Migration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(migrationsResource, c.ns, migration), &v1alpha1.Migration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Migration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMigrations) UpdateStatus(ctx context.Context, migration *v1alpha1.Migration, opts v1.UpdateOptions) (*v1alpha1.Migration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(migrationsResource, "status", c.ns, migration), &v1alpha1.Migration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Migration), err
}

// Delete takes name of the migration and deletes it. Returns an error if one occurs.
func (c *FakeMigrations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(migrationsResource, c.ns, name, opts), &v1alpha1.Migration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMigrations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(migrationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MigrationList{})
	return err
}

// Patch applies the patch and returns the patched migration.
func (c *FakeMigrations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Migration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(migrationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Migration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Migration), err
}
