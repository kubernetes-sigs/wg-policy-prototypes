package params

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

var (
	Params TrivyArgs
)

func ParseArguments() {
	flag.StringVar(&Params.Name, "name", "trivy", "name of policy report")
	flag.StringVar(&Params.Category, "category", "vulnerabilityReport", "category of the policy report")

	if home := homedir.HomeDir(); home != "" {
		flag.StringVar(&Params.Kubeconfig, "kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		flag.StringVar(&Params.Kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()
}
