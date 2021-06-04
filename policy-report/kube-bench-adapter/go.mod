module github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/kube-bench-adapter

go 1.16

require (
	github.com/aquasecurity/kube-bench v0.6.0
	github.com/pelletier/go-toml v1.8.1 // indirect
	k8s.io/api v0.20.5
	k8s.io/apimachinery v0.20.5
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/code-generator v0.20.1
	k8s.io/klog v0.3.3
	sigs.k8s.io/controller-runtime v0.8.3
)

replace k8s.io/client-go => k8s.io/client-go v0.20.5

replace k8s.io/api => k8s.io/api v0.20.5
