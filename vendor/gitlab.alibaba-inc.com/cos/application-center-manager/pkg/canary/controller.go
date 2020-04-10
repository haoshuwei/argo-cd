package canary

import (
	appv1beta1 "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/apis/alibabacloud/v1beta1"
)

type Controller interface {
	IsPrimaryReady(canary *appv1beta1.Canary) error
	IsCanaryReady(canary *appv1beta1.Canary) (bool, error)
	GetMetadata(canary *appv1beta1.Canary) (string, map[string]int32, error)
	SyncStatus(canary *appv1beta1.Canary, status appv1beta1.CanaryStatus) error
	SetStatusFailedChecks(canary *appv1beta1.Canary, val int) error
	SetStatusWeight(canary *appv1beta1.Canary, val int) error
	SetStatusIterations(canary *appv1beta1.Canary, val int) error
	SetStatusPhase(canary *appv1beta1.Canary, phase appv1beta1.CanaryPhase) error
	Initialize(canary *appv1beta1.Canary) error
	Promote(canary *appv1beta1.Canary) error
	HasTargetChanged(canary *appv1beta1.Canary) (bool, error)
	HaveDependenciesChanged(canary *appv1beta1.Canary) (bool, error)
	ScaleToZero(canary *appv1beta1.Canary) error
	ScaleFromZero(canary *appv1beta1.Canary) error
	Finalize(canary *appv1beta1.Canary) error
}
