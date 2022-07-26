/*
Copyright 2022.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +name:"VectorSpec"
// +weight:"200"
type _hugoVectorSpec interface{} //nolint:deadcode,unused

// +name:"Vector"
// +version:"v1alpha1"
// +description:"Vector system configuration"
type _metaVectorSpec interface{} //nolint:deadcode,unused

// VectorSpec defines the desired state of Vector
type VectorSpec struct {
	// Reference to the vector system. Each of the `vectorRef`s can manage a agent daemonset and a aggregator statefulset.
	VectorRef      string          `json:"vectorRef,omitempty"`
	AgentSpec      *AgentSpec      `json:"agent,omitempty"`
	AggregatorSpec *AggregatorSpec `json:"aggregator,omitempty"`
}

// VectorStatus defines the observed state of Vector
type VectorStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=vectors,scope=Cluster,categories=vector-all
// +kubebuilder:storageversion

// Vector is the Schema for the vectors API
type Vector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VectorSpec   `json:"spec,omitempty"`
	Status VectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VectorList contains a list of Vector
type VectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vector `json:"items"`
}

const (
	DefaultAgentImageRepository                = "fluent/fluent-bit"
	DefaultAgentImageTag                       = "1.9.5"
	DefaultAggregatorImageRepository           = "ghcr.io/banzaicloud/fluentd"
	DefaultAggregatorImageTag                  = "v1.14.6-alpine-5"
	DefaultAggregatorBufferStorageVolumeName   = "aggregator-buffer"
	DefaultAggregatorVolumeModeImageRepository = "busybox"
	DefaultAggregatorVolumeModeImageTag        = "latest"
)

func init() {
	SchemeBuilder.Register(&Vector{}, &VectorList{})
}
