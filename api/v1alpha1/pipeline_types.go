package v1alpha1

import (
	"github.com/spaghettifunk/vector-operator/pkg/sdk/vector/sources"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +name:"Pipeline"
// +weight:"200"
type _hugoPipeline interface{} //nolint:deadcode,unused

// +name:"Pipeline"
// +version:"v1alpha1"
// +description:"PipelineSpec Source Transformer Sink"
type _metaPipeline interface{} //nolint:deadcode,unused

// +kubebuilder:object:generate=true

type PipelineSpec struct {
	VectorRef  string       `json:"vectorRef,omitempty"`
	Sources    []*Source    `json:"sources,omitempty"`
	Transforms []*Transform `json:"transforms,omitempty"`
	Sinks      []*Sink      `json:"sinks,omitempty"`
}

type Source struct {
	Name              string                     `json:"name,omitempty"`
	ApacheMetricsSpec *sources.ApacheMetricsSpec `json:"apacheMetrics,omitempty"`
	DemoLogsSpec      *sources.DemoLogsSpec      `json:"demoLogs,omitempty"`
	VectorSpec        *sources.VectorSpec        `json:"vector,omitempty"`
}

type Transform struct {
	Name string `json:"name,omitempty"`
}

type Sink struct {
	Name string `json:"name,omitempty"`
}

// PipelineStatus defines the observed state of Pipeline
type PipelineStatus struct {
	Active        *bool    `json:"active,omitempty"`
	Problems      []string `json:"problems,omitempty"`
	ProblemsCount int      `json:"problemsCount,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=vector-all
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Active",type="boolean",JSONPath=".status.active",description="Is the output active?"
// +kubebuilder:printcolumn:name="Problems",type="integer",JSONPath=".status.problemsCount",description="Number of problems"
// +kubebuilder:storageversion

// Pipeline is the Schema for the pipelines API
type Pipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PipelineSpec   `json:"spec,omitempty"`
	Status PipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PipelineList contains a list of Pipeline
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pipeline `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Pipeline{}, &PipelineList{})
}
