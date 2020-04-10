package controller

import (
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appv1beta1 "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/apis/alibabacloud/v1beta1"
	"gitlab.alibaba-inc.com/cos/application-center-manager/pkg/canary"
	"gitlab.alibaba-inc.com/cos/application-center-manager/router"
)

const (
	MetricsProviderServiceSuffix = ":service"
)

// scheduleCanaries synchronises the canary map with the jobs map,
// for new canaries new jobs are created and started
// for the removed canaries the jobs are stopped and deleted
func (c *Controller) scheduleCanaries() {
	current := make(map[string]string)
	stats := make(map[string]int)

	c.canaries.Range(func(key interface{}, value interface{}) bool {
		cn := value.(*appv1beta1.Canary)

		// format: <name>.<namespace>
		name := key.(string)
		current[name] = fmt.Sprintf("%s.%s", cn.Spec.TargetRef.Name, cn.Namespace)

		job, exists := c.jobs[name]
		// schedule new job for existing job with different analysis interval or non-existing job
		if (exists && job.GetCanaryAnalysisInterval() != cn.GetAnalysisInterval()) || !exists {
			if exists {
				job.Stop()
			}

			newJob := CanaryJob{
				Name:             cn.Name,
				Namespace:        cn.Namespace,
				function:         c.advanceCanary,
				done:             make(chan bool),
				ticker:           time.NewTicker(cn.GetAnalysisInterval()),
				analysisInterval: cn.GetAnalysisInterval(),
			}

			c.jobs[name] = newJob
			newJob.Start()
		}

		// compute canaries per namespace total
		t, ok := stats[cn.Namespace]
		if !ok {
			stats[cn.Namespace] = 1
		} else {
			stats[cn.Namespace] = t + 1
		}
		return true
	})

	// cleanup deleted jobs
	for job := range c.jobs {
		if _, exists := current[job]; !exists {
			c.jobs[job].Stop()
			delete(c.jobs, job)
		}
	}

	// check if multiple canaries have the same target
	for canaryName, targetName := range current {
		for name, target := range current {
			if name != canaryName && target == targetName {
				c.logger.With("canary", canaryName).
					Errorf("Bad things will happen! Found more than one canary with the same target %s", targetName)
			}
		}
	}

}

