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
	"crypto"
	"encoding/base64"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ValidationResultSpec defines the desired state of ValidationResult
type ValidationResultSpec struct {
	Plugin string `json:"plugin"`
	// The number of rules in the validator plugin spec, hence the number of expected ValidationResults.
	// +kubebuilder:validation:Minimum=1
	ExpectedResults int `json:"expectedResults"`
}

type ValidationState string

const (
	ValidationFailed     ValidationState = "Failed"
	ValidationInProgress ValidationState = "InProgress"
	ValidationSucceeded  ValidationState = "Succeeded"
)

type SinkState string

const (
	SinkEmitNone      SinkState = "N/A"
	SinkEmitFailed    SinkState = "Failed"
	SinkEmitSucceeded SinkState = "Succeeded"
)

// ValidationResultStatus defines the observed state of ValidationResult
type ValidationResultStatus struct {
	State     ValidationState `json:"state"`
	SinkState SinkState       `json:"sinkState,omitempty"`

	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []ValidationCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

type ValidationCondition struct {
	// Unique, one-word description of the validation type associated with the condition.
	ValidationType string `json:"validationType"`
	// Unique, one-word description of the validation rule associated with the condition.
	ValidationRule string `json:"validationRule"`
	// Human-readable message indicating details about the last transition.
	Message string `json:"message,omitempty"`
	// Human-readable messages indicating additional details for the last transition.
	Details []string `json:"details,omitempty"`
	// Human-readable messages indicating additional failure details for the last transition.
	Failures []string `json:"failures,omitempty"`
	// True if the validation rule succeeded, otherwise False
	Status corev1.ConditionStatus `json:"status"`
	// Timestamp of most recent execution of the validation rule associated with the condition.
	LastValidationTime metav1.Time `json:"lastValidationTime"`
}

// DefaultValidationCondition returns a default ValidationCondition
func DefaultValidationCondition() ValidationCondition {
	return ValidationCondition{
		Details:            make([]string, 0),
		Status:             corev1.ConditionTrue,
		LastValidationTime: metav1.Now(),
	}
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

func (r *ValidationResult) Hash() string {
	digester := crypto.MD5.New()

	fmt.Fprint(digester, r.ObjectMeta.UID)
	fmt.Fprint(digester, r.Spec)
	fmt.Fprint(digester, r.Status.State)

	for _, condition := range r.Status.Conditions {
		c := condition.DeepCopy()
		c.LastValidationTime = metav1.Time{}
		fmt.Fprint(digester, c)
	}

	hash := digester.Sum(nil)
	return base64.StdEncoding.EncodeToString(hash)
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
