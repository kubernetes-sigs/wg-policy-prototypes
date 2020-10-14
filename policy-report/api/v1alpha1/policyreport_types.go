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
	Pass int `json:"pass"`

	// Fail provides the count of policies whose requirements were not met
	Fail int `json:"fail"`

	// Warn provides the count of unscored policies whose requirements were not met
	Warn int `json:"warn"`

	// Error provides the count of policies that could not be evaluated
	Error int `json:"error"`

	// Skip indicates the count of policies that were not selected for evaluation
	Skip int `json:"skip"`
}

// PolicyStatus has one of the following values:
//   - Pass: indicates that the policy requirements are met
//   - Fail: indicates that the policy requirements are not met
//   - Warn: indicates that the policy requirements and not met, and the policy is not scored
//   - Error: indicates that the policy could not be evaluated
//   - Skip: indicates that the policy was not selected based on user inputs or applicability
//
// +kubebuilder:validation:Enum=Pass;Fail;Warn;Error;Skip
type PolicyStatus string

// PolicyReportResult provides the result for an individual policy
type PolicyReportResult struct {

	// Policy is the name of the policy
	Policy string `json:"policy"`

	// Rule is the name of the policy rule
	// +optional
	Rule string `json:"rule,omitempty"`

	// Resources is an list of resources
	Resources []*ResourceStatus `json:"resources,omitempty"`

	// Message is a short user friendly description of the policy rule
	Message string `json:"message,omitempty"`

	// Scored indicates if this policy rule is scored
	Scored bool `json:"scored,omitempty"`

	// Data provides additional information for the policy rule
	Data map[string]string `json:"data,omitempty"`
}

// ResourceStatus provides the resource status
type ResourceStatus struct {

	// Resource is an optional reference to the resource check bu the policy rule
	Resource *corev1.ObjectReference `json:"resource,omitempty"`

	// ResourceSelector is an optional selector for policy results that apply to multiple resources.
	// For example, a policy result may apply to all pods that match a label.
	// Either a Resource or a ResourceSelector can be specified. If neither are provided, the
	// result is assumed to be for the policy report scope.
	// +optional
	ResourceSelector *metav1.LabelSelector `json:"resourceSelector,omitempty"`

	// Status indicates the result of the policy rule check
	Status PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Kind",type=string,JSONPath=`.scope.kind`,priority=1
// +kubebuilder:printcolumn:name="Name",type=string,JSONPath=`.scope.name`,priority=1
// +kubebuilder:printcolumn:name="Pass",type=integer,JSONPath=`.summary.pass`
// +kubebuilder:printcolumn:name="Fail",type=integer,JSONPath=`.summary.fail`
// +kubebuilder:printcolumn:name="Warn",type=integer,JSONPath=`.summary.warn`
// +kubebuilder:printcolumn:name="Error",type=integer,JSONPath=`.summary.error`
// +kubebuilder:printcolumn:name="Skip",type=integer,JSONPath=`.summary.skip`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

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

// +kubebuilder:object:root=true

// PolicyReportList contains a list of PolicyReport
type PolicyReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyReport `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PolicyReport{}, &PolicyReportList{})
}
