package cmd

import (
	"context"
	"fmt"
	"log"

	policyreport "sigs.k8s.io/wg-policy-prototypes/policy-report/api/v1alpha2"

	client "github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/trivy-adapter/pkg/generated/v1alpha2/clientset/versioned"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog"
)

func Write(r []policyreport.PolicyReport, kubeconfigPath string) error {

	var kubeconfig *rest.Config

	cfg, err := rest.InClusterConfig()
	if err != nil {
		cfg, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			klog.Fatalf("Error building kubeconfig: %s", err.Error())
			fmt.Println("error: ", err)
			return nil
		}
	}
	kubeconfig = cfg
	clientset, err := client.NewForConfig(kubeconfig)
	if err != nil {
		fmt.Println("error: ", err)
			return nil
	}
	
	for _, report := range r {
		var r *policyreport.PolicyReport
		r = &report
		policyReport := clientset.Wgpolicyk8sV1alpha2().PolicyReports(r.Namespace)
		// Check for existing Policy Reports
		result, getErr := policyReport.Get(context.Background(), r.Name, metav1.GetOptions{})
		fmt.Println(result)
		// Create new Policy Report if not found
		if errors.IsNotFound(getErr) {
			fmt.Println("creating policy report...")

			result, err = policyReport.Create(context.Background(), r, metav1.CreateOptions{})
			if err != nil {
				fmt.Println("error: ", err)
				return nil
			}
		} else {

			// Update existing Policy Report
			fmt.Println("updating policy report...")
			retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {

				getObj, err := policyReport.Get(context.Background(), r.GetName(), metav1.GetOptions{})
				if errors.IsNotFound(err) {
					// This doesnt ever happen even if it is already deleted or not found
					log.Printf("%v not found", report.GetName())
					return nil
				}

				if err != nil {
					return err
				}

				r.SetResourceVersion(getObj.GetResourceVersion())

				_, updateErr := policyReport.Update(context.Background(), r, metav1.UpdateOptions{})
				return updateErr
			})
			if retryErr != nil {
				panic(fmt.Errorf("update failed: %v", retryErr))
			}
			fmt.Println("updated policy report...")
		}
	}
	return nil
}