func (c *Controller) advanceCanary(name string, namespace string) {
	c.logger.Info("advanceCanary get in")
	// check if the canary exists
	cd, err := c.appClient.AppV1beta1().Canaries(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		c.logger.With("canary", fmt.Sprintf("%s.%s", name, namespace)).
			Errorf("Canary %s.%s not found", name, namespace)
		return
	}

	// override the global provider if one is specified in the canary spec
	provider := c.meshProvider
	if cd.Spec.Provider != "" {
		provider = cd.Spec.Provider
	}
	c.logger.Infof("Current provider is %s", provider)

	// init controller based on target kind
	c.logger.Infof("Init %s kind controller", cd.Spec.TargetRef.Kind)
	canaryController := c.canaryFactory.Controller(cd.Spec.TargetRef.Kind)
	labelSelector, ports, err := canaryController.GetMetadata(cd)
	if err != nil {
		c.recordEventWarningf(cd, "%v", err)
		return
	}
	c.logger.Infof("labelSelector %++v ports %++v", labelSelector, ports)

	// init Kubernetes router
	kubeRouter := c.routerFactory.KubernetesRouter(cd.Spec.TargetRef.Kind, labelSelector, map[string]string{}, ports)
	if err := kubeRouter.Initialize(cd); err != nil {
		c.recordEventWarningf(cd, "%v", err)
		return
	}

	// create primary
	err = canaryController.Initialize(cd)
	if err != nil {
		c.recordEventWarningf(cd, "%v", err)
		return
	}

	// init mesh router
	meshRouter := c.routerFactory.MeshRouter(provider)

	// create or update svc
	if err := kubeRouter.Reconcile(cd); err != nil {
		c.recordEventWarningf(cd, "%v", err)
		return
	}

	// create or update mesh routes
	if err := meshRouter.Reconcile(cd); err != nil {
		c.recordEventWarningf(cd, "%v", err)
		return
	}

	// check for changes
	shouldAdvance, err := c.shouldAdvance(cd, canaryController)
	if err != nil {
		c.recordEventWarningf(cd, "%v", err)
		return
	}
	c.logger.Infof("shouldAdvance %v", shouldAdvance)

	// check primary status
	if err := canaryController.IsPrimaryReady(cd); err != nil {
		c.recordEventWarningf(cd, "%v", err)
		return
	}

	c.logger.Info("1")
	// get the routing settings
	primaryWeight, canaryWeight, _, err := meshRouter.GetRoutes(cd)
	if err != nil {
		c.recordEventWarningf(cd, "%v", err)
		return
	}

	c.logger.Info("2")
	// check if canary analysis should start (canary revision has changes) or continue
	if ok := c.checkCanaryStatus(cd, canaryController, shouldAdvance); !ok {
		return
	}

	c.logger.Info("2.1")
	// check if canary revision changed during analysis
	if restart := c.hasCanaryRevisionChanged(cd, canaryController); restart {
		c.recordEventInfof(cd, "New revision detected! Restarting analysis for %s.%s",
			cd.Spec.TargetRef.Name, cd.Namespace)

		// route all traffic back to primary
		primaryWeight = 100
		canaryWeight = 0
		if err := meshRouter.SetRoutes(cd, primaryWeight, canaryWeight, false); err != nil {
			c.recordEventWarningf(cd, "%v", err)
			return
		}

		// reset status
		status := appv1beta1.CanaryStatus{
			Phase:        appv1beta1.CanaryPhaseProgressing,
			CanaryWeight: 0,
			FailedChecks: 0,
			Iterations:   0,
		}
		if err := canaryController.SyncStatus(cd, status); err != nil {
			c.recordEventWarningf(cd, "%v", err)
		}
		return
	} else {
		c.logger.Infof("canary changed %v", restart)
	}

	c.logger.Info("3")
	// check canary status
	var retriable = true
	retriable, err = canaryController.IsCanaryReady(cd)
	if err != nil && retriable {
		c.recordEventWarningf(cd, "%v", err)
		return
	}

	c.logger.Infof("cd status %++v", cd.Status)
	if cd.Status.Phase == appv1beta1.CanaryPhaseBegin {
		c.runCanary(cd, canaryController, meshRouter, false,
			cd.Status.CanaryWeight, 100-cd.Status.CanaryWeight, 100)
	}

	if cd.Status.Phase == appv1beta1.CanaryPhaseRollback {
		c.rollback(cd, canaryController, meshRouter)
	}
}

func (c *Controller) runCanary(canary *appv1beta1.Canary, canaryController canary.Controller,
	meshRouter router.Interface, mirrored bool, canaryWeight int, primaryWeight int, maxWeight int) {
	primaryName := fmt.Sprintf("%s-primary", canary.Spec.TargetRef.Name)

	// increase traffic weight
	if canaryWeight < maxWeight {

		if primaryWeight < 0 {
			primaryWeight = 0
		}
		if canaryWeight > 100 {
			canaryWeight = 100
		}

		if err := meshRouter.SetRoutes(canary, primaryWeight, canaryWeight, mirrored); err != nil {
			c.recordEventWarningf(canary, "%v", err)
			return
		}

		if err := canaryController.SetStatusWeight(canary, canaryWeight); err != nil {
			c.recordEventWarningf(canary, "%v", err)
			return
		}

		c.recordEventInfof(canary, "Advance %s.%s canary weight %v", canary.Name, canary.Namespace, canaryWeight)
		// update status phase

		if err := canaryController.SyncStatus(canary, appv1beta1.CanaryStatus{Phase: appv1beta1.CanaryPhaseBegin}); err != nil {
			c.logger.With("canary", fmt.Sprintf("%s.%s", canary.Name, canary.Namespace)).Errorf("%v", err)
		}
		return
	}

	// promote canary - max weight reached
	if canaryWeight >= maxWeight {

		// update primary spec
		c.recordEventInfof(canary, "Copying %s.%s template spec to %s.%s",
			canary.Spec.TargetRef.Name, canary.Namespace, primaryName, canary.Namespace)
		if err := canaryController.Promote(canary); err != nil {
			c.recordEventWarningf(canary, "%v", err)
			return
		}

		// update status phase
		if err := canaryController.SetStatusPhase(canary, appv1beta1.CanaryPhasePromoting); err != nil {
			c.recordEventWarningf(canary, "%v", err)
			return
		}
	}
}

