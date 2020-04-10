package main

import (
	"flag"
	"fmt"
	"gitlab.alibaba-inc.com/cos/application-center-manager/pkg/canary"
	clientset "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/client/clientset/versioned"
	informers "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/client/informers/externalversions"
	"gitlab.alibaba-inc.com/cos/application-center-manager/pkg/controller"
	"gitlab.alibaba-inc.com/cos/application-center-manager/pkg/logger"
	"gitlab.alibaba-inc.com/cos/application-center-manager/router"
	"gitlab.alibaba-inc.com/cos/application-center-manager/server"
	"gitlab.alibaba-inc.com/cos/application-center-manager/signals"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"strings"
	"time"
)

var (
	masterURL                string
	kubeConfig               string
	logLevel                 string
	zapEncoding              string
	selectorLabels           string
	port                     string
	ingressAnnotationsPrefix string
	enableConfigTracking     bool
	controlLoopInterval      time.Duration
	meshProvider             string
	eventWebhook             string
	threadiness              int
)

func init() {
	flag.StringVar(&zapEncoding, "zap-encoding", "json", "Zap logger encoding.")
	flag.StringVar(&port, "port", "8080", "Port to listen on.")
	flag.StringVar(&logLevel, "log-level", "debug", "Log level can be: debug, info, warning, error.")
	flag.StringVar(&kubeConfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&selectorLabels, "selector-labels", "app,name,app.kubernetes.io/name", "List of pod labels that Flagger uses to create pod selectors.")
	flag.StringVar(&ingressAnnotationsPrefix, "ingress-annotations-prefix", "nginx.ingress.kubernetes.io", "Annotations prefix for ingresses.")
	flag.BoolVar(&enableConfigTracking, "enable-config-tracking", true, "Enable secrets and configmaps tracking.")
	flag.DurationVar(&controlLoopInterval, "control-loop-interval", 10*time.Second, "Kubernetes API sync interval.")
	flag.StringVar(&meshProvider, "mesh-provider", "istio", "Service mesh provider, can be istio, linkerd, appmesh, supergloo, nginx or smi.")
	flag.StringVar(&eventWebhook, "event-webhook", "", "Webhook for publishing flagger events")
	flag.IntVar(&threadiness, "threadiness", 2, "Worker concurrency.")
}

func main() {
	flag.Parse()

	fmt.Printf("Manager Version %++v", GetVersion())
	logger, err := logger.NewLoggerWithEncoding(logLevel, zapEncoding)
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}
	defer logger.Sync()

	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeConfig)
	if err != nil {
		logger.Fatalf("Error building kubeconfig: %v", err)
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		logger.Fatalf("Error building kubernetes clientset: %v", err)
	}

	appClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		logger.Fatalf("Error building appctl clientset: %s", err.Error())
	}

	verifyCRDs(appClient, logger)
	infos := startInformers(appClient, logger, stopCh)

	labels := strings.Split(selectorLabels, ",")
	if len(labels) < 1 {
		logger.Fatalf("At least one selector label is required")
	}

	// start HTTP server
	go server.ListenAndServe(port, 3*time.Second, logger, stopCh)

	routerFactory := router.NewFactory(cfg, kubeClient, appClient, ingressAnnotationsPrefix, logger)

	var configTracker canary.Tracker
	if enableConfigTracking {
		configTracker = &canary.ConfigTracker{
			Logger:     logger,
			KubeClient: kubeClient,
			AppClient:  appClient,
		}
	} else {
		configTracker = &canary.NopTracker{}
	}

	canaryFactory := canary.NewFactory(kubeClient, appClient, configTracker, labels, logger)

	c := controller.NewController(
		kubeClient,
		appClient,
		infos,
		controlLoopInterval,
		logger,
		canaryFactory,
		routerFactory,
		meshProvider,
		fromEnv("EVENT_WEBHOOK_URL", eventWebhook),
	)

	if err := c.Run(threadiness, stopCh); err != nil {
		logger.Fatalf("Error running controller: %v", err)
	}
}

func verifyCRDs(appClient clientset.Interface, logger *zap.SugaredLogger) {
	_, err := appClient.AppV1beta1().Canaries("").List(metav1.ListOptions{Limit: 1})
	if err != nil {
		logger.Fatalf("Canary CRD is not registered %v", err)
	}
}

func startInformers(appClient clientset.Interface, logger *zap.SugaredLogger, stopCh <-chan struct{}) controller.Informers {
	appInformerFactory := informers.NewSharedInformerFactoryWithOptions(appClient, time.Second*30, informers.WithNamespace(""))

	logger.Info("Waiting for canary informer cache to sync")
	canaryInformer := appInformerFactory.App().V1beta1().Canaries()
	go canaryInformer.Informer().Run(stopCh)
	if ok := cache.WaitForNamedCacheSync("application-center-manager", stopCh, canaryInformer.Informer().HasSynced); !ok {
		logger.Fatalf("failed to wait for cache to sync")
	}

	return controller.Informers{
		CanaryInformer: canaryInformer,
	}
}

func fromEnv(envVar string, defaultVal string) string {
	if v := os.Getenv(envVar); v != "" {
		return v
	}
	return defaultVal
}
