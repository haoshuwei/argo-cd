package clientpool

import (
	"fmt"
	appclientset "github.com/argoproj/argo-cd/pkg/client/clientset/versioned"
	"github.com/jinzhu/copier"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

type ClientPool struct {
	cache         *cache.Cache
	loadingRules  *clientcmd.ClientConfigLoadingRules
	overrides     *clientcmd.ConfigOverrides
	kubeClientset kubernetes.Interface
	appClientset  appclientset.Interface
}

var pool *ClientPool

func InitClientPool(
	loadingRules *clientcmd.ClientConfigLoadingRules,
	overrides *clientcmd.ConfigOverrides,
	kubeClientset kubernetes.Interface,
	appClientset appclientset.Interface,
) {
	log.Infof("init client pool")
	pool = &ClientPool{
		cache:         cache.New(5*time.Minute, 10*time.Minute),
		loadingRules:  loadingRules,
		overrides:     overrides,
		kubeClientset: kubeClientset,
		appClientset:  appClientset,
	}
}

func GetPool() *ClientPool {
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

	restConfig, err := p.buildRestConfig(impersonate)
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

	restConfig, err := p.buildRestConfig(impersonate)
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

func (p ClientPool) buildRestConfig(impersonate string) (*rest.Config, error) {
	overrides := clientcmd.ConfigOverrides{}
	err := copier.Copy(&overrides, p.overrides)
	if err != nil {
		return nil, err
	}

	overrides.AuthInfo.Impersonate = impersonate
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(p.loadingRules, &overrides)
	return config.ClientConfig()
}

func getAppClientCacheKey(impersonate string) string {
	return fmt.Sprintf("%s-appclient", impersonate)
}

func getKubeClientCacheKey(impersonate string) string {
	return fmt.Sprintf("%s-kubeclient", impersonate)
}
