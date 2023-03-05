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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ContextSpec defines the desired state of Context
type ContextSpec struct {
	// Provider is name of the provider to reference the provider in the same namespace
	Provider *LocalObjectRef `json:"provider"`
	// Owner represents the owner of the context
	Owner *Owner `json:"owner"`
	// +immutable
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Name is immutable"
	// +kubebuilder:validation:MaxLength=512
	Name string `json:"name"`
	// Env represents environment variable of the context
	Env []*EnvVarFrom `json:"env,omitempty"`
	// EnvFrom is set of environment variables from the secrets
	EnvFrom *EnvFrom `json:"envFrom,omitempty"`
	// DeleteRemoteWhenDelete is flag to delete remote when delete object
	DeleteRemoteWhenDelete bool `json:"deleteRemoteWhenDelete,omitempty"`
}

// ContextStatus defines the observed state of Context
type ContextStatus struct {
	// UUID of the context
	Id string `json:"id,omitempty"`
	// EnvVars contains the environment name and it's bcrypted value to compare when the context is updated
	Env map[string][]byte `json:"env,omitempty"`
	// Created timestamp of the context
	CreatedAt string `json:"createdAt,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Context is the Schema for the contexts API
type Context struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContextSpec   `json:"spec,omitempty"`
	Status ContextStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ContextList contains a list of Context
type ContextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Context `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Context{}, &ContextList{})
}
