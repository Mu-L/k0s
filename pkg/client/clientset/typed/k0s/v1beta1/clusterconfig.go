// SPDX-FileCopyrightText: k0s authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"

	k0sv1beta1 "github.com/k0sproject/k0s/pkg/apis/k0s/v1beta1"
	scheme "github.com/k0sproject/k0s/pkg/client/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ClusterConfigsGetter has a method to return a ClusterConfigInterface.
// A group's client should implement this interface.
type ClusterConfigsGetter interface {
	ClusterConfigs(namespace string) ClusterConfigInterface
}

// ClusterConfigInterface has methods to work with ClusterConfig resources.
type ClusterConfigInterface interface {
	Create(ctx context.Context, clusterConfig *k0sv1beta1.ClusterConfig, opts v1.CreateOptions) (*k0sv1beta1.ClusterConfig, error)
	Update(ctx context.Context, clusterConfig *k0sv1beta1.ClusterConfig, opts v1.UpdateOptions) (*k0sv1beta1.ClusterConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*k0sv1beta1.ClusterConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*k0sv1beta1.ClusterConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	ClusterConfigExpansion
}

// clusterConfigs implements ClusterConfigInterface
type clusterConfigs struct {
	*gentype.ClientWithList[*k0sv1beta1.ClusterConfig, *k0sv1beta1.ClusterConfigList]
}

// newClusterConfigs returns a ClusterConfigs
func newClusterConfigs(c *K0sV1beta1Client, namespace string) *clusterConfigs {
	return &clusterConfigs{
		gentype.NewClientWithList[*k0sv1beta1.ClusterConfig, *k0sv1beta1.ClusterConfigList](
			"clusterconfigs",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *k0sv1beta1.ClusterConfig { return &k0sv1beta1.ClusterConfig{} },
			func() *k0sv1beta1.ClusterConfigList { return &k0sv1beta1.ClusterConfigList{} },
		),
	}
}
