package report

import (
	"strconv"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kubebench "github.com/aquasecurity/kube-bench/check"
	policyreport "github.com/mritunjaysharma394/policy-report-prototype/pkg/apis/wgpolicyk8s.io/v1alpha2"
)

func New(cisResults *kubebench.OverallControls, name string, category string) (*policyreport.PolicyReport, error) {

	report := &policyreport.PolicyReport{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Summary: policyreport.PolicyReportSummary{
			Pass: cisResults.Totals.Pass,
			Fail: cisResults.Totals.Fail,
			Warn: cisResults.Totals.Warn,
		},
	}

	for _, control := range cisResults.Controls {
		for _, group := range control.Groups {
			for _, check := range group.Checks {
				_ = check
				r := newResult(category, control, group, check)
				report.Results = append(report.Results, r)
			}
		}
	}

	return report, nil
}

func newResult(category string, control *kubebench.Controls, group *kubebench.Group, check *kubebench.Check) *policyreport.PolicyReportResult {
	return &policyreport.PolicyReportResult{
		Policy:      control.Text,
		Rule:        group.Text,
		Category:    category,
		Result:      convertState(check.State),
		Scored:      check.Scored,
		Description: check.Text,
		Properties: map[string]string{
			"index":           check.ID,
			"audit":           check.Audit,
			"AuditEnv":        check.AuditEnv,
			"AuditConfig":     check.AuditConfig,
			"type":            check.Type,
			"remediation":     check.Remediation,
			"test_info":       check.TestInfo[0],
			"actual_value":    check.ActualValue,
			"IsMultiple":      strconv.FormatBool(check.IsMultiple),
			"expected_result": check.ExpectedResult,
			"reason":          check.Reason,
		},
	}
}

func convertState(s kubebench.State) policyreport.PolicyResult {

	str := strings.ToLower(string(s))

	return policyreport.PolicyResult(str)
}
