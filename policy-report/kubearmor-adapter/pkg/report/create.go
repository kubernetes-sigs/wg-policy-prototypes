package report

import (
	"strconv"

	pb "github.com/kubearmor/KubeArmor/protobuf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/wg-policy-prototypes/policy-report/pkg/api/wgpolicyk8s.io/v1alpha2"
)

const PolicyReportSource string = "KubeArmor Policy Engine"

//slice of policy reports
var policyReports = make(map[string]*v1alpha2.PolicyReport)

func New(alert *pb.Alert, namespace string) (*v1alpha2.PolicyReport, error) {

	//policy report doesn't exist and needs to be created
	if policyReports[namespace] == nil {

		policyReports[namespace] = &v1alpha2.PolicyReport{
			ObjectMeta: metav1.ObjectMeta{
				Name: "kubearmor-policy-report",
			},
			Summary: v1alpha2.PolicyReportSummary{
				Fail: 0,
			},
		}
	}
	r := newResult(alert)
	if r.Result == "fail" {
		policyReports[namespace].Summary.Fail++
	}
	policyReports[namespace].Results = append(policyReports[namespace].Results, r)
	return policyReports[namespace], nil
}

func newResult(Alert *pb.Alert) *v1alpha2.PolicyReportResult {

	var sev string

	if Alert.Severity == "1" || Alert.Severity == "2" {
		sev = "low"
	} else if Alert.Severity == "3" || Alert.Severity == "4" || Alert.Severity == "5" {
		sev = "medium"
	} else {
		sev = "high"
	}
	return &v1alpha2.PolicyReportResult{

		Source:      PolicyReportSource,
		Policy:      Alert.PolicyName,
		Scored:      false,
		Timestamp:   metav1.Timestamp{Seconds: Alert.Timestamp, Nanos: int32(Alert.Timestamp)},
		Severity:    v1alpha2.PolicyResultSeverity(sev),
		Result:      "fail",
		Description: Alert.Message,
		Category:    Alert.Type,
		Properties: map[string]string{
			"updated_time":   Alert.UpdatedTime,
			"cluster_name":   Alert.ClusterName,
			"host_name":      Alert.HostName,
			"namespace_name": Alert.NamespaceName,
			"pod_name":       Alert.PodName,
			"container_id":   Alert.ContainerID,
			"container_name": Alert.ContainerName,
			"host_pid":       strconv.Itoa(int(Alert.HostPID)),
			"ppid":           strconv.Itoa(int(Alert.PPID)),
			"pid":            strconv.Itoa(int(Alert.PID)),
			"tags":           Alert.Tags,
			"source":         Alert.Source,
			"operation":      Alert.Operation,
			"resource":       Alert.Resource,
			"data":           Alert.Data,
			"action":         Alert.Action,
			"result":         Alert.Result,
		},
	}
}
