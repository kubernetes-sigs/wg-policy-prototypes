package params

import (
	"time"
)

type KubeBenchArgs struct {
	Name               string
	Category           string
	Kubeconfig         string
	KubebenchYAML      string
	KubebenchImg       string
	KubebenchTargets   string
	KubebenchVersion   string
	KubebenchBenchmark string
	Timeout            time.Duration
}
