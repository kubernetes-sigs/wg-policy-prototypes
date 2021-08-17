package cmd

import (
	"flag"
	"fmt"
	"io"

	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/trivy-adapter/pkg/imgvuln"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

const (
	shortMessage = "Trivy adapter security toolkit"
	longMessage  = `Trivy adapter security toolkit
imgvuln CLI can be used to find risks, such as vulnerabilities or insecure
pod descriptors, in Kubernetes workloads. By default, the risk assessment
reports are stored in the policy report crd.
$ kubectl create deployment nginx --image nginx:1.16
Run the vulnerability scanner to generate policy reports:
$ %[1]s scan policyreports deployment/nginx
`
)

func NewRootCmd(buildInfo imgvuln.BuildInfo, args []string, outWriter io.Writer, errWriter io.Writer) *cobra.Command {
	var cf *genericclioptions.ConfigFlags

	rootCmd := &cobra.Command{
		Use:           "imgvuln",
		Short:         shortMessage,
		Long:          fmt.Sprintf(longMessage, buildInfo.Executable),
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cf = genericclioptions.NewConfigFlags(true)

	rootCmd.AddCommand(NewInitCmd(buildInfo, cf))
	rootCmd.AddCommand(NewScanCmd(buildInfo, cf))

	SetGlobalFlags(cf, rootCmd)

	rootCmd.SetArgs(args[1:])
	rootCmd.SetOut(outWriter)
	rootCmd.SetErr(errWriter)

	return rootCmd
}

// Run is the entry point of the imgvuln CLI. It runs the specified
// command based on the specified args.
func Run(version imgvuln.BuildInfo, args []string, outWriter io.Writer, errWriter io.Writer) error {

	initFlags()

	return NewRootCmd(version, args, outWriter, errWriter).Execute()
}

func initFlags() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	// Hide all klog flags except for -v 
	flag.CommandLine.VisitAll(func(f *flag.Flag) {
		if f.Name != "v" {
			pflag.Lookup(f.Name).Hidden = true
		}
	})
}
