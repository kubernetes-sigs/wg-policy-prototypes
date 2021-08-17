package cmd

import (
	"context"

	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/trivy-adapter/pkg/imgvuln"
	"github.com/spf13/cobra"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewInitCmd(buildInfo imgvuln.BuildInfo, cf *genericclioptions.ConfigFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Create Kubernetes resources used by imgvuln",
		Long: `Create all the resources used by imgvuln. It will create the following in your
Kubernetes cluster:
 - RBAC objects:
   - The "imgvuln" ClusterRole
   - The "imgvuln" ClusterRoleBinding
 - The "imgvuln" namespace with the following objects:
   - The "imgvuln" service account
   - The "imgvuln" ConfigMap
   - The "imgvuln" secret
The "imgvuln" ConfigMap and the "imgvuln" secret contain the default
config parameters.
All resources created by this command can be removed from the cluster using
the "cleanup" command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			kubeConfig, err := cf.ToRESTConfig()
			if err != nil {
				return err
			}
			kubeClientset, err := kubernetes.NewForConfig(kubeConfig)
			if err != nil {
				return err
			}
			apiExtensionsClientset, err := apiextensionsv1.NewForConfig(kubeConfig)
			if err != nil {
				return err
			}
			scheme := imgvuln.NewScheme()
			kubeClient, err := client.New(kubeConfig, client.Options{Scheme: scheme})
			if err != nil {
				return err
			}
			configManager := imgvuln.NewConfigManager(kubeClientset, imgvuln.NamespaceName)
			installer := NewInstaller(buildInfo, kubeClientset, apiExtensionsClientset, kubeClient, configManager)
			return installer.Install(context.Background())
		},
	}
	return cmd
}