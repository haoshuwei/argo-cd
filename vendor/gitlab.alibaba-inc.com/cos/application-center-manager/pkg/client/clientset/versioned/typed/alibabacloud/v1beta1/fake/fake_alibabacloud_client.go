// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/client/clientset/versioned/typed/alibabacloud/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAppV1beta1 struct {
	*testing.Fake
}

func (c *FakeAppV1beta1) Canaries(namespace string) v1beta1.CanaryInterface {
	return &FakeCanaries{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAppV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
