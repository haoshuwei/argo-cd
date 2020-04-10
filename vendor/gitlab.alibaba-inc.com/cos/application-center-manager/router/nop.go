package router

import (
	appv1beat1 "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/apis/alibabacloud/v1beta1"
)

// NopRouter no-operation router
type NopRouter struct {
}

func (*NopRouter) Reconcile(_ *appv1beat1.Canary) error {
	return nil
}

func (*NopRouter) SetRoutes(_ *appv1beat1.Canary, _ int, _ int, _ bool) error {
	return nil
}

func (*NopRouter) GetRoutes(canary *appv1beat1.Canary) (primaryWeight int, canaryWeight int, mirror bool, err error) {
	if canary.Status.Iterations > 0 {
		return 0, 100, false, nil
	}
	return 100, 0, false, nil
}

func (c *NopRouter) Finalize(_ *appv1beat1.Canary) error {
	return nil
}
