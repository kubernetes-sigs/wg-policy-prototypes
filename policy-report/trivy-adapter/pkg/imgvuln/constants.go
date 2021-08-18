package imgvuln

const (
	// NamespaceName the name of the namespace in which imgvuln stores its
	// configuration and where it runs scan jobs.
	NamespaceName = "trivy-adapter"

	// ServiceAccountName the name of the service account used to provide
	// identity for scan jobs run by imgvuln.
	ServiceAccountName = "trivy-adapter"

	// ConfigMapName the name of the ConfigMap where imgvuln stores its
	// configuration.
	ConfigMapName = "trivy-adapter"

	// SecretName the name of the secret where imgvuln stores is sensitive
	// configuration.
	SecretName = "trivy-adapter"
)

const (
	LabelResourceKind      = "trivy-adapter.resource.kind"
	LabelResourceName      = "trivy-adapter.resource.name"
	LabelResourceNamespace = "trivy-adapter.resource.namespace"
	LabelContainerName     = "trivy-adapter.container.name"
	LabelPodSpecHash       = "pod-spec-hash"
	LabelPluginConfigHash  = "plugin-config-hash"

	LabelVulnerabilityReportScanner = "vulnerabilityReport.scanner"

	LabelK8SAppManagedBy = "app.kubernetes.io/managed-by"
	Appimgvuln         = "trivy-adapter"
)

const (
	AnnotationContainerImages    = "trivy-adapter.container-images"
	AnnotationScanJobAnnotations = "scanJob.annotations"
)