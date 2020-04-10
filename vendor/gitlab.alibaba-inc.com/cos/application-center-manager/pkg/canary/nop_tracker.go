package canary

import (
	appv1beat1 "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/apis/alibabacloud/v1beta1"
	corev1 "k8s.io/api/core/v1"
)

// NopTracker no-operation tracker
type NopTracker struct{}

func (nt *NopTracker) GetTargetConfigs(*appv1beat1.Canary) (map[string]ConfigRef, error) {
	res := make(map[string]ConfigRef)
	return res, nil
}

func (nt *NopTracker) GetConfigRefs(*appv1beat1.Canary) (*map[string]string, error) {
	return nil, nil
}

func (nt *NopTracker) HasConfigChanged(*appv1beat1.Canary) (bool, error) {
	return false, nil
}

func (nt *NopTracker) CreatePrimaryConfigs(*appv1beat1.Canary, map[string]ConfigRef) error {
	return nil
}

func (nt *NopTracker) ApplyPrimaryConfigs(spec corev1.PodSpec, _ map[string]ConfigRef) corev1.PodSpec {
	return spec
}
