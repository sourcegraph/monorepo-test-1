/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "github.com/sourcegraph/monorepo-test-1/kubernetes-3/federation/apis/federation/v1beta1"
)

// FakeClusters implements ClusterInterface
type FakeClusters struct {
	Fake *FakeFederationV1beta1
}

var clustersResource = schema.GroupVersionResource{Group: "federation", Version: "v1beta1", Resource: "clusters"}

func (c *FakeClusters) Create(cluster *v1beta1.Cluster) (result *v1beta1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clustersResource, cluster), &v1beta1.Cluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Cluster), err
}

func (c *FakeClusters) Update(cluster *v1beta1.Cluster) (result *v1beta1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clustersResource, cluster), &v1beta1.Cluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Cluster), err
}

func (c *FakeClusters) UpdateStatus(cluster *v1beta1.Cluster) (*v1beta1.Cluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clustersResource, "status", cluster), &v1beta1.Cluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Cluster), err
}

func (c *FakeClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clustersResource, name), &v1beta1.Cluster{})
	return err
}

func (c *FakeClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clustersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.ClusterList{})
	return err
}

func (c *FakeClusters) Get(name string, options v1.GetOptions) (result *v1beta1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clustersResource, name), &v1beta1.Cluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Cluster), err
}

func (c *FakeClusters) List(opts v1.ListOptions) (result *v1beta1.ClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clustersResource, opts), &v1beta1.ClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ClusterList{}
	for _, item := range obj.(*v1beta1.ClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusters.
func (c *FakeClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clustersResource, opts))
}

// Patch applies the patch and returns the patched cluster.
func (c *FakeClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clustersResource, name, data, subresources...), &v1beta1.Cluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Cluster), err
}
