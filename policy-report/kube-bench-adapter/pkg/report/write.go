package report

import (
	"context"
	"fmt"
	"log"

	clusterpolicyreport "github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/kube-bench-adapter/pkg/apis/wgpolicyk8s.io/v1alpha2"

	client "github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/kube-bench-adapter/pkg/generated/v1alpha2/clientset/versioned"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog"
)

func Write(r *clusterpolicyreport.ClusterPolicyReport, kubeconfigPath string) (*clusterpolicyreport.ClusterPolicyReport, error) {
	var kubeconfig *rest.Config

	cfg, err := rest.InClusterConfig()
	if err != nil {
		cfg, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			klog.Fatalf("Error building kubeconfig: %s", err.Error())
			return nil, err
		}
	}
	kubeconfig = cfg
	clientset, err := client.NewForConfig(kubeconfig)
	if err != nil {
		return nil, err
	}
	policyReport := clientset.Wgpolicyk8sV1alpha2().ClusterPolicyReports()

	// Check for existing Policy Reports
	result, getErr := policyReport.Get(context.Background(), r.Name, metav1.GetOptions{})
	// Create new Policy Report if not found
	if errors.IsNotFound(getErr) {
		fmt.Println("creating cluster-wide policy report...")

		result, err = policyReport.Create(context.Background(), r, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	} else {

		// Update existing Policy Report
		fmt.Println("updating policy report...")
		retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {

			getObj, err := policyReport.Get(context.Background(), r.GetName(), metav1.GetOptions{})
			if errors.IsNotFound(err) {
				// This doesnt ever happen even if it is already deleted or not found
				log.Printf("%v not found", r.GetName())
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

	return result, nil
}
