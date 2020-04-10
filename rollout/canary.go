package rollout

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gitlab.alibaba-inc.com/cos/application-center-manager/pkg/apis/alibabacloud/v1beta1"
	clientset "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"os/user"
)

func GetAppClient(kubeConfig string) *clientset.Clientset {
	if kubeConfig == "~/.kube/config" {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		kubeConfig = fmt.Sprintf("%s/.kube/config", usr.HomeDir)
	}
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}
	appClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building appctl clientset: %s", err.Error())
	}

	return appClient
}

type Canary struct {
	appClient *clientset.Clientset
}

func NewCanaryController(kubeConfig string) *Canary {
	appClient := GetAppClient(kubeConfig)
	c := &Canary{}
	c.appClient = appClient
	return c
}

func (c *Canary) SetWeight(canaryWeight int, canaryName, namespace string) {
	canaryObj, err := c.appClient.AppV1beta1().Canaries(namespace).Get(canaryName, v1.GetOptions{})
	if err != nil {
		log.Fatalf("Get canary object %s:%s error %v", namespace, canaryName, err)
	}
	if canaryObj.Status.Phase != v1beta1.CanaryPhaseProgressing {
		log.Fatalf("Current canary %s status is %s, can not support set weight", canaryName, canaryObj.Status.Phase)
	}
	canaryObj.Status.CanaryWeight = canaryWeight
	canaryObj.Status.Phase = v1beta1.CanaryPhaseBegin
	_, err = c.appClient.AppV1beta1().Canaries(namespace).UpdateStatus(canaryObj)
	if err != nil {
		log.Fatalf("Update %s:%s status error %v", namespace, canaryName, err)
	}
}

func (c *Canary) ListCanary(namespace string) {
	canarys, err := c.appClient.AppV1beta1().Canaries(namespace).List(v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range canarys.Items {
		log.Infof("Name")
		log.Info()
		log.Info(c.Name)
	}
}
