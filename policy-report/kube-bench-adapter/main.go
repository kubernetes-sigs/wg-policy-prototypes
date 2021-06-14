// package
package main

import (
	"fmt"
	"os"

	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/kube-bench-adapter/pkg/report"

	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/kube-bench-adapter/pkg/kubebench"
	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/kube-bench-adapter/pkg/params"
)

func main() {
	//parse arguments
	params.ParseArguments()

	//run kube-bench job
	cis, err := kubebench.RunJob(&params.Params)
	if err != nil {
		fmt.Printf("failed to run job of kube-bench: %v \n", err)
		os.Exit(-1)
	}

	// create policy report
	r, err := report.New(cis, params.Params.Name, params.Params.Category)
	if err != nil {
		fmt.Printf("failed to create policy reports: %v \n", err)
		os.Exit(-1)
	}

	// write policy report
	r, err = report.Write(r, params.Params.Namespace, params.Params.Kubeconfig)
	if err != nil {
		fmt.Printf("failed to create policy reports: %v \n", err)
		os.Exit(-1)
	}

	fmt.Printf("wrote policy report %s/%s \n", r.Namespace, r.Name)
}
