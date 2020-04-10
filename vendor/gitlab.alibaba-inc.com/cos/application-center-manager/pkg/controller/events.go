package controller

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"

	appv1beta1 "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/apis/alibabacloud/v1beta1"
)

func (c *Controller) recordEventInfof(r *appv1beta1.Canary, template string, args ...interface{}) {
	c.logger.With("canary", fmt.Sprintf("%s.%s", r.Name, r.Namespace)).Infof(template, args...)
	c.eventRecorder.Event(r, corev1.EventTypeNormal, "Synced", fmt.Sprintf(template, args...))
	c.sendEventToWebhook(r, corev1.EventTypeNormal, template, args)
}

func (c *Controller) recordEventErrorf(r *appv1beta1.Canary, template string, args ...interface{}) {
	c.logger.With("canary", fmt.Sprintf("%s.%s", r.Name, r.Namespace)).Errorf(template, args...)
	c.eventRecorder.Event(r, corev1.EventTypeWarning, "Synced", fmt.Sprintf(template, args...))
	c.sendEventToWebhook(r, corev1.EventTypeWarning, template, args)
}

func (c *Controller) recordEventWarningf(r *appv1beta1.Canary, template string, args ...interface{}) {
	c.logger.With("canary", fmt.Sprintf("%s.%s", r.Name, r.Namespace)).Infof(template, args...)
	c.eventRecorder.Event(r, corev1.EventTypeWarning, "Synced", fmt.Sprintf(template, args...))
	c.sendEventToWebhook(r, corev1.EventTypeWarning, template, args)
}

func (c *Controller) sendEventToWebhook(r *appv1beta1.Canary, eventType, template string, args []interface{}) {
	webhookOverride := false
	if r.GetAnalysis() == nil {
		return
	}
	for _, canaryWebhook := range r.GetAnalysis().Webhooks {
		if canaryWebhook.Type == appv1beta1.EventHook {
			webhookOverride = true
			err := CallEventWebhook(r, canaryWebhook.URL, fmt.Sprintf(template, args...), eventType)
			if err != nil {
				c.logger.With("canary", fmt.Sprintf("%s.%s", r.Name, r.Namespace)).Errorf("error sending event to webhook: %s", err)
			}
		}
	}

	if c.eventWebhook != "" && !webhookOverride {
		err := CallEventWebhook(r, c.eventWebhook, fmt.Sprintf(template, args...), eventType)
		if err != nil {
			c.logger.With("canary", fmt.Sprintf("%s.%s", r.Name, r.Namespace)).Errorf("error sending event to webhook: %s", err)
		}
	}
}
