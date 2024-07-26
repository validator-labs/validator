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

// ValidatorConfigSpec defines the desired state of ValidatorConfig.
type ValidatorConfigSpec struct {
	// HelmConfig defines the configuration for the Helm repository.
	HelmConfig HelmConfig `json:"helmConfig" yaml:"helmConfig"`

	// Plugins defines the configuration for the validator plugins.
	Plugins []HelmRelease `json:"plugins,omitempty" yaml:"plugins,omitempty"`

	// Sink defines the configuration for the notification sink.
	Sink *Sink `json:"sink,omitempty" yaml:"sink,omitempty"`
}

// Sink defines the configuration for a notification sink.
type Sink struct {
	// Type of the sink.
	// +kubebuilder:validation:Enum=alertmanager;slack
	Type string `json:"type" yaml:"type"`

	// SecretName is the name of a K8s secret containing configuration details for the sink.
	SecretName string `json:"secretName" yaml:"secretName"`
}

// HelmRelease defines the configuration for a Helm chart release.
type HelmRelease struct {
	// Name of the Helm chart.
	Name string `json:"name" yaml:"name"`

	// Version of the Helm chart.
	Version string `json:"version" yaml:"version"`

	// Values defines the values to be passed to the Helm chart.
	Values string `json:"values" yaml:"values"`
}

// HelmConfig defines the configuration for a Helm repository.
type HelmConfig struct {
	// Repository URL of the Helm chart.
	Repository string `json:"repository" yaml:"repository"`

	// CAFile is the path to the CA certificate for the Helm repository.
	CAFile string `json:"caFile,omitempty" yaml:"caFile,omitempty"`

	// InsecureSkipTLSVerify skips the verification of the server's certificate chain and host name.
	InsecureSkipTLSVerify bool `json:"insecureSkipVerify,omitempty" yaml:"insecureSkipVerify,omitempty"`

	// AuthSecretName is the name of the K8s secret containing the authentication details for the Helm repository.
	AuthSecretName string `json:"authSecretName,omitempty" yaml:"authSecretName,omitempty"`
}

// ValidatorConfigStatus defines the observed state of ValidatorConfig
type ValidatorConfigStatus struct {
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []ValidatorPluginCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// ValidatorPluginCondition describes the state of a Validator plugin.
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

// HelmChartDeployedCondition defines whether the helm chart was installed/pulled/upgraded correctly.
const HelmChartDeployedCondition ConditionType = "HelmChartDeployed"

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ValidatorConfig is the Schema for the validatorconfigs API.
type ValidatorConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ValidatorConfigSpec   `json:"spec,omitempty"`
	Status ValidatorConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ValidatorConfigList contains a list of ValidatorConfig.
type ValidatorConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ValidatorConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ValidatorConfig{}, &ValidatorConfigList{})
}
