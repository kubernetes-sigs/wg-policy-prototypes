package params

import (
	"flag"
	"path/filepath"
	"time"

	"k8s.io/client-go/util/homedir"
)

var (
	Params KubeBenchArgs
)

func ParseArguments() {
	flag.StringVar(&Params.Name, "name", "kube-bench", "name of policy report")
	flag.StringVar(&Params.Category, "category", "CIS Benchmarks", "category of the policy report")
	flag.StringVar(&Params.Namespace, "namespace", "default", "namespace of kube-bench job")
	flag.StringVar(&Params.KubebenchYAML, "yaml", "job.yaml", "YAML for kube-bench job")
	flag.StringVar(&Params.KubebenchTargets, "kube-bench-targets", "master,node,etcd,policies", "targets for benchmark of kube-bench job")
	flag.StringVar(&Params.KubebenchVersion, "kube-bench-version", "", "specify the Kubernetes version for kube-bench job")
	flag.StringVar(&Params.KubebenchBenchmark, "kube-bench-benchmark", "", "specify the benchmark for kube-bench job")

	Params.KubebenchImg = *flag.String("kubebenchImg", "aquasec/kube-bench:latest", "kube-bench image used as part of this test")
	Params.Timeout = *flag.Duration("timeout", 10*time.Minute, "Test Timeout")

	if home := homedir.HomeDir(); home != "" {
		flag.StringVar(&Params.Kubeconfig, "kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		flag.StringVar(&Params.Kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()
}
