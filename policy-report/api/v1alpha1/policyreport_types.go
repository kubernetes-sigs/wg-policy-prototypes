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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PolicyReportSummary provides a status count summary
type PolicyReportSummary struct {

	// Pass provides the count of policies whose requirements were met
	// +optional
	Pass int `json:"pass"`

	// Fail provides the count of policies whose requirements were not met
	// +optional
	Fail int `json:"fail"`

	// Warn provides the count of unscored policies whose requirements were not met
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
//   - partial-pass: the policy requirements are met but requires manual assessment
//   - fail: the policy requirements are not met
//   - warn: the policy requirements and not met and the policy is not scored
//   - error: the policy could not be evaluated
//   - skip: the policy was not selected based on user inputs or applicability
//
// +kubebuilder:validation:Enum=pass;fail;warn;error;skip
type PolicyResult string

// PolicyImpact has one of the following values:
//   - high
//   - low
//   - medium
// +kubebuilder:validation:Enum=high;low;medium
type PolicyImpact string

// PolicySeverity has one of the following values:
//   - high
//   - low
//   - medium
// +kubebuilder:validation:Enum=high;low;medium
type PolicySeverity string

// PolicyStatus has one of the following values:
//   - completed: indicates that the policy requirements are implemented
//   - partial: indicates that the policy requirements are partially implemented
//   - planned: indicates that the policy requirements are not implemented
// +kubebuilder:validation:Enum=completed;partial;planned
type PolicyStatus string

// PolicyReportResult provides the result for an individual policy
type PolicyReportResult struct {

	// Policy is the name of the policy
	Policy string `json:"policy"`

	// Rule is the name of the policy rule
	// +optional
	Rule string `json:"rule,omitempty"`

	// Subjects is an optional reference to the checked Kubernetes resources
	// +optional
	Subjects []*corev1.ObjectReference `json:"resources,omitempty"`

	// SubjectSelector is an optional label selector for checked Kubernetes resources.
	// For example, a policy result may apply to all pods that match a label.
	// Either a Subject or a SubjectSelector can be specified. If neither are provided, the
	// result is assumed to be for the policy report scope.
	// +optional
	SubjectSelector *metav1.LabelSelector `json:"resourceSelector,omitempty"`

	// Description is a short user friendly description of the policy rule
	Description string `json:"message,omitempty"`

	// Result indicates the result of the policy rule check
	Result PolicyResult `json:"result,omitempty"`

	// Scored indicates if this policy rule is scored
	Scored bool `json:"scored,omitempty"`

	// Data provides additional information for the policy rule
	Data map[string]string `json:"data,omitempty"`

	// Category indicates policy category
	// +optional
	Category string `json:"category,omitempty"`

	// Impact indicates policy criticality (pre-assessment)
	// +optional
	Impact PolicyImpact `json:"impact,omitempty"`

	// Severity indicates policy check result criticality (post-assessment)
	// +optional
	Severity PolicySeverity `json:"severity,omitempty"`

	// Status indicates the policy implementation readiness
	// +optional
	Status PolicyStatus `json:"status,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
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

	// Scope is an optional reference to the report scope (e.g. a Deployment, Namespace, or Node)
	// +optional
	Scope *corev1.ObjectReference `json:"scope,omitempty"`

	// ScopeSelector is an optional selector for multiple scopes (e.g. Pods).
	// Either one of, or none of, but not both of, Scope or ScopeSelector should be specified.
	// +optional
	ScopeSelector *metav1.LabelSelector `json:"scopeSelector,omitempty"`

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
