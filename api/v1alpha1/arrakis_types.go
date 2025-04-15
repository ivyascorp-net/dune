/*
Copyright 2025.

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
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ArrakisSpec defines the desired state of Arrakis.
type ArrakisSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// System

}

type System struct {
	Name           string            `json:"name,omitempty"`
	Image          string            `json:"image,omitempty"`
	Replicas       string            `json:"replicas,omitempty"`
	AppSettings    appsv1.Deployment `json:"appSettings,omitempty"`
	ServiceAccount v1.ServiceAccount `json:"serviceAccount,omitempty"`
	Service        v1.Service        `json:"service,omitempty"`
}
type ArrakisPhase string

const (
	ArrakisPhasePending  ArrakisPhase = "Pending"
	ArrakisPhaseDeployed ArrakisPhase = "Deployed"
	ArrakisPhaseFailed   ArrakisPhase = "Failed"
)

// ArrakisStatus defines the observed state of Arrakis.
type ArrakisStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Phase ArrakisPhase
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Arrakis is the Schema for the arrakis API.
type Arrakis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArrakisSpec   `json:"spec,omitempty"`
	Status ArrakisStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ArrakisList contains a list of Arrakis.
type ArrakisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Arrakis `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Arrakis{}, &ArrakisList{})
}
