package kubebench

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	kubebench "github.com/aquasecurity/kube-bench/check"
	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/kube-bench-adapter/pkg/params"
	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	yaml "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

func getClientSet(kubeconfigPath string) (*kubernetes.Clientset, error) {
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

	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		return nil, err
	}

	return clientset, nil

}
func RunJob(params *params.KubeBenchArgs) (*kubebench.OverallControls, error) {

	clientset, err := getClientSet(params.Kubeconfig)
	if err != nil {
		return nil, err
	}
	var jobName string
	jobName, err = deployJob(context.Background(), clientset, params)
	if err != nil {
		return nil, err
	}

	p, err := findPodForJob(context.Background(), clientset, params, jobName, params.Timeout)
	if err != nil {
		return nil, err
	}

	output, err := getPodLogs(context.Background(), clientset, jobName, p)
	if err != nil {
		return nil, err
	}

	err = clientset.BatchV1().Jobs(params.Namespace).Delete(context.Background(), jobName, metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}

	err = clientset.CoreV1().Pods(params.Namespace).Delete(context.Background(), p.Name, metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}

	controls, err := convert(output)
	if err != nil {
		return nil, err
	}

	return controls, nil

}

func deployJob(ctx context.Context, clientset *kubernetes.Clientset, params *params.KubeBenchArgs) (string, error) {

	jobYAML, err := embedYAMLs(params.KubebenchYAML)
	if err != nil {
		return "", err
	}

	decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(jobYAML), len(jobYAML))
	job := &batchv1.Job{}
	if err := decoder.Decode(job); err != nil {
		return "", err
	}
	jobName := job.GetName()
	job.Spec.Template.Spec.Containers[0].Image = params.KubebenchImg
	job.Spec.Template.Spec.Containers[0].Args = []string{"--json"}
	job.Spec.Template.Spec.Containers[0].Args = []string{"--version", params.KubebenchVersion}
	job.Spec.Template.Spec.Containers[0].Args = []string{"--benchmark", params.KubebenchBenchmark}
	job.Spec.Template.Spec.Containers[0].Args = []string{"run", "--targets", params.KubebenchTargets, "--json"}
	_, err = clientset.BatchV1().Jobs(params.Namespace).Create(ctx, job, metav1.CreateOptions{})

	return jobName, err
}

func findPodForJob(ctx context.Context, clientset *kubernetes.Clientset, params *params.KubeBenchArgs, jobName string, duration time.Duration) (*apiv1.Pod, error) {
	failedPods := make(map[string]struct{})
	selector := fmt.Sprintf("job-name=%s", jobName)
	timeout := time.After(duration)
	for {
		time.Sleep(3 * time.Second)
	podfailed:
		select {
		case <-timeout:
			return nil, fmt.Errorf("podList - timed out: no Pod found for Job %s", jobName)
		default:
			pods, err := clientset.CoreV1().Pods(params.Namespace).List(ctx, metav1.ListOptions{
				LabelSelector: selector,
			})
			if err != nil {
				return nil, err
			}
			fmt.Printf("Found (%d) pods\n", len(pods.Items))
			for _, cp := range pods.Items {
				if _, found := failedPods[cp.Name]; found {
					continue
				}

				if strings.HasPrefix(cp.Name, jobName) {
					fmt.Printf("pod (%s) - %#v\n", cp.Name, cp.Status.Phase)
					if cp.Status.Phase == apiv1.PodSucceeded {
						return &cp, nil
					}

					if cp.Status.Phase == apiv1.PodFailed {
						fmt.Printf("pod (%s) - %s - retrying...\n", cp.Name, cp.Status.Phase)
						fmt.Print(getPodLogs(ctx, clientset, jobName, &cp))
						failedPods[cp.Name] = struct{}{}
						break podfailed
					}
				}
			}
		}
	}
}

func getPodLogs(ctx context.Context, clientset *kubernetes.Clientset, jobName string, pod *apiv1.Pod) (string, error) {
	podLogOpts := apiv1.PodLogOptions{}
	req := clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOpts)
	podLogs, err := req.Stream(ctx)
	if err != nil {
		return "", err
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func convert(jsonString string) (*kubebench.OverallControls, error) {
	jsonDataReader := strings.NewReader(jsonString)
	decoder := json.NewDecoder(jsonDataReader)

	var controls kubebench.OverallControls
	if err := decoder.Decode(&controls); err != nil {
		return nil, err
	}

	return &controls, nil
}
