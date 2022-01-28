package report

import (
	"strconv"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kubebench "github.com/aquasecurity/kube-bench/check"
	clusterpolicyreport "sigs.k8s.io/wg-policy-prototypes/policy-report/pkg/api/wgpolicyk8s.io/v1alpha2"
)

const PolicyReportSource string = "Kube Bench"

func New(cisResults *kubebench.OverallControls, name string, category string) (*clusterpolicyreport.ClusterPolicyReport, error) {

	report := &clusterpolicyreport.ClusterPolicyReport{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Summary: clusterpolicyreport.PolicyReportSummary{
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

func newResult(category string, control *kubebench.Controls, group *kubebench.Group, check *kubebench.Check) *clusterpolicyreport.PolicyReportResult {
	return &clusterpolicyreport.PolicyReportResult{
		Policy:      check.ID + " " + control.Text,
		Rule:        group.Text,
		Source:      PolicyReportSource,
		Category:    category,
		Result:      convertState(check.State),
		Scored:      check.Scored,
		Description: check.Text,
		Timestamp:   metav1.Timestamp{Seconds: int64(time.Now().Second()), Nanos: int32(time.Now().Nanosecond())},
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

func convertState(s kubebench.State) clusterpolicyreport.PolicyResult {

	str := strings.ToLower(string(s))

	return clusterpolicyreport.PolicyResult(str)
}
