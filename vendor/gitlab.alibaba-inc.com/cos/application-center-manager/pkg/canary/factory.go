package canary

import (
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"

	clientset "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/client/clientset/versioned"
)

type Factory struct {
	kubeClient    kubernetes.Interface
	appClient     clientset.Interface
	logger        *zap.SugaredLogger
	configTracker Tracker
	labels        []string
}

func NewFactory(kubeClient kubernetes.Interface,
	appClient clientset.Interface,
	configTracker Tracker,
	labels []string,
	logger *zap.SugaredLogger) *Factory {
	return &Factory{
		kubeClient:    kubeClient,
		appClient:     appClient,
		logger:        logger,
		configTracker: configTracker,
		labels:        labels,
	}
}

func (factory *Factory) Controller(kind string) Controller {
	deploymentCtrl := &DeploymentController{
		logger:        factory.logger,
		kubeClient:    factory.kubeClient,
		appClient:     factory.appClient,
		labels:        factory.labels,
		configTracker: factory.configTracker,
	}
	serviceCtrl := &ServiceController{
		logger:     factory.logger,
		kubeClient: factory.kubeClient,
		appClient:  factory.appClient,
	}

	switch kind {
	case "Deployment":
		return deploymentCtrl
	case "Service":
		return serviceCtrl
	default:
		return deploymentCtrl
	}
}
