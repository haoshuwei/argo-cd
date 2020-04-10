package router

import appv1beat1 "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/apis/alibabacloud/v1beta1"

const configAnnotation = "alibaba-cloud-appctl.kubernetes.io/original-configuration"
const kubectlAnnotation = "kubectl.kubernetes.io/last-applied-configuration"

type Interface interface {
	Reconcile(canary *appv1beat1.Canary) error
	SetRoutes(canary *appv1beat1.Canary, primaryWeight int, canaryWeight int, mirrored bool) error
	GetRoutes(canary *appv1beat1.Canary) (primaryWeight int, canaryWeight int, mirrored bool, err error)
	Finalize(canary *appv1beat1.Canary) error
}
