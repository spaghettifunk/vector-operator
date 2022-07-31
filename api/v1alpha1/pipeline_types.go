package v1alpha1

// +name:"Pipeline"
// +weight:"200"
type _hugoPipeline interface{} //nolint:deadcode,unused

// +name:"Pipeline"
// +version:"v1alpha1"
// +description:"PipelineSpec Source Transformer Sink"
type _metaPipeline interface{} //nolint:deadcode,unused

type PipelineSpec struct {
	Sources    []*Source    `json:"sources,omitempty"`
	Transforms []*Transform `json:"transforms,omitempty"`
	Sinks      []*Sink      `json:"sinks,omitempty"`
}

type Source struct {
	Name string `json:"name,omitempty"`
}

type Transform struct {
	Name string `json:"name,omitempty"`
}

type Sink struct {
	Name string `json:"name,omitempty"`
}
