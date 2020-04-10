package canary

import (
	appv1beat1 "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/apis/alibabacloud/v1beta1"
	corev1 "k8s.io/api/core/v1"
)

type Tracker interface {
	GetTargetConfigs(cd *appv1beat1.Canary) (map[string]ConfigRef, error)
	GetConfigRefs(cd *appv1beat1.Canary) (*map[string]string, error)
	HasConfigChanged(cd *appv1beat1.Canary) (bool, error)
	CreatePrimaryConfigs(cd *appv1beat1.Canary, refs map[string]ConfigRef) error
	ApplyPrimaryConfigs(spec corev1.PodSpec, refs map[string]ConfigRef) corev1.PodSpec
}
