package canary

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
	"log"

	appv1beta1 "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/apis/alibabacloud/v1beta1"
	clientset "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/client/clientset/versioned"
)

func syncCanaryStatus(appClient clientset.Interface, cd *appv1beta1.Canary, status appv1beta1.CanaryStatus, canaryResource interface{}, setAll func(cdCopy *appv1beta1.Canary)) error {
	hash := computeHash(canaryResource)
	log.Print("syncCanaryStatus get in ")
	name, ns := cd.GetName(), cd.GetNamespace()
	err := retry.RetryOnConflict(retry.DefaultBackoff, func() (err error) {
		cd, err = appClient.AppV1beta1().Canaries(ns).Get(name, metav1.GetOptions{})
		if err != nil {
			return fmt.Errorf("canary %s.%s get query failed: %w", name, ns, err)
		}

		log.Print("syncCanaryStatus  123123123123")
		cdCopy := cd.DeepCopy()
		if status.Phase == appv1beta1.CanaryPhaseBegin {
			cdCopy.Status.Phase = appv1beta1.CanaryPhaseProgressing
		} else {
			cdCopy.Status.Phase = status.Phase
		}

		log.Printf("cd.ResourceVersion %s", cd.ResourceVersion)
		cdCopy.Status.CanaryWeight = status.CanaryWeight
		cdCopy.Status.FailedChecks = status.FailedChecks
		cdCopy.Status.Iterations = status.Iterations
		cdCopy.Status.LastAppliedSpec = hash
		cdCopy.Status.LastTransitionTime = metav1.Now()
		setAll(cdCopy)

		if ok, conditions := MakeStatusConditions(cd, status.Phase); ok {
			cdCopy.Status.Conditions = conditions
		}

		if err = updateStatusWithUpgrade(appClient, cdCopy); err != nil {
			return fmt.Errorf("updateStatusWithUpgrade failed: %w", err)
		}
		return
	})

	if err != nil {
		return fmt.Errorf("failed after retries: %w", err)
	}
	return nil
}

func setStatusFailedChecks(appClient clientset.Interface, cd *appv1beta1.Canary, val int) error {
	firstTry := true
	name, ns := cd.GetName(), cd.GetNamespace()
	err := retry.RetryOnConflict(retry.DefaultBackoff, func() (err error) {
		if !firstTry {
			cd, err = appClient.AppV1beta1().Canaries(ns).Get(name, metav1.GetOptions{})
			if err != nil {
				return fmt.Errorf("canary %s.%s get query failed: %w", name, ns, err)
			}
		}
		cdCopy := cd.DeepCopy()
		cdCopy.Status.FailedChecks = val
		cdCopy.Status.LastTransitionTime = metav1.Now()

		if err = updateStatusWithUpgrade(appClient, cdCopy); err != nil {
			return fmt.Errorf("updateStatusWithUpgrade failed: %w", err)
		}
		firstTry = false
		return
	})
	if err != nil {
		return fmt.Errorf("failed after retries: %w", err)
	}
	return nil
}

func setStatusWeight(appClient clientset.Interface, cd *appv1beta1.Canary, val int) error {
	firstTry := true
	name, ns := cd.GetName(), cd.GetNamespace()
	err := retry.RetryOnConflict(retry.DefaultBackoff, func() (err error) {
		if !firstTry {
			cd, err = appClient.AppV1beta1().Canaries(ns).Get(name, metav1.GetOptions{})
			if err != nil {
				return fmt.Errorf("canary %s.%s get query failed: %w", name, ns, err)
			}
		}
		cdCopy := cd.DeepCopy()
		cdCopy.Status.CanaryWeight = val
		cdCopy.Status.LastTransitionTime = metav1.Now()

		if err = updateStatusWithUpgrade(appClient, cdCopy); err != nil {
			return fmt.Errorf("updateStatusWithUpgrade failed: %w", err)
		}
		firstTry = false
		return
	})
	if err != nil {
		return fmt.Errorf("failed after retries: %w", err)
	}
	return nil
}

func setStatusIterations(appClient clientset.Interface, cd *appv1beta1.Canary, val int) error {
	firstTry := true
	name, ns := cd.GetName(), cd.GetNamespace()
	err := retry.RetryOnConflict(retry.DefaultBackoff, func() (err error) {
		if !firstTry {
			cd, err = appClient.AppV1beta1().Canaries(ns).Get(name, metav1.GetOptions{})
			if err != nil {
				return fmt.Errorf("canary %s.%s get query failed: %w", name, ns, err)
			}
		}

		cdCopy := cd.DeepCopy()
		cdCopy.Status.Iterations = val
		cdCopy.Status.LastTransitionTime = metav1.Now()

		if err = updateStatusWithUpgrade(appClient, cdCopy); err != nil {
			return fmt.Errorf("updateStatusWithUpgrade failed: %w", err)
		}
		firstTry = false
		return
	})

	if err != nil {
		return fmt.Errorf("failed after retries: %w", err)
	}
	return nil
}

