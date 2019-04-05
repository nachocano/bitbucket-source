// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/nachocano/bitbucket-source/pkg/client/clientset/versioned/typed/sources/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSourcesV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSourcesV1alpha1) BitBucketSources(namespace string) v1alpha1.BitBucketSourceInterface {
	return &FakeBitBucketSources{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSourcesV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
