// package
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/kube-bench-adapter/pkg/report"

	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/kube-bench-adapter/pkg/kubebench"

	"k8s.io/client-go/util/homedir"
)

var (
	name               string
	namespace          string
	category           string
	kubeconfig         string
	kubebenchYAML      string
	kubebenchImg       string
	kubebenchTargets   string
	kubebenchVersion   string
	kubebenchBenchmark string
	timeout            time.Duration
)

func parseArguments() {
	flag.StringVar(&name, "name", "kube-bench", "name of policy report")
	flag.StringVar(&namespace, "namespace", "default", "namespace of the cluster")
	flag.StringVar(&category, "category", "CIS Benchmarks", "category of the policy report")
	flag.StringVar(&kubebenchYAML, "yaml", "job.yaml", "YAML for kube-bench job")
	flag.StringVar(&kubebenchTargets, "kube-bench-targets", "master,node,etcd,policies", "targets for benchmark of kube-bench job")
	flag.StringVar(&kubebenchVersion, "kube-bench-version", "", "specify the Kubernetes version for kube-bench job")
	flag.StringVar(&kubebenchBenchmark, "kube-bench-benchmark", "", "specify the benchmark for kube-bench job")

	kubebenchImg = *flag.String("kubebenchImg", "aquasec/kube-bench:latest", "kube-bench image used as part of this test")
	timeout = *flag.Duration("timeout", 10*time.Minute, "Test Timeout")

	if home := homedir.HomeDir(); home != "" {
		flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()
}

func main() {
	parseArguments()

	//run kube-bench job
	cis, err := kubebench.RunJob(kubeconfig, kubebenchYAML, kubebenchImg, kubebenchVersion, kubebenchBenchmark, kubebenchTargets, timeout)
	if err != nil {
		fmt.Printf("failed to run job of kube-bench: %v \n", err)
		os.Exit(-1)
	}

	// create policy report
	r, err := report.New(cis, name, category)
	if err != nil {
		fmt.Printf("failed to create policy reports: %v \n", err)
		os.Exit(-1)
	}

	// write policy report
	r, err = report.Write(r, namespace, kubeconfig)
	if err != nil {
		fmt.Printf("failed to create policy reports: %v \n", err)
		os.Exit(-1)
	}

	fmt.Printf("wrote policy report %s/%s \n", r.Namespace, r.Name)
}
