package sinks

// +kubebuilder:object:generate=true

type VectorSpec struct {
	Type SinkType `json:"type,omitempty"`
}
