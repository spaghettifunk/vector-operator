package v1alpha1

import "github.com/spaghettifunk/vector-operator/pkg/sdk/vector/sources"

// +name:"Pipeline"
// +weight:"200"
type _hugoPipeline interface{} //nolint:deadcode,unused

// +name:"Pipeline"
// +version:"v1alpha1"
// +description:"PipelineSpec Source Transformer Sink"
type _metaPipeline interface{} //nolint:deadcode,unused

// +kubebuilder:object:generate=true

type PipelineSpec struct {
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
