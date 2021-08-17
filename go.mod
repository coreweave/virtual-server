module github.com/coreweave/virtual-server

go 1.15

require (
	github.com/kr/text v0.2.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9 // indirect
	golang.org/x/net v0.0.0-20201209123823-ac852fbbde11 // indirect
	golang.org/x/tools v0.0.0-20200725200936-102e7d357031 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.20.1
	k8s.io/apiextensions-apiserver v0.19.6 // indirect
	k8s.io/apimachinery v0.20.1
	kubevirt.io/client-go v0.36.0
	kubevirt.io/containerized-data-importer v1.26.1
	sigs.k8s.io/controller-runtime v0.7.1
)

replace (
	github.com/altoros/gosigma => github.com/juju/gosigma v0.0.0-20200420012028-063911838a9e
	github.com/hashicorp/raft => github.com/juju/raft v2.0.0-20200420012049-88ad3b3f0a54+incompatible
	github.com/openshift/api => github.com/openshift/api v0.0.0-20191219222812-2987a591a72c
	github.com/openshift/client-go => github.com/openshift/client-go v0.0.0-20191125132246-f6563a70e19a
	github.com/operator-framework/operator-lifecycle-manager => github.com/operator-framework/operator-lifecycle-manager v0.0.0-20190128024246-5eb7ae5bdb7a
	k8s.io/client-go => k8s.io/client-go v0.20.0
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.16.4
	kubevirt.io/containerized-data-importer => kubevirt.io/containerized-data-importer v1.26.1
)
