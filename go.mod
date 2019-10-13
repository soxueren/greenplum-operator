module github.com/soxueren/greenplum-operator

require (
	github.com/Azure/go-autorest/autorest v0.9.2 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.7.0 // indirect
	github.com/DeanThompson/ginpprof v0.0.0-20190408063150-3be636683586
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc
	github.com/chilts/sid v0.0.0-20190607042430-660e94789ec9
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.25.4
	github.com/go-openapi/spec v0.19.0
	github.com/golang/groupcache v0.0.0-20190702054246-869f871628b6 // indirect
	github.com/googleapis/gnostic v0.3.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20190430165422-3e4dfb77656c // indirect
	github.com/gorilla/websocket v1.4.0
	github.com/hashicorp/consul/api v1.2.0 // indirect
	github.com/hashicorp/memberlist v0.1.5 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.11.1
	github.com/nats-io/nats-server/v2 v2.1.0 // indirect
	github.com/operator-framework/operator-sdk v0.0.0-20191012024916-f419ad3f3dc5
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190710185942-9d28bd7c0945 // indirect
	github.com/spf13/pflag v1.0.3
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.5.1
	go.uber.org/atomic v1.4.0 // indirect
	golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // indirect
	google.golang.org/grpc v1.23.0 // indirect
	k8s.io/api v0.0.0-20190918155943-95b840bb6a1f
	k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/kube-openapi v0.0.0-20190401085232-94e1e7b7574c
	k8s.io/utils v0.0.0-20190712204705-3dccf664f023 // indirect
	sigs.k8s.io/controller-runtime v0.2.0
)

// Pinned to kubernetes-1.14.1
replace (
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190409022649-727a075fdec8
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190409021438-1a26190bd76a
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190409023720-1bc0c81fa51d
)

replace (
	// Indirect operator-sdk dependencies use git.apache.org, which is frequently
	// down. The github mirror should be used instead.
	// Locking to a specific version (from 'go mod graph'):
	git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v0.0.0-20190906230412-69b4126ece6b
	github.com/coreos/prometheus-operator => github.com/coreos/prometheus-operator v0.31.1
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.6.1
	// Pinned to v2.10.0 (kubernetes-1.14.1) so https://proxy.golang.org can
	// resolve it correctly.
	github.com/prometheus/prometheus => github.com/prometheus/prometheus v0.0.0-20190525122359-d20e84d0fb64
)

go 1.13
