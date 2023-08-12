/*
Copyright 2020 The Kubernetes authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StatusFilter is used by PolicyReport generators to write only those reports whose status is specified by the filters
// +kubebuilder:validation:Enum=pass;fail;warn;error;skip
type StatusFilter string

type Limits struct {
	// MaxResults is the maximum number of results contained in the report
	// +optional
	MaxResults int `json:"maxResults"`

	// StatusFilter indicates that the PolicyReport contains only those reports with statuses specified in this list
	// +kubebuilder:validation:UniqueItems=true
	// +optional
	StatusFilter []*StatusFilter `json:"statusFilter,omitempty"`
}

type PolicyReportConfiguration struct {
	Limits Limits `json:"limits"`
}

// PolicyReportSummary provides a status count summary
type PolicyReportSummary struct {

	// Pass provides the count of policies whose requirements were met
	// +optional
	Pass int `json:"pass"`

	// Fail provides the count of policies whose requirements were not met
	// +optional
	Fail int `json:"fail"`

	// Warn provides the count of non-scored policies whose requirements were not met
	// +optional
	Warn int `json:"warn"`

	// Error provides the count of policies that could not be evaluated
	// +optional
	Error int `json:"error"`

	// Skip indicates the count of policies that were not selected for evaluation
	// +optional
	Skip int `json:"skip"`
}

// PolicyResult has one of the following values:
//   - pass: the policy requirements are met
//   - fail: the policy requirements are not met
//   - warn: the policy requirements are not met and the policy is not scored
//   - error: the policy could not be evaluated
//   - skip: the policy was not selected based on user inputs or applicability
//
// +kubebuilder:validation:Enum=pass;fail;warn;error;skip
type PolicyResult string

// PolicyResultSeverity has one of the following values:
//   - critical
//   - high
//   - low
//   - medium
//   - info
//
// +kubebuilder:validation:Enum=critical;high;low;medium;info
type PolicyResultSeverity string

// PolicyReportResult provides the result for an individual policy
type PolicyReportResult struct {

	// Source is an identifier for the policy engine that manages this report
	// If the Source is specified at this level, it will override the Source
	// field set at the PolicyReport level
	// +optional
	Source string `json:"source"`

	// Policy is the name or identifier of the policy
	Policy string `json:"policy"`

	// Rule is the name or identifier of the rule within the policy
	// +optional
	Rule string `json:"rule,omitempty"`

	// Category indicates policy category
	// +optional
	Category string `json:"category,omitempty"`

	// Severity indicates policy check result criticality
	// +optional
	Severity PolicyResultSeverity `json:"severity,omitempty"`

	// Timestamp indicates the time the result was found
	Timestamp metav1.Timestamp `json:"timestamp,omitempty"`

	// Result indicates the outcome of the policy rule execution
	Result PolicyResult `json:"result,omitempty"`

	// Scored indicates if this result is scored
	Scored bool `json:"scored,omitempty"`

	// Subjects is an optional reference to the checked Kubernetes resources
	// +optional
	Subjects []*corev1.ObjectReference `json:"resources,omitempty"`

	// ResourceSelector is an optional label selector for checked Kubernetes resources.
	// For example, a policy result may apply to all pods that match a label.
	// Either a Subject or a ResourceSelector can be specified. If neither are provided, the
	// result is assumed to be for the policy report scope.
	// +optional
	ResourceSelector *metav1.LabelSelector `json:"resourceSelector,omitempty"`

	// Description is a short user friendly message for the policy rule
	Description string `json:"message,omitempty"`

	// Properties provides additional information for the policy rule
	Properties map[string]string `json:"properties,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Kind",type=string,JSONPath=`.scope.kind`,priority=1
// +kubebuilder:printcolumn:name="Name",type=string,JSONPath=`.scope.name`,priority=1
// +kubebuilder:printcolumn:name="Pass",type=integer,JSONPath=`.summary.pass`
// +kubebuilder:printcolumn:name="Fail",type=integer,JSONPath=`.summary.fail`
// +kubebuilder:printcolumn:name="Warn",type=integer,JSONPath=`.summary.warn`
// +kubebuilder:printcolumn:name="Error",type=integer,JSONPath=`.summary.error`
// +kubebuilder:printcolumn:name="Skip",type=integer,JSONPath=`.summary.skip`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:shortName=polr

// PolicyReport is the Schema for the policyreports API
type PolicyReport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Source is an identifier for the source e.g. a policy engine that manages this report.
	// Use this field if all the results are produced by a single policy engine.
	// If the results are produced by multiple sources e.g. different engines or scanners,
	// then use the Source field at the PolicyReportResult level.
	// +optional
	Source string `json:"source"`

	// Scope is an optional reference to the report scope (e.g. a Deployment, Namespace, or Node)
	// +optional
	Scope *corev1.ObjectReference `json:"scope,omitempty"`

	// ScopeSelector is an optional selector for multiple scopes (e.g. Pods).
	// Either one of, or none of, but not both of, Scope or ScopeSelector should be specified.
	// +optional
	ScopeSelector *metav1.LabelSelector `json:"scopeSelector,omitempty"`

	// Configuration is an optional field which can be used to specify
	// a contract between PolicyReport generators and consumers
	// +optional
	Configuration *PolicyReportConfiguration `json:"configuration,omitempty"`

	// PolicyReportSummary provides a summary of results
	// +optional
	Summary PolicyReportSummary `json:"summary,omitempty"`

	// PolicyReportResult provides result details
	// +optional
	Results []*PolicyReportResult `json:"results,omitempty"`
}

// PolicyReportList contains a list of PolicyReport
// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PolicyReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyReport `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PolicyReport{}, &PolicyReportList{})
}
