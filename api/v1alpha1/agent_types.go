package v1alpha1

import (
	// "github.com/banzaicloud/operator-tools/pkg/typeoverride"
	// "github.com/banzaicloud/operator-tools/pkg/volume"
	"strconv"

	"github.com/banzaicloud/operator-tools/pkg/typeoverride"
	corev1 "k8s.io/api/core/v1"
)

// +name:"AgentSpec"
// +weight:"200"
type _hugoAgentSpec interface{} //nolint:deadcode,unused

// +name:"AgentSpec"
// +version:"v1beta1"
// +description:"AgentSpec defines the desired state of Agent"
type _metaAgentSpec interface{} //nolint:deadcode,unused

// +kubebuilder:object:generate=true

// AgentSpec defines the desired state of Agent
type AgentSpec struct {
	Pipeline                *PipelineSpec                `json:"pipeline,omitempty"`
	DaemonSetAnnotations    map[string]string            `json:"daemonsetAnnotations,omitempty"`
	Annotations             map[string]string            `json:"annotations,omitempty"`
	Labels                  map[string]string            `json:"labels,omitempty"`
	EnvVars                 []corev1.EnvVar              `json:"envVars,omitempty"`
	Image                   ImageSpec                    `json:"image,omitempty"`
	TargetHost              string                       `json:"targetHost,omitempty"`
	TargetPort              int32                        `json:"targetPort,omitempty"`
	LogLevel                string                       `json:"logLevel,omitempty" plugin:"default:info"`
	Resources               corev1.ResourceRequirements  `json:"resources,omitempty"`
	Tolerations             []corev1.Toleration          `json:"tolerations,omitempty"`
	NodeSelector            map[string]string            `json:"nodeSelector,omitempty"`
	Affinity                *corev1.Affinity             `json:"affinity,omitempty"`
	Metrics                 *Metrics                     `json:"metrics,omitempty"`
	Security                *Security                    `json:"security,omitempty"`
	MountPath               string                       `json:"mountPath,omitempty"`
	ExtraVolumeMounts       []*VolumeMount               `json:"extraVolumeMounts,omitempty"`
	LivenessProbe           *corev1.Probe                `json:"livenessProbe,omitempty"`
	LivenessDefaultCheck    bool                         `json:"livenessDefaultCheck,omitempty"`
	ReadinessProbe          *corev1.Probe                `json:"readinessProbe,omitempty"`
	ServiceAccountOverrides *typeoverride.ServiceAccount `json:"serviceAccount,omitempty"`
}

// GetPrometheusPortFromAnnotation gets the port value from annotation
func (spec AgentSpec) GetPrometheusPortFromAnnotation() int32 {
	var err error
	var port int64
	if spec.Annotations != nil {
		port, err = strconv.ParseInt(spec.Annotations["prometheus.io/port"], 10, 32)
		if err != nil {
			panic(err)
		}
	}
	return int32(port)
}

// +kubebuilder:object:generate=true

// VolumeMount defines source and destination folders of a hostPath type pod mount
type VolumeMount struct {
	// Source folder
	// +kubebuilder:validation:Pattern=^/.+$
	Source string `json:"source"`
	// Destination Folder
	// +kubebuilder:validation:Pattern=^/.+$
	Destination string `json:"destination"`
	// Mount Mode
	ReadOnly *bool `json:"readOnly,omitempty"`
}
