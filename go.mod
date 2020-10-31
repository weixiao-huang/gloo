module github.com/solo-io/gloo

go 1.14

require (
	contrib.go.opencensus.io/exporter/stackdriver v0.12.5 // indirect
	github.com/Masterminds/semver/v3 v3.0.3
	github.com/Netflix/go-expect v0.0.0-20180928190340-9d1f4485533b
	github.com/avast/retry-go v2.4.3+incompatible
	github.com/aws/aws-sdk-go v1.34.9
	github.com/cncf/udpa/go v0.0.0-20200629203442-efcf912fb354
	github.com/cratonica/2goarray v0.0.0-20190331194516-514510793eaa
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/elazarl/goproxy v0.0.0-20190711103511-473e67f1d7d2 // indirect
	github.com/elazarl/goproxy/ext v0.0.0-20190711103511-473e67f1d7d2 // indirect
	github.com/envoyproxy/go-control-plane v0.9.6-0.20200622162421-8d76a6f21925
	github.com/envoyproxy/protoc-gen-validate v0.4.0
	github.com/fgrosse/zaptest v1.1.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/go-openapi/loads v0.19.4
	github.com/go-openapi/spec v0.19.4
	github.com/go-openapi/swag v0.19.5
	github.com/go-swagger/go-swagger v0.21.0
	github.com/gogo/googleapis v1.3.1
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.2
	github.com/google/go-containerregistry v0.1.3 // indirect
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-github/v31 v31.0.0
	github.com/gorilla/mux v1.7.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4
	github.com/hashicorp/consul/api v1.3.0
	github.com/hashicorp/go-multierror v1.0.0
	github.com/hashicorp/go-uuid v1.0.2-0.20191001231223-f32f5fe8d6a8
	github.com/hashicorp/vault/api v1.0.5-0.20191108163347-bdd38fca2cff
	github.com/hinshun/vt10x v0.0.0-20180809195222-d55458df857c
	github.com/inconshreveable/go-update v0.0.0-20160112193335-8152e7eb6ccf
	github.com/jhump/protoreflect v1.5.0
	github.com/k0kubun/pp v2.3.0+incompatible
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/keybase/go-ps v0.0.0-20190827175125-91aafc93ba19
	github.com/mattbaird/jsonpatch v0.0.0-20200820163806-098863c1fc24 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/mitchellh/hashstructure v1.0.0
	github.com/mitchellh/reflectwalk v1.0.1
	github.com/olekukonko/tablewriter v0.0.4
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/opencontainers/go-digest v1.0.0-rc1
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/prometheus/client_golang v1.2.1
	github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4
	github.com/prometheus/prometheus v2.5.0+incompatible
	github.com/rotisserie/eris v0.4.0
	github.com/sergi/go-diff v1.0.0
	github.com/sirupsen/logrus v1.7.0 // indirect
	github.com/solo-io/envoy-operator v0.1.4
	github.com/solo-io/go-list-licenses v0.0.0-20191023220251-171e4740d00f
	github.com/solo-io/go-utils v0.16.6
	github.com/solo-io/protoc-gen-ext v0.0.9
	github.com/solo-io/reporting-client v0.1.2
	github.com/solo-io/skv2 v0.8.1
	github.com/solo-io/solo-apis v0.0.0-20200717214114-6a1daa5a5d05
	github.com/solo-io/solo-kit v0.13.13
	github.com/solo-io/wasm/tools/wasme/pkg v0.0.0-20200922223809-7fff932e9e9a
	github.com/spf13/afero v1.2.2
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
	go.opencensus.io v0.22.4
	go.uber.org/multierr v1.5.0
	go.uber.org/zap v1.15.0
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/mod v0.3.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	golang.org/x/sys v0.0.0-20200826173525-f9321e4c35a6 // indirect
	golang.org/x/tools v0.0.0-20200916195026-c9a70fc28ce3
	google.golang.org/genproto v0.0.0-20200626011028-ee7919e894b5
	google.golang.org/grpc v1.29.1
	gopkg.in/AlecAivazis/survey.v1 v1.8.7
	gopkg.in/src-d/go-billy.v4 v4.3.0 // indirect
	gopkg.in/yaml.v2 v2.3.0
	helm.sh/helm/v3 v3.1.2
	k8s.io/api v0.18.8
	k8s.io/apiextensions-apiserver v0.18.2
	k8s.io/apimachinery v0.18.8
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/code-generator v0.18.2
	k8s.io/helm v2.16.1+incompatible
	k8s.io/kube-openapi v0.0.0-20200121204235-bf4fb3bd569c // indirect
	k8s.io/kubectl v0.17.2
	k8s.io/kubernetes v1.17.1
	k8s.io/utils v0.17.1
	knative.dev/pkg v0.0.0-20191203174735-3444316bdeef
	knative.dev/serving v0.10.0
	sigs.k8s.io/controller-runtime v0.5.8
	sigs.k8s.io/yaml v1.2.0
)

replace (
	cloud.google.com/go => cloud.google.com/go v0.52.0 // cloud.google.com/go/iam breaks with v0.53.0+
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.0+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
	github.com/golang/mock => github.com/golang/mock v1.4.3
	github.com/golang/protobuf => github.com/golang/protobuf v1.4.1
	github.com/spf13/cobra => github.com/spf13/cobra v0.0.5 // override wasme requirement
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20200117163144-32f20d992d24 // override wasme requirement

	// kube 1.17
	k8s.io/api => k8s.io/api v0.17.1
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.1
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.1
	k8s.io/apiserver => k8s.io/apiserver v0.17.1
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.17.1
	k8s.io/client-go => k8s.io/client-go v0.17.1
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.17.1
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.17.1
	k8s.io/code-generator => k8s.io/code-generator v0.17.1
	k8s.io/component-base => k8s.io/component-base v0.17.1
	k8s.io/cri-api => k8s.io/cri-api v0.17.1
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.17.1
	k8s.io/gengo => k8s.io/gengo v0.0.0-20190822140433-26a664648505
	k8s.io/heapster => k8s.io/heapster v1.2.0-beta.1
	k8s.io/klog => github.com/stefanprodan/klog v0.0.0-20190418165334-9cbb78b20423
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.17.1
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.17.1
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190816220812-743ec37842bf
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.17.1
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.17.1
	k8s.io/kubectl => k8s.io/kubectl v0.17.1
	k8s.io/kubelet => k8s.io/kubelet v0.17.1
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.17.1
	k8s.io/metrics => k8s.io/metrics v0.17.1
	k8s.io/node-api => k8s.io/node-api v0.17.1
	k8s.io/repo-infra => k8s.io/repo-infra v0.0.0-20181204233714-00fe14e3d1a3
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.17.1
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.17.1
	k8s.io/sample-controller => k8s.io/sample-controller v0.17.1
	k8s.io/utils => k8s.io/utils v0.0.0-20190801114015-581e00157fb1
)
