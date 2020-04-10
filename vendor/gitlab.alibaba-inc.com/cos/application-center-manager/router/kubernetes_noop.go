package router

import (
	"gitlab.alibaba-inc.com/cos/application-center-manager/pkg/apis/alibabacloud/v1beta1"
)

// KubernetesNoopRouter manages nothing. This is useful when one uses Flagger for progressive delivery of
// services that are not load-balanced by a Kubernetes service
type KubernetesNoopRouter struct {
}

func (c *KubernetesNoopRouter) Initialize(_ *v1beta1.Canary) error {
	return nil
}

func (c *KubernetesNoopRouter) Reconcile(_ *v1beta1.Canary) error {
	return nil
}

func (c *KubernetesNoopRouter) Finalize(_ *v1beta1.Canary) error {
	return nil
}
