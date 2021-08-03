package plugin

import (
	"fmt"

	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/image-vuln-adapter/pkg/vulnerabilityreport"
	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/image-vuln-adapter/pkg/ext"
	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/image-vuln-adapter/pkg/plugin/trivy"
	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/image-vuln-adapter/pkg/imgvuln"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Resolver struct {
	buildInfo          imgvuln.BuildInfo
	config             imgvuln.ConfigData
	namespace          string
	serviceAccountName string
	client             client.Client
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) WithBuildInfo(buildInfo imgvuln.BuildInfo) *Resolver {
	r.buildInfo = buildInfo
	return r
}

func (r *Resolver) WithConfig(config imgvuln.ConfigData) *Resolver {
	r.config = config
	return r
}

func (r *Resolver) WithNamespace(namespace string) *Resolver {
	r.namespace = namespace
	return r
}

func (r *Resolver) WithServiceAccountName(name string) *Resolver {
	r.serviceAccountName = name
	return r
}

func (r *Resolver) WithClient(client client.Client) *Resolver {
	r.client = client
	return r
}

// GetVulnerabilityPlugin is a factory method that instantiates the vulnerabilityreport.Plugin.
//
// imgvuln currently supports Trivy scanner in Standalone and ClientServer
// mode, and Aqua Enterprise scanner.
//
// You could add your own scanner by implementing the vulnerabilityreport.Plugin interface.
func (r *Resolver) GetVulnerabilityPlugin() (vulnerabilityreport.Plugin, imgvuln.PluginContext, error) {
	scanner, err := r.config.GetVulnerabilityReportsScanner()
	if err != nil {
		return nil, nil, err
	}

	pluginContext := imgvuln.NewPluginContext().
		WithName(string(scanner)).
		WithNamespace(r.namespace).
		WithServiceAccountName(r.serviceAccountName).
		WithClient(r.client).
		Get()

	switch scanner {
	case imgvuln.Trivy:
		return trivy.NewPlugin(ext.NewSystemClock(), ext.NewGoogleUUIDGenerator()), pluginContext, nil
	}
	return nil, nil, fmt.Errorf("unsupported vulnerability scanner plugin: %s", scanner)
}