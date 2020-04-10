package canary

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	flaggerv1 "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/apis/alibabacloud/v1beta1"
)

// SyncStatus encodes the canary pod spec and updates the canary status
func (c *DeploymentController) SyncStatus(cd *flaggerv1.Canary, status flaggerv1.CanaryStatus) error {
	dep, err := c.kubeClient.AppsV1().Deployments(cd.Namespace).Get(cd.Spec.TargetRef.Name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("deployment %s.%s get query error: %w", cd.Spec.TargetRef.Name, cd.Namespace, err)
	}

	configs, err := c.configTracker.GetConfigRefs(cd)
	if err != nil {
		return fmt.Errorf("GetConfigRefs failed: %w", err)
	}

	return syncCanaryStatus(c.appClient, cd, status, dep.Spec.Template, func(cdCopy *flaggerv1.Canary) {
		cdCopy.Status.TrackedConfigs = configs
	})
}

// SetStatusFailedChecks updates the canary failed checks counter
func (c *DeploymentController) SetStatusFailedChecks(cd *flaggerv1.Canary, val int) error {
	return setStatusFailedChecks(c.appClient, cd, val)
}

// SetStatusWeight updates the canary status weight value
func (c *DeploymentController) SetStatusWeight(cd *flaggerv1.Canary, val int) error {
	return setStatusWeight(c.appClient, cd, val)
}

// SetStatusIterations updates the canary status iterations value
func (c *DeploymentController) SetStatusIterations(cd *flaggerv1.Canary, val int) error {
	return setStatusIterations(c.appClient, cd, val)
}

// SetStatusPhase updates the canary status phase
func (c *DeploymentController) SetStatusPhase(cd *flaggerv1.Canary, phase flaggerv1.CanaryPhase) error {
	return setStatusPhase(c.appClient, cd, phase)
}
