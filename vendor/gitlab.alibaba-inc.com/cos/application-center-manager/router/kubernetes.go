package router

import (
	"gitlab.alibaba-inc.com/cos/application-center-manager/pkg/apis/alibabacloud/v1beta1"
)

// KubernetesRouter manages Kubernetes services
type KubernetesRouter interface {
	// Initialize creates or updates the primary and canary services
	Initialize(canary *v1beta1.Canary) error
	// Reconcile creates or updates the main service
	Reconcile(canary *v1beta1.Canary) error
	// Revert router
	Finalize(canary *v1beta1.Canary) error
}
