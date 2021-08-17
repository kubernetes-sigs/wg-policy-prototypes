package imgvuln

const (
	// NamespaceName the name of the namespace in which imgvuln stores its
	// configuration and where it runs scan jobs.
	NamespaceName = "imgvuln"

	// ServiceAccountName the name of the service account used to provide
	// identity for scan jobs run by imgvuln.
	ServiceAccountName = "imgvuln"

	// ConfigMapName the name of the ConfigMap where imgvuln stores its
	// configuration.
	ConfigMapName = "imgvuln"

	// SecretName the name of the secret where imgvuln stores is sensitive
	// configuration.
	SecretName = "imgvuln"
)

const (
	LabelResourceKind      = "imgvuln.resource.kind"
	LabelResourceName      = "imgvuln.resource.name"
	LabelResourceNamespace = "imgvuln.resource.namespace"
	LabelContainerName     = "imgvuln.container.name"
	LabelPodSpecHash       = "pod-spec-hash"
	LabelPluginConfigHash  = "plugin-config-hash"

	LabelVulnerabilityReportScanner = "vulnerabilityReport.scanner"

	LabelK8SAppManagedBy = "app.kubernetes.io/managed-by"
	Appimgvuln         = "imgvuln"
)

const (
	AnnotationContainerImages    = "imgvuln.container-images"
	AnnotationScanJobAnnotations = "scanJob.annotations"
)