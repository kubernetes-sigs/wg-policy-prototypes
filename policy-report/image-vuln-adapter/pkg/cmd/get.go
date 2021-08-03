package cmd

import (
	"io"

	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/image-vuln-adapter/pkg/imgvuln"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func NewGetCmd(buildInfo imgvuln.BuildInfo, cf *genericclioptions.ConfigFlags, outWriter io.Writer) *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get security reports",
	}
	getCmd.AddCommand(NewGetVulnerabilitiesCmd(buildInfo.Executable, cf, outWriter))
	getCmd.PersistentFlags().StringP("output", "o", "yaml", "Output format. One of yaml|json")

	return getCmd
}