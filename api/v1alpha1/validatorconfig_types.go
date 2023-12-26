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

// ValidatorConfigSpec defines the desired state of ValidatorConfig
type ValidatorConfigSpec struct {
	Plugins []HelmRelease `json:"plugins,omitempty" yaml:"plugins,omitempty"`
	Sink    *Sink         `json:"sink,omitempty" yaml:"sink,omitempty"`
}

type Sink struct {
	// +kubebuilder:validation:Enum=alertmanager;slack
	Type string `json:"type" yaml:"type"`
	// Name of a K8s secret containing configuration details for the sink
	SecretName string `json:"secretName" yaml:"secretName"`
}

type HelmRelease struct {
	Chart  HelmChart `json:"chart" yaml:"chart"`
	Values string    `json:"values" yaml:"values"`
}

type HelmChart struct {
	Name                  string `json:"name" yaml:"name"`
	Repository            string `json:"repository" yaml:"repository"`
	Version               string `json:"version" yaml:"version"`
	CaFile                string `json:"caFile,omitempty" yaml:"caFile,omitempty"`
	InsecureSkipTlsVerify bool   `json:"insecureSkipVerify,omitempty" yaml:"insecureSkipVerify,omitempty"`
	AuthSecretName        string `json:"authSecretName,omitempty" yaml:"authSecretName,omitempty"`
}

// ValidatorConfigStatus defines the observed state of ValidatorConfig
type ValidatorConfigStatus struct {
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []ValidatorPluginCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

type ValidatorPluginCondition struct {
	// Type of condition in CamelCase or in foo.example.com/CamelCase.
	// Many .condition.type values are consistent across resources like Available, but because arbitrary conditions
	// can be useful (see .node.status.conditions), the ability to deconflict is important.
	// +required
	Type ConditionType `json:"type"`

	// Name of the Validator plugin.
	// +required
	PluginName string `json:"pluginName"`

	// Status of the condition, one of True, False, Unknown.
	// +required
	Status corev1.ConditionStatus `json:"status"`

	// A human readable message indicating details about the transition.
	// This field may be empty.
	// +optional
	Message string `json:"message,omitempty"`

	// Last time the condition transitioned from one status to another.
	// This should be when the underlying condition changed. If that is not known, then using the time when
	// the API field changed is acceptable.
	// +required
	LastTransitionTime metav1.Time `json:"lastUpdatedTime"`
}

// ConditionType is a valid value for Condition.Type.
type ConditionType string

// HelmChartDeployedCondition defines the helm chart deployed condition type that defines if the helm chart was deployed correctly.
const HelmChartDeployedCondition ConditionType = "HelmChartDeployed"

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ValidatorConfig is the Schema for the validatorconfigs API
type ValidatorConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ValidatorConfigSpec   `json:"spec,omitempty"`
	Status ValidatorConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ValidatorConfigList contains a list of ValidatorConfig
type ValidatorConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ValidatorConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ValidatorConfig{}, &ValidatorConfigList{})
}
