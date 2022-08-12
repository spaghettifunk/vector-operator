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

	util "github.com/banzaicloud/operator-tools/pkg/utils"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
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
	// Resources like agent and aggregator will run in this namespace as welv.
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

// SetDefaults fills empty attributes
func (v *Vector) SetDefaults() error {
	if v.Spec.AggregatorSpec != nil { // nolint:nestif
		if v.Spec.AggregatorSpec.Image.Repository == "" {
			v.Spec.AggregatorSpec.Image.Repository = DefaultAgentImageRepository
		}
		if v.Spec.AggregatorSpec.Image.Tag == "" {
			v.Spec.AggregatorSpec.Image.Tag = DefaultAgentImageTag
		}
		if v.Spec.AggregatorSpec.Image.PullPolicy == "" {
			v.Spec.AggregatorSpec.Image.PullPolicy = "IfNotPresent"
		}
		if v.Spec.AggregatorSpec.Annotations == nil {
			v.Spec.AggregatorSpec.Annotations = make(map[string]string)
		}
		if v.Spec.AggregatorSpec.Metrics != nil {
			if v.Spec.AggregatorSpec.Metrics.Path == "" {
				v.Spec.AggregatorSpec.Metrics.Path = "/metrics"
			}
			if v.Spec.AggregatorSpec.Metrics.Port == 0 {
				v.Spec.AggregatorSpec.Metrics.Port = 24231
			}
			if v.Spec.AggregatorSpec.Metrics.Timeout == "" {
				v.Spec.AggregatorSpec.Metrics.Timeout = "5s"
			}
			if v.Spec.AggregatorSpec.Metrics.Interval == "" {
				v.Spec.AggregatorSpec.Metrics.Interval = "15s"
			}
			if v.Spec.AggregatorSpec.Metrics.PrometheusAnnotations {
				v.Spec.AggregatorSpec.Annotations["prometheus.io/scrape"] = "true"

				v.Spec.AggregatorSpec.Annotations["prometheus.io/path"] = v.Spec.AggregatorSpec.Metrics.Path
				v.Spec.AggregatorSpec.Annotations["prometheus.io/port"] = fmt.Sprintf("%d", v.Spec.AggregatorSpec.Metrics.Port)
			}
		}
		if v.Spec.AggregatorSpec.Resources.Limits == nil {
			v.Spec.AggregatorSpec.Resources.Limits = v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("1G"),
				v1.ResourceCPU:    resource.MustParse("1"),
			}
		}
		if v.Spec.AggregatorSpec.Resources.Requests == nil {
			v.Spec.AggregatorSpec.Resources.Requests = v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("512M"),
				v1.ResourceCPU:    resource.MustParse("500m"),
			}
		}
		if v.Spec.AggregatorSpec.LivenessProbe == nil {
			if v.Spec.AggregatorSpec.LivenessDefaultCheck {
				v.Spec.AggregatorSpec.LivenessProbe = &v1.Probe{
					ProbeHandler: v1.ProbeHandler{
						Exec: &v1.ExecAction{Command: []string{"/bin/healthy.sh"}},
					},
					InitialDelaySeconds: 600,
					TimeoutSeconds:      0,
					PeriodSeconds:       60,
					SuccessThreshold:    0,
					FailureThreshold:    0,
				}
			}
		}
		if v.Spec.AggregatorSpec.ReadinessDefaultCheck.BufferFreeSpace {
			if v.Spec.AggregatorSpec.ReadinessDefaultCheck.BufferFreeSpaceThreshold == 0 {
				v.Spec.AggregatorSpec.ReadinessDefaultCheck.BufferFreeSpaceThreshold = 90
			}
		}
		if v.Spec.AggregatorSpec.ReadinessDefaultCheck.BufferFileNumber {
			if v.Spec.AggregatorSpec.ReadinessDefaultCheck.BufferFileNumberMax == 0 {
				v.Spec.AggregatorSpec.ReadinessDefaultCheck.BufferFileNumberMax = 5000
			}
		}
		if v.Spec.AggregatorSpec.ReadinessDefaultCheck.InitialDelaySeconds == 0 {
			v.Spec.AggregatorSpec.ReadinessDefaultCheck.InitialDelaySeconds = 5
		}
		if v.Spec.AggregatorSpec.ReadinessDefaultCheck.TimeoutSeconds == 0 {
			v.Spec.AggregatorSpec.ReadinessDefaultCheck.TimeoutSeconds = 3
		}
		if v.Spec.AggregatorSpec.ReadinessDefaultCheck.PeriodSeconds == 0 {
			v.Spec.AggregatorSpec.ReadinessDefaultCheck.PeriodSeconds = 30
		}
		if v.Spec.AggregatorSpec.ReadinessDefaultCheck.SuccessThreshold == 0 {
			v.Spec.AggregatorSpec.ReadinessDefaultCheck.SuccessThreshold = 3
		}
		if v.Spec.AggregatorSpec.ReadinessDefaultCheck.FailureThreshold == 0 {
			v.Spec.AggregatorSpec.ReadinessDefaultCheck.FailureThreshold = 1
		}
		if v.Spec.AggregatorSpec.MountPath == "" {
			v.Spec.AggregatorSpec.MountPath = "/etc/vector/"
		}
	}

	if v.Spec.AgentSpec != nil { // nolint:nestif
		if v.Spec.AgentSpec.Image.Repository == "" {
			v.Spec.AgentSpec.Image.Repository = DefaultAggregatorImageRepository
		}
		if v.Spec.AgentSpec.Image.Tag == "" {
			v.Spec.AgentSpec.Image.Tag = DefaultAggregatorImageTag
		}
		if v.Spec.AgentSpec.Image.PullPolicy == "" {
			v.Spec.AgentSpec.Image.PullPolicy = "IfNotPresent"
		}
		if v.Spec.AgentSpec.LogLevel == "" {
			v.Spec.AgentSpec.LogLevel = "info"
		}
		if v.Spec.AgentSpec.Resources.Limits == nil {
			v.Spec.AgentSpec.Resources.Limits = v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("1G"),
				v1.ResourceCPU:    resource.MustParse("1"),
			}
		}
		if v.Spec.AgentSpec.Resources.Requests == nil {
			v.Spec.AgentSpec.Resources.Requests = v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("512M"),
				v1.ResourceCPU:    resource.MustParse("500m"),
			}
		}
		if v.Spec.AgentSpec.Annotations == nil {
			v.Spec.AgentSpec.Annotations = make(map[string]string)
		}
		if v.Spec.AgentSpec.Metrics != nil {
			if v.Spec.AgentSpec.Metrics.Path == "" {
				v.Spec.AgentSpec.Metrics.Path = "/api/v1/metrics/prometheus"
			}
			if v.Spec.AgentSpec.Metrics.Port == 0 {
				v.Spec.AgentSpec.Metrics.Port = 2020
			}
			if v.Spec.AgentSpec.Metrics.Timeout == "" {
				v.Spec.AgentSpec.Metrics.Timeout = "5s"
			}
			if v.Spec.AgentSpec.Metrics.Interval == "" {
				v.Spec.AgentSpec.Metrics.Interval = "15s"
			}
			if v.Spec.AgentSpec.Metrics.PrometheusAnnotations {
				v.Spec.AgentSpec.Annotations["prometheus.io/scrape"] = "true"
				v.Spec.AgentSpec.Annotations["prometheus.io/path"] = v.Spec.AgentSpec.Metrics.Path
				v.Spec.AgentSpec.Annotations["prometheus.io/port"] = fmt.Sprintf("%d", v.Spec.AgentSpec.Metrics.Port)
			}
		} else if v.Spec.AgentSpec.LivenessDefaultCheck {
			v.Spec.AgentSpec.Metrics = &Metrics{
				Port: 2020,
				Path: "/",
			}
		}
		if v.Spec.AgentSpec.LivenessProbe == nil {
			if v.Spec.AgentSpec.LivenessDefaultCheck {
				v.Spec.AgentSpec.LivenessProbe = &v1.Probe{
					ProbeHandler: v1.ProbeHandler{
						HTTPGet: &v1.HTTPGetAction{
							Path: v.Spec.AgentSpec.Metrics.Path,
							Port: intstr.IntOrString{
								IntVal: v.Spec.AgentSpec.Metrics.Port,
							},
						}},
					InitialDelaySeconds: 10,
					TimeoutSeconds:      0,
					PeriodSeconds:       10,
					SuccessThreshold:    0,
					FailureThreshold:    3,
				}
			}
		}
		if v.Spec.AgentSpec.MountPath == "" {
			v.Spec.AgentSpec.MountPath = "/etc/vector/"
		}
		if v.Spec.AgentSpec.Security == nil {
			v.Spec.AgentSpec.Security = &Security{}
		}
		if v.Spec.AgentSpec.Security.RoleBasedAccessControlCreate == nil {
			v.Spec.AgentSpec.Security.RoleBasedAccessControlCreate = util.BoolPointer(true)
		}
	}

	return nil
}

// AggregatorObjectMeta creates an objectMeta for resource Aggregator
func (v *Vector) AggregatorObjectMeta(name, component string) metav1.ObjectMeta {
	o := metav1.ObjectMeta{
		Name:      v.QualifiedName(name),
		Namespace: v.Spec.ControlNamespace,
		Labels:    v.GetAggregatorLabels(component),
		OwnerReferences: []metav1.OwnerReference{
			{
				APIVersion: v.APIVersion,
				Kind:       v.Kind,
				Name:       v.Name,
				UID:        v.UID,
				Controller: util.BoolPointer(true),
			},
		},
	}
	return o
}

func (v *Vector) GetAggregatorLabels(component string) map[string]string {
	return util.MergeLabels(
		v.Spec.AggregatorSpec.Labels,
		map[string]string{
			"app.kubernetes.io/name":      "aggregator",
			"app.kubernetes.io/component": component,
		},
		GenerateVectorRefLabels(v.ObjectMeta.GetName()),
	)
}

func GenerateVectorRefLabels(vectorRef string) map[string]string {
	return map[string]string{"app.kubernetes.io/managed-by": vectorRef}
}
