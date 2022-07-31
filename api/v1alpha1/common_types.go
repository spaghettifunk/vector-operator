package v1alpha1

import (
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	corev1 "k8s.io/api/core/v1"
)

// +name:"Common"
// +weight:"200"
type _hugoCommon interface{} //nolint:deadcode,unused

// +name:"Common"
// +version:"v1alpha1"
// +description:"ImageSpec Metrics Security"
type _metaCommon interface{} //nolint:deadcode,unused

// ImageSpec struct hold information about image specification
type ImageSpec struct {
	Repository       string                        `json:"repository,omitempty"`
	Tag              string                        `json:"tag,omitempty"`
	PullPolicy       string                        `json:"pullPolicy,omitempty"`
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
}

func (s ImageSpec) RepositoryWithTag() string {
	res := s.Repository
	if s.Tag != "" {
		res += ":" + s.Tag
	}
	return res
}

// Metrics defines the service monitor endpoints
type Metrics struct {
	Interval              string               `json:"interval,omitempty"`
	Timeout               string               `json:"timeout,omitempty"`
	Port                  int32                `json:"port,omitempty"`
	Path                  string               `json:"path,omitempty"`
	ServiceMonitor        bool                 `json:"serviceMonitor,omitempty"`
	ServiceMonitorConfig  ServiceMonitorConfig `json:"serviceMonitorConfig,omitempty"`
	PrometheusAnnotations bool                 `json:"prometheusAnnotations,omitempty"`
	PrometheusRules       bool                 `json:"prometheusRules,omitempty"`
}

// ServiceMonitorConfig defines the ServiceMonitor properties
type ServiceMonitorConfig struct {
	AdditionalLabels   map[string]string   `json:"additionalLabels,omitempty"`
	HonorLabels        bool                `json:"honorLabels,omitempty"`
	Relabelings        []*v1.RelabelConfig `json:"relabelings,omitempty"`
	MetricsRelabelings []*v1.RelabelConfig `json:"metricRelabelings,omitempty"`
	Scheme             string              `json:"scheme,omitempty"`
	TLSConfig          *v1.TLSConfig       `json:"tlsConfig,omitempty"`
}

// Security defines Fluentd, Fluentbit deployment security properties
type Security struct {
	ServiceAccount               string                     `json:"serviceAccount,omitempty"`
	RoleBasedAccessControlCreate *bool                      `json:"roleBasedAccessControlCreate,omitempty"`
	PodSecurityPolicyCreate      bool                       `json:"podSecurityPolicyCreate,omitempty"`
	SecurityContext              *corev1.SecurityContext    `json:"securityContext,omitempty"`
	PodSecurityContext           *corev1.PodSecurityContext `json:"podSecurityContext,omitempty"`
}

// ReadinessDefaultCheck Enable default readiness checks
type ReadinessDefaultCheck struct {
	// Enable default Readiness check it'll fail if the buffer volume free space exceeds the `readinessDefaultThreshold` percentage (90%).
	BufferFreeSpace          bool  `json:"bufferFreeSpace,omitempty"`
	BufferFreeSpaceThreshold int32 `json:"bufferFreeSpaceThreshold,omitempty"`
	BufferFileNumber         bool  `json:"bufferFileNumber,omitempty"`
	BufferFileNumberMax      int32 `json:"bufferFileNumberMax,omitempty"`
	InitialDelaySeconds      int32 `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds           int32 `json:"timeoutSeconds,omitempty"`
	PeriodSeconds            int32 `json:"periodSeconds,omitempty"`
	SuccessThreshold         int32 `json:"successThreshold,omitempty"`
	FailureThreshold         int32 `json:"failureThreshold,omitempty"`
}
