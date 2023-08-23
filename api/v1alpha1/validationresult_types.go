/*
Copyright 2023.

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

// ValidationResultSpec defines the desired state of ValidationResult
type ValidationResultSpec struct {
	Plugin string `json:"plugin"`
}

type ValidationState string

const (
	ValidationFailed     ValidationState = "Failed"
	ValidationInProgress ValidationState = "InProgress"
	ValidationSucceeded  ValidationState = "Succeeded"
)

// ValidationResultStatus defines the observed state of ValidationResult
type ValidationResultStatus struct {
	State ValidationState `json:"state"`

	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []ValidationCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

type ValidationCondition struct {
	Status             corev1.ConditionStatus `json:"status"`
	LastTransitionTime metav1.Time            `json:"lastTransitionTime,omitempty"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// Human-readable message indicating details about last transition.
	Message  string   `json:"message,omitempty"`
	Failures []string `json:"failures,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="Age"
//+kubebuilder:printcolumn:name="Plugin",type="string",JSONPath=".spec.plugin",description="Plugin"
//+kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state",description="State"

// ValidationResult is the Schema for the validationresults API
type ValidationResult struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ValidationResultSpec   `json:"spec,omitempty"`
	Status ValidationResultStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ValidationResultList contains a list of ValidationResult
type ValidationResultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ValidationResult `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ValidationResult{}, &ValidationResultList{})
}