func (c *Controller) runAnalysis(canary *appv1beta1.Canary) bool {
	// run external checks
	for _, webhook := range canary.GetAnalysis().Webhooks {
		if webhook.Type == "" || webhook.Type == appv1beta1.RolloutHook {
			err := CallWebhook(canary.Name, canary.Namespace, appv1beta1.CanaryPhaseProgressing, webhook)
			if err != nil {
				c.recordEventWarningf(canary, "Halt %s.%s advancement external check %s failed %v",
					canary.Name, canary.Namespace, webhook.Name, err)
				return false
			}
		}
	}

	/*
		ok := c.runBuiltinMetricChecks(canary)
		if !ok {
			return ok
		}

		ok = c.runMetricChecks(canary)
		if !ok {
			return ok
		}
	*/
	return true
}

func (c *Controller) shouldSkipAnalysis(canary *appv1beta1.Canary, canaryController canary.Controller, meshRouter router.Interface) bool {
	if !canary.SkipAnalysis() {
		return false
	}

	// route all traffic to primary
	primaryWeight := 100
	canaryWeight := 0
	if err := meshRouter.SetRoutes(canary, primaryWeight, canaryWeight, false); err != nil {
		c.recordEventWarningf(canary, "%v", err)
		return false
	}

	// copy spec and configs from canary to primary
	c.recordEventInfof(canary, "Copying %s.%s template spec to %s-primary.%s",
		canary.Spec.TargetRef.Name, canary.Namespace, canary.Spec.TargetRef.Name, canary.Namespace)
	if err := canaryController.Promote(canary); err != nil {
		c.recordEventWarningf(canary, "%v", err)
		return false
	}

	// shutdown canary
	if err := canaryController.ScaleToZero(canary); err != nil {
		c.recordEventWarningf(canary, "%v", err)
		return false
	}

	// update status phase
	if err := canaryController.SetStatusPhase(canary, appv1beta1.CanaryPhaseSucceeded); err != nil {
		c.recordEventWarningf(canary, "%v", err)
		return false
	}

	// notify
	c.recordEventInfof(canary, "Promotion completed! Canary analysis was skipped for %s.%s",
		canary.Spec.TargetRef.Name, canary.Namespace)

	return true
}

func (c *Controller) shouldAdvance(canary *appv1beta1.Canary, canaryController canary.Controller) (bool, error) {
	if canary.Status.LastAppliedSpec == "" ||
		canary.Status.Phase == appv1beta1.CanaryPhaseInitializing ||
		canary.Status.Phase == appv1beta1.CanaryPhaseProgressing ||
		canary.Status.Phase == appv1beta1.CanaryPhaseWaiting ||
		canary.Status.Phase == appv1beta1.CanaryPhasePromoting ||
		canary.Status.Phase == appv1beta1.CanaryPhaseFinalising {
		return true, nil
	}

	newTarget, err := canaryController.HasTargetChanged(canary)
	if err != nil {
		return false, err
	}
	if newTarget {
		return newTarget, nil
	}

	newCfg, err := canaryController.HaveDependenciesChanged(canary)
	if err != nil {
		return false, err
	}

	return newCfg, nil

}

