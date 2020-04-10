module gitlab.alibaba-inc.com/cos/application-center-manager

go 1.14

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/google/go-cmp v0.4.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.5.1
	github.com/stretchr/testify v1.5.1 // indirect
	go.uber.org/zap v1.14.1
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2 // indirect
	k8s.io/api v0.17.4
	k8s.io/apimachinery v0.17.4
	k8s.io/client-go v0.17.4
	k8s.io/utils v0.0.0-20191114184206-e782cd3c129f
	sigs.k8s.io/controller-tools v0.2.1 // indirect
)

replace k8s.io/klog => github.com/stefanprodan/klog v0.0.0-20190418165334-9cbb78b20423
