package cmd

import (
	_ "embed"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/scheme"
)

var (
	//go:embed crd/vulnerabilityreports.crd.yaml
	vulnerabilityReportsCRD []byte
)

func GetVulnerabilityReportsCRD() (apiextensionsv1.CustomResourceDefinition, error) {
	return getCRDFromBytes(vulnerabilityReportsCRD)
}

func getCRDFromBytes(bytes []byte) (apiextensionsv1.CustomResourceDefinition, error) {
	var crd apiextensionsv1.CustomResourceDefinition
	_, _, err := scheme.Codecs.UniversalDecoder().Decode(bytes, nil, &crd)
	if err != nil {
		return apiextensionsv1.CustomResourceDefinition{}, err
	}
	return crd, nil
}