func (c *Controller) checkCanaryStatus(canary *appv1beta1.Canary, canaryController canary.Controller, shouldAdvance bool) bool {
	if canary.Status.Phase == appv1beta1.CanaryPhaseProgressing ||
		canary.Status.Phase == appv1beta1.CanaryPhasePromoting ||
		canary.Status.Phase == appv1beta1.CanaryPhaseFinalising ||
		canary.Status.Phase == appv1beta1.CanaryPhaseBegin {
		return true
	}

	if canary.Status.Phase == "" || canary.Status.Phase == appv1beta1.CanaryPhaseInitializing {
		if err := canaryController.SyncStatus(canary, appv1beta1.CanaryStatus{Phase: appv1beta1.CanaryPhaseInitialized}); err != nil {
			c.logger.With("canary", fmt.Sprintf("%s.%s", canary.Name, canary.Namespace)).Errorf("%v", err)
			return false
		}
		c.recordEventInfof(canary, "Initialization done! %s.%s", canary.Name, canary.Namespace)
		return false
	}

	if shouldAdvance {
		canaryPhaseProgressing := canary.DeepCopy()
		canaryPhaseProgressing.Status.Phase = appv1beta1.CanaryPhaseProgressing
		c.recordEventInfof(canaryPhaseProgressing, "New revision detected! Scaling up %s.%s", canaryPhaseProgressing.Spec.TargetRef.Name, canaryPhaseProgressing.Namespace)

		if err := canaryController.ScaleFromZero(canary); err != nil {
			c.recordEventErrorf(canary, "%v", err)
			return false
		}
		if err := canaryController.SyncStatus(canary, appv1beta1.CanaryStatus{Phase: appv1beta1.CanaryPhaseProgressing}); err != nil {
			c.logger.With("canary", fmt.Sprintf("%s.%s", canary.Name, canary.Namespace)).Errorf("%v", err)
			return false
		}
		return false
	}
	return false
}

func (c *Controller) hasCanaryRevisionChanged(canary *appv1beta1.Canary, canaryController canary.Controller) bool {
	if canary.Status.Phase == appv1beta1.CanaryPhaseProgressing {
		if diff, _ := canaryController.HasTargetChanged(canary); diff {
			return true
		}
		if diff, _ := canaryController.HaveDependenciesChanged(canary); diff {
			return true
		}
	}
	return false
}

func (c *Controller) rollback(canary *appv1beta1.Canary, canaryController canary.Controller, meshRouter router.Interface) {
	if canary.Status.FailedChecks >= canary.GetAnalysisThreshold() {
		c.recordEventWarningf(canary, "Rolling back %s.%s failed checks threshold reached %v",
			canary.Name, canary.Namespace, canary.Status.FailedChecks)
	}

	// route all traffic back to primary
	primaryWeight := 100
	canaryWeight := 0
	if err := meshRouter.SetRoutes(canary, primaryWeight, canaryWeight, false); err != nil {
		c.recordEventWarningf(canary, "%v", err)
		return
	}

	canaryPhaseFailed := canary.DeepCopy()
	canaryPhaseFailed.Status.Phase = appv1beta1.CanaryPhaseFailed
	c.recordEventWarningf(canaryPhaseFailed, "Canary failed! Scaling down %s.%s",
		canaryPhaseFailed.Name, canaryPhaseFailed.Namespace)

	// shutdown canary
	if err := canaryController.ScaleToZero(canary); err != nil {
		c.recordEventWarningf(canary, "%v", err)
		return
	}

	// mark canary as failed
	if err := canaryController.SyncStatus(canary, appv1beta1.CanaryStatus{Phase: appv1beta1.CanaryPhaseFailed, CanaryWeight: 0}); err != nil {
		c.logger.With("canary", fmt.Sprintf("%s.%s", canary.Name, canary.Namespace)).Errorf("%v", err)
		return
	}

}
