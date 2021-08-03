module github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/image-vuln-adapter

go 1.16

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/google/go-containerregistry v0.6.0
	github.com/google/uuid v1.2.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	k8s.io/api v0.21.3
	k8s.io/apiextensions-apiserver v0.21.2
	k8s.io/apimachinery v0.21.3
	k8s.io/cli-runtime v0.21.3
	k8s.io/client-go v0.21.3
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20210527160623-6fdb442a123b
	sigs.k8s.io/controller-runtime v0.9.2
)
