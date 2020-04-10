package router

import (
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"

	clientset "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/client/clientset/versioned"
)

type Factory struct {
	kubeConfig               *restclient.Config
	kubeClient               kubernetes.Interface
	appClient                clientset.Interface
	ingressAnnotationsPrefix string
	logger                   *zap.SugaredLogger
}

func NewFactory(kubeConfig *restclient.Config, kubeClient kubernetes.Interface,
	appClient clientset.Interface,
	ingressAnnotationsPrefix string,
	logger *zap.SugaredLogger) *Factory {
	return &Factory{
		kubeConfig:               kubeConfig,
		kubeClient:               kubeClient,
		appClient:                appClient,
		ingressAnnotationsPrefix: ingressAnnotationsPrefix,
		logger:                   logger,
	}
}

// KubernetesRouter returns a KubernetesRouter interface implementation
func (factory *Factory) KubernetesRouter(kind string, labelSelector string, annotations map[string]string, ports map[string]int32) KubernetesRouter {
	switch kind {
	case "Service":
		return &KubernetesNoopRouter{}
	default: // Daemonset or Deployment
		return &KubernetesDefaultRouter{
			logger:        factory.logger,
			appClient:     factory.appClient,
			kubeClient:    factory.kubeClient,
			labelSelector: labelSelector,
			annotations:   annotations,
			ports:         ports,
		}
	}
}

// MeshRouter returns a service mesh router
func (factory *Factory) MeshRouter(provider string) Interface {
	switch {
	case provider == "none":
		return &NopRouter{}
	case provider == "kubernetes":
		return &NopRouter{}
	case provider == "nginx":
		return &IngressRouter{
			logger:            factory.logger,
			kubeClient:        factory.kubeClient,
			annotationsPrefix: factory.ingressAnnotationsPrefix,
		}
	default:
		return &IngressRouter{
			logger:            factory.logger,
			kubeClient:        factory.kubeClient,
			annotationsPrefix: factory.ingressAnnotationsPrefix,
		}
	}
}
