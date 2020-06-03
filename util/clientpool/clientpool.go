package clientpool

import (
	"fmt"
	appclientset "github.com/argoproj/argo-cd/pkg/client/clientset/versioned"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"time"
)

type ClientPool struct {
	cache         *cache.Cache
	kubeClientset kubernetes.Interface
	appClientset  appclientset.Interface
}

var pool *ClientPool

func InitClientPool(kubeClientset kubernetes.Interface, appClientset appclientset.Interface) {
	log.Infof("init client pool")
	pool = &ClientPool{
		cache:         cache.New(5*time.Minute, 10*time.Minute),
		kubeClientset: kubeClientset,
		appClientset:  appClientset,
	}
}

func GetPool() *ClientPool {
	log.Infof("get client pool: %v", pool)
	return pool
}

func (p ClientPool) GetAppClientset(ctx context.Context) appclientset.Interface {
	impersonate := getImpersonateFromContext(ctx)
	if impersonate == "" {
		log.Infof("impersonate identity not provided, using default client")
		return p.appClientset
	}

	cachedClient, found := p.cache.Get(getAppClientCacheKey(impersonate))
	if found {
		log.Infof("appclient cache hit for %s", impersonate)
		return cachedClient.(appclientset.Interface)
	}

	restConfig, err := buildRestConfig(impersonate)
	if err != nil {
		log.Warnf("failed to build client config for %s, using default client", impersonate)
		return p.appClientset
	}
	clientset, err := appclientset.NewForConfig(restConfig)
	if err != nil {
		log.Warnf("failed to build clientset for %s, using default client", impersonate)
		return p.appClientset
	}

	p.cache.Set(getAppClientCacheKey(impersonate), clientset, cache.DefaultExpiration)
	log.Infof("create new appclient for %s", impersonate)
	return clientset
}

func (p ClientPool) GetKubeClientset(ctx context.Context) kubernetes.Interface {
	impersonate := getImpersonateFromContext(ctx)
	if impersonate == "" {
		log.Infof("impersonate identity not provided, using default client")
		return p.kubeClientset
	}

	cachedClient, found := p.cache.Get(getKubeClientCacheKey(impersonate))
	if found {
		log.Infof("appclient cache hit for %s", impersonate)
		return cachedClient.(kubernetes.Interface)
	}

	restConfig, err := buildRestConfig(impersonate)
	if err != nil {
		log.Warnf("failed to build client config for %s, using default client", impersonate)
		return p.kubeClientset
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		log.Warnf("failed to build clientset for %s, using default client", impersonate)
		return p.kubeClientset
	}

	p.cache.Set(getKubeClientCacheKey(impersonate), clientset, cache.DefaultExpiration)
	log.Infof("create new appclient for %s", impersonate)
	return clientset
}

func getImpersonateFromContext(ctx context.Context) string {
	userName := ""
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		val := md.Get("impersonate-user")
		if len(val) > 0 {
			userName = val[0]
		}

		log.Infof("Impersonate: %s", userName)
	}

	return userName
}

func buildRestConfig(impersonate string) (*rest.Config, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	overrides := clientcmd.ConfigOverrides{
		AuthInfo: clientcmdapi.AuthInfo{
			Impersonate: impersonate,
		},
	}
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &overrides)
	return config.ClientConfig()
}

func getAppClientCacheKey(impersonate string) string {
	return fmt.Sprintf("%s-appclient", impersonate)
}

func getKubeClientCacheKey(impersonate string) string {
	return fmt.Sprintf("%s-kubeclient", impersonate)
}
