package v1alpha1

import (
	"github.com/banzaicloud/operator-tools/pkg/typeoverride"
	"github.com/banzaicloud/operator-tools/pkg/volume"
	corev1 "k8s.io/api/core/v1"
)

// +name:"AggregatorSpec"
// +weight:"200"
type _hugoAggregatorSpec interface{} //nolint:deadcode,unused

// +name:"AggregatorSpec"
// +version:"v1beta1"
// +description:"AggregatorSpec defines the desired state of Aggregator"
type _metaAggregatorSpec interface{} //nolint:deadcode,unused

// +kubebuilder:object:generate=true

// AggregatorSpec defines the desired state of Aggregator
type AggregatorSpec struct {
	Enable                 bool              `json:"enable,omitempty"`
	Pipeline               *PipelineSpec     `json:"pipeline,omitempty"`
	StatefulSetAnnotations map[string]string `json:"statefulsetAnnotations,omitempty"`
	Annotations            map[string]string `json:"annotations,omitempty"`
	ConfigCheckAnnotations map[string]string `json:"configCheckAnnotations,omitempty"`
	Labels                 map[string]string `json:"labels,omitempty"`
	EnvVars                []corev1.EnvVar   `json:"envVars,omitempty"`
	Image                  ImageSpec         `json:"image,omitempty"`
	DisablePvc             bool              `json:"disablePvc,omitempty"`
	// BufferStorageVolume is by default configured as PVC using AggregatorPvcSpec
	// +docLink:"volume.KubernetesVolume,https://github.com/banzaicloud/operator-tools/tree/master/docs/types"
	BufferStorageVolume   volume.KubernetesVolume     `json:"bufferStorageVolume,omitempty"`
	ExtraVolumes          []ExtraVolume               `json:"extraVolumes,omitempty"`
	Resources             corev1.ResourceRequirements `json:"resources,omitempty"`
	ConfigCheckResources  corev1.ResourceRequirements `json:"configCheckResources,omitempty"`
	LivenessProbe         *corev1.Probe               `json:"livenessProbe,omitempty"`
	LivenessDefaultCheck  bool                        `json:"livenessDefaultCheck,omitempty"`
	ReadinessProbe        *corev1.Probe               `json:"readinessProbe,omitempty"`
	ReadinessDefaultCheck ReadinessDefaultCheck       `json:"readinessDefaultCheck,omitempty"`
	Port                  int32                       `json:"port,omitempty"`
	Tolerations           []corev1.Toleration         `json:"tolerations,omitempty"`
	NodeSelector          map[string]string           `json:"nodeSelector,omitempty"`
	Affinity              *corev1.Affinity            `json:"affinity,omitempty"`
	Metrics               *Metrics                    `json:"metrics,omitempty"`
	Security              *Security                   `json:"security,omitempty"`
	Scaling               *AggregatorScaling          `json:"scaling,omitempty"`
	// +kubebuilder:validation:enum=fatal,error,warn,info,debug,trace
	LogLevel                string                       `json:"logLevel,omitempty"`
	ServiceAccountOverrides *typeoverride.ServiceAccount `json:"serviceAccount,omitempty"`
}

// +kubebuilder:object:generate=true

// ExtraVolume defines the Aggregator extra volumes
type ExtraVolume struct {
	VolumeName    string                   `json:"volumeName,omitempty"`
	Path          string                   `json:"path,omitempty"`
	ContainerName string                   `json:"containerName,omitempty"`
	Volume        *volume.KubernetesVolume `json:"volume,omitempty"`
}

func (e *ExtraVolume) GetVolume() (corev1.Volume, error) {
	return e.Volume.GetVolume(e.VolumeName)
}

func (e *ExtraVolume) ApplyVolumeForPodSpec(spec *corev1.PodSpec) error {
	return e.Volume.ApplyVolumeForPodSpec(e.VolumeName, e.ContainerName, e.Path, spec)
}

// +kubebuilder:object:generate=true

// AggregatorScaling enables configuring the scaling behaviour of the Aggregator statefulset
type AggregatorScaling struct {
	Replicas            int    `json:"replicas,omitempty"`
	PodManagementPolicy string `json:"podManagementPolicy,omitempty"`
}
