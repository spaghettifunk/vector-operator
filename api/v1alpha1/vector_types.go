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
	"fmt"

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
	VectorRef         string             `json:"vectorRef,omitempty"`
	AgentSpec         *AgentSpec         `json:"agent,omitempty"`
	AggregatorSpec    *AggregatorSpec    `json:"aggregator,omitempty"`
	GlobalOptionsSpec *GlobalOptionsSpec `json:"globalOptions,omitempty"`
	// Limit namespaces to watch Transform and Sink custom resources.
	WatchNamespaces []string `json:"watchNamespaces,omitempty"`
	// Namespace for cluster wide configuration resources like Transforms and Sinks.
	// This should be a protected namespace from regular users.
	// Resources like agent and aggregator will run in this namespace as well.
	ControlNamespace string `json:"controlNamespace"`
	// Allow configuration of cluster resources from any namespace. Mutually exclusive with ControlNamespace restriction of Cluster resources
	AllowClusterResourcesFromAllNamespaces bool `json:"allowClusterResourcesFromAllNamespaces,omitempty"`
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
	DefaultAgentImageRepository                = "timberio/vector"
	DefaultAgentImageTag                       = "0.23.0-alpine"
	DefaultAggregatorImageRepository           = "timberio/vector"
	DefaultAggregatorImageTag                  = "0.23.0-alpine"
	DefaultAggregatorBufferStorageVolumeName   = "aggregator-buffer"
	DefaultAggregatorVolumeModeImageRepository = "busybox"
	DefaultAggregatorVolumeModeImageTag        = "latest"
)

// QualifiedName is the "logging-resource" name combined
func (v *Vector) QualifiedName(name string) string {
	return fmt.Sprintf("%s-%s", v.Name, name)
}

func init() {
	SchemeBuilder.Register(&Vector{}, &VectorList{})
}
