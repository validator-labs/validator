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

// Package v1alpha1 contains API Schema definitions for the validation v1alpha1 API group
// +kubebuilder:object:generate=true
// +groupName=validation.spectrocloud.labs
package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// APIVersion is the API version used to reference v1alpha1 objects.
	APIVersion = GroupVersion.String()

	// Group is the API group used to reference validator objects.
	Group = "validation.spectrocloud.labs"

	// GroupVersion is group version used to register these objects.
	GroupVersion = schema.GroupVersion{Group: Group, Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme

	// ValidatorConfigKind is the kind of the ValidatorConfig object.
	ValidatorConfigKind = reflect.TypeOf(ValidatorConfig{}).Name()

	// ValidatorConfigGroupResource is the name of the ValidatorConfig resource.
	ValidatorConfigGroupResource = schema.GroupResource{Group: Group, Resource: "validatorconfigs"}

	// ValidationResultKind is the kind of the ValidationResult object.
	ValidationResultKind = reflect.TypeOf(ValidationResult{}).Name()

	// ValidationResultGroupResource is the name of the ValidationResult resource.
	ValidationResultGroupResource = schema.GroupResource{Group: Group, Resource: "validationresults"}
)