func setStatusPhase(appClient clientset.Interface, cd *appv1beta1.Canary, phase appv1beta1.CanaryPhase) error {
	firstTry := true
	name, ns := cd.GetName(), cd.GetNamespace()
	err := retry.RetryOnConflict(retry.DefaultBackoff, func() (err error) {
		if !firstTry {
			cd, err = appClient.AppV1beta1().Canaries(ns).Get(name, metav1.GetOptions{})
			if err != nil {
				return fmt.Errorf("canary %s.%s get query failed: %w", name, ns, err)
			}
		}

		cdCopy := cd.DeepCopy()
		cdCopy.Status.Phase = phase
		cdCopy.Status.LastTransitionTime = metav1.Now()

		if phase != appv1beta1.CanaryPhaseProgressing && phase != appv1beta1.CanaryPhaseWaiting {
			cdCopy.Status.CanaryWeight = 0
			cdCopy.Status.Iterations = 0
		}

		// on promotion set primary spec hash
		if phase == appv1beta1.CanaryPhaseInitialized || phase == appv1beta1.CanaryPhaseSucceeded {
			cdCopy.Status.LastPromotedSpec = cd.Status.LastAppliedSpec
		}

		if ok, conditions := MakeStatusConditions(cdCopy, phase); ok {
			cdCopy.Status.Conditions = conditions
		}

		if err = updateStatusWithUpgrade(appClient, cdCopy); err != nil {
			return fmt.Errorf("updateStatusWithUpgrade failed: %w", err)
		}
		firstTry = false
		return
	})
	if err != nil {
		return fmt.Errorf("failed after retries: %w", err)
	}
	return nil
}

// getStatusCondition returns a condition based on type
func getStatusCondition(status appv1beta1.CanaryStatus, conditionType appv1beta1.CanaryConditionType) *appv1beta1.CanaryCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == conditionType {
			return &c
		}
	}
	return nil
}

// MakeStatusCondition updates the canary status conditions based on canary phase
func MakeStatusConditions(cd *appv1beta1.Canary,
	phase appv1beta1.CanaryPhase) (bool, []appv1beta1.CanaryCondition) {
	currentCondition := getStatusCondition(cd.Status, appv1beta1.PromotedType)

	message := fmt.Sprintf("New %s detected, starting initialization.", cd.Spec.TargetRef.Kind)
	status := corev1.ConditionUnknown
	switch phase {
	case appv1beta1.CanaryPhaseInitializing:
		status = corev1.ConditionUnknown
		message = fmt.Sprintf("New %s detected, starting initialization.", cd.Spec.TargetRef.Kind)
	case appv1beta1.CanaryPhaseInitialized:
		status = corev1.ConditionTrue
		message = fmt.Sprintf("%s initialization completed.", cd.Spec.TargetRef.Kind)
	case appv1beta1.CanaryPhaseWaiting:
		status = corev1.ConditionUnknown
		message = "Waiting for approval."
	case appv1beta1.CanaryPhaseProgressing:
		status = corev1.ConditionUnknown
		message = "New revision detected, starting canary analysis."
	case appv1beta1.CanaryPhasePromoting:
		status = corev1.ConditionUnknown
		message = "Canary analysis completed, starting primary rolling update."
	case appv1beta1.CanaryPhaseFinalising:
		status = corev1.ConditionUnknown
		message = "Canary analysis completed, routing all traffic to primary."
	case appv1beta1.CanaryPhaseSucceeded:
		status = corev1.ConditionTrue
		message = "Canary analysis completed successfully, promotion finished."
	case appv1beta1.CanaryPhaseFailed:
		status = corev1.ConditionFalse
		message = fmt.Sprintf("Canary analysis failed, %s scaled to zero.", cd.Spec.TargetRef.Kind)
	case appv1beta1.CanaryPhaseBegin:
		status = corev1.ConditionTrue
		message = fmt.Sprintf("Canary weight changed to %d .", cd.Status.CanaryWeight)
	}

	newCondition := &appv1beta1.CanaryCondition{
		Type:               appv1beta1.PromotedType,
		Status:             status,
		LastUpdateTime:     metav1.Now(),
		LastTransitionTime: metav1.Now(),
		Message:            message,
		Reason:             string(phase),
	}

	if currentCondition != nil &&
		currentCondition.Status == newCondition.Status &&
		currentCondition.Reason == newCondition.Reason {
		return false, nil
	}

	if currentCondition != nil && currentCondition.Status == newCondition.Status {
		newCondition.LastTransitionTime = currentCondition.LastTransitionTime
	}

	return true, []appv1beta1.CanaryCondition{*newCondition}
}

// updateStatusWithUpgrade tries to update the status sub-resource
// if the status update fails with:
// Canary.flagger.app is invalid: apiVersion: Invalid value: flagger.app/v1alpha3: must be flagger.app/v1beta1
// then the canary object will be updated to the latest API version
func updateStatusWithUpgrade(appClient clientset.Interface, cd *appv1beta1.Canary) error {
	_, err := appClient.AppV1beta1().Canaries(cd.Namespace).UpdateStatus(cd)
	if err != nil {
		// upgrade alpha resource
		if _, updateErr := appClient.AppV1beta1().Canaries(cd.Namespace).Update(cd); updateErr != nil {
			return fmt.Errorf("updating canary %s.%s from v1alpha to v1beta failed: %w", cd.Name, cd.Namespace, updateErr)
		}
		// retry status update
		_, err = appClient.AppV1beta1().Canaries(cd.Namespace).UpdateStatus(cd)
	}

	if err != nil {
		return fmt.Errorf("updating canary %s.%s status failed: %w", cd.Name, cd.Namespace, err)
	}
	return err
}
