package report

import (
	//"strconv"
	//"strings"

	clusterpolicyreport "github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/image-vuln-adapter/pkg/apis/wgpolicyk8s.io/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func New(name string, category string) (*clusterpolicyreport.ClusterPolicyReport, error) {

	report := &clusterpolicyreport.ClusterPolicyReport{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Summary: clusterpolicyreport.PolicyReportSummary{
			Pass: 2,
			Fail: 4,
			Warn: 6,
		},
	}

	r := newResult(category)
	report.Results = append(report.Results, r)

	return report, nil
}

func newResult(category string) *clusterpolicyreport.PolicyReportResult {
	return &clusterpolicyreport.PolicyReportResult{
		Policy:   "vulnerabilityReport",
		Rule:     "no rule",
		Category: category,
		Result:   "pass",
		//Timestamp:   metav1.Timestamp,
		Severity:    "low",
		Scored:      true,
		Description: "Testing Policy Report",
		Properties: map[string]string{
			"index":           "0000-000000-0000-000000",
			"audit":           "test",
			"AuditEnv":        "test",
			"AuditConfig":     "test",
			"type":            "test",
			"remediation":     "test",
			"test_info":       "test",
			"actual_value":    "test",
			"IsMultiple":      "test",
			"expected_result": "test",
			"reason":          "test",
		},
	}
}
