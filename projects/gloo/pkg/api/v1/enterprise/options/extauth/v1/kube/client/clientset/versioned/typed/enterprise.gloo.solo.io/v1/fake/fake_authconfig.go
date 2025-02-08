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

package fake

import (
	"context"

	enterprisegloosoloiov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1/kube/apis/enterprise.gloo.solo.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAuthConfigs implements AuthConfigInterface
type FakeAuthConfigs struct {
	Fake *FakeEnterpriseV1
	ns   string
}

var authconfigsResource = schema.GroupVersionResource{Group: "enterprise.gloo.solo.io", Version: "v1", Resource: "authconfigs"}

var authconfigsKind = schema.GroupVersionKind{Group: "enterprise.gloo.solo.io", Version: "v1", Kind: "AuthConfig"}

// Get takes name of the authConfig, and returns the corresponding authConfig object, and an error if there is any.
func (c *FakeAuthConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *enterprisegloosoloiov1.AuthConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(authconfigsResource, c.ns, name), &enterprisegloosoloiov1.AuthConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*enterprisegloosoloiov1.AuthConfig), err
}

// List takes label and field selectors, and returns the list of AuthConfigs that match those selectors.
func (c *FakeAuthConfigs) List(ctx context.Context, opts v1.ListOptions) (result *enterprisegloosoloiov1.AuthConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(authconfigsResource, authconfigsKind, c.ns, opts), &enterprisegloosoloiov1.AuthConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &enterprisegloosoloiov1.AuthConfigList{ListMeta: obj.(*enterprisegloosoloiov1.AuthConfigList).ListMeta}
	for _, item := range obj.(*enterprisegloosoloiov1.AuthConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested authConfigs.
func (c *FakeAuthConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(authconfigsResource, c.ns, opts))

}

// Create takes the representation of a authConfig and creates it.  Returns the server's representation of the authConfig, and an error, if there is any.
func (c *FakeAuthConfigs) Create(ctx context.Context, authConfig *enterprisegloosoloiov1.AuthConfig, opts v1.CreateOptions) (result *enterprisegloosoloiov1.AuthConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(authconfigsResource, c.ns, authConfig), &enterprisegloosoloiov1.AuthConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*enterprisegloosoloiov1.AuthConfig), err
}

// Update takes the representation of a authConfig and updates it. Returns the server's representation of the authConfig, and an error, if there is any.
func (c *FakeAuthConfigs) Update(ctx context.Context, authConfig *enterprisegloosoloiov1.AuthConfig, opts v1.UpdateOptions) (result *enterprisegloosoloiov1.AuthConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(authconfigsResource, c.ns, authConfig), &enterprisegloosoloiov1.AuthConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*enterprisegloosoloiov1.AuthConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAuthConfigs) UpdateStatus(ctx context.Context, authConfig *enterprisegloosoloiov1.AuthConfig, opts v1.UpdateOptions) (*enterprisegloosoloiov1.AuthConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(authconfigsResource, "status", c.ns, authConfig), &enterprisegloosoloiov1.AuthConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*enterprisegloosoloiov1.AuthConfig), err
}

// Delete takes name of the authConfig and deletes it. Returns an error if one occurs.
func (c *FakeAuthConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(authconfigsResource, c.ns, name), &enterprisegloosoloiov1.AuthConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAuthConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(authconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &enterprisegloosoloiov1.AuthConfigList{})
	return err
}

// Patch applies the patch and returns the patched authConfig.
func (c *FakeAuthConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *enterprisegloosoloiov1.AuthConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(authconfigsResource, c.ns, name, pt, data, subresources...), &enterprisegloosoloiov1.AuthConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*enterprisegloosoloiov1.AuthConfig), err
}
