package main

import (
	"fmt"
	"os"

	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/image-vuln-adapter/pkg/params"
	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/image-vuln-adapter/pkg/report"
)

func main() {

	//parse arguments
	params.ParseArguments()

	// create policy report
	r, err := report.New(params.Params.Name, params.Params.Category)
	if err != nil {
		fmt.Printf("failed to create policy reports: %v \n", err)
		os.Exit(-1)
	}

	// write policy report
	r, err = report.Write(r, params.Params.Kubeconfig)
	if err != nil {
		fmt.Printf("failed to create policy reports: %v \n", err)
		os.Exit(-1)
	}

	fmt.Printf("wrote policy report %s \n", r.Name)
	fmt.Println(r)
}