/*
Copyright 2024.

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

// ParseKeyWordSpec defines the desired state of ParseKeyWord
type ParseKeyWordSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ParseKeyWord. Edit parsekeyword_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ParseKeyWordStatus defines the observed state of ParseKeyWord
type ParseKeyWordStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ParseKeyWord is the Schema for the parsekeywords API
type ParseKeyWord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ParseKeyWordSpec   `json:"spec,omitempty"`
	Status ParseKeyWordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ParseKeyWordList contains a list of ParseKeyWord
type ParseKeyWordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ParseKeyWord `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ParseKeyWord{}, &ParseKeyWordList{})
}
