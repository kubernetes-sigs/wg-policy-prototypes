package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/trivy-adapter/pkg/cmd"
	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/trivy-adapter/pkg/imgvuln"
	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/trivy-adapter/pkg/params"
	"k8s.io/klog"

	// Load all known auth plugins
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var (
	// These variables are populated by GoReleaser via ldflags
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// main is the entrypoint of the imgvuln CLI executable command.
func main() {
	//parse arguments
	params.ParseArguments()

	defer klog.Flush()
	klog.InitFlags(nil)

	if err := cmd.Run(imgvuln.BuildInfo{
		Version:    version,
		Commit:     commit,
		Date:       date,
		Executable: executable(os.Args),
	}, os.Args, os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func executable(args []string) string {
	if strings.HasPrefix(filepath.Base(args[0]), "kubectl-") {
		return "kubectl trivy-adapter"
	}
	return "trivy-adapter"
}
