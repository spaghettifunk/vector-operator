package sources

// +kubebuilder:object:generate=true

// ref: https://vector.dev/docs/reference/configuration/sources/demo_logs/
type DemoLogsSpec struct {
	Type SourceType `json:"type,omitempty"`
	// The total number of lines to output. By default the source continuously prints logs (infinitely).
	Count int `json:"count,omitempty"`
	// Configures in which way frames are decoded into events.
	Decoding *Decoding `json:"decoding,omitempty"`
	// The format of the randomly generated output.
	Format string `json:"format,omitempty"`
	// Configures in which way incoming byte sequences are split up into byte frames.
	Framing *Framing `json:"framing,omitempty"`
	// The amount of time, in seconds, to pause between each batch of output lines. The default is one
	// batch per second. In order to remove the delay and output batches as quickly as possible, set interval to 0.0.
	Interval string `json:"interval,omitempty"`
	// The list of lines to output.
	Lines []string `json:"lines,omitempty"`
	// If true, each output line starts with an increasing sequence number, beginning with 0.
	Sequence bool `json:"sequence,omitempty"`
}

type Decoding struct {
	// The decoding method.
	Codec string `json:"codec,omitempty"`
}

// +kubebuilder:object:generate=true

type Framing struct {
	// Options for character_delimited framing.
	// Relevant when: method = `character_delimited`
	CharacterDelimited *CharacterDelimited `json:"characterDelimited,omitempty"`
	// The framing method.
	Method string `json:"method,omitempty"`
	// Options for newline_delimited framing.
	// Relevant when: method = `newline_delimited`
	// The maximum frame length limit. Any frames longer than max_length bytes will be discarded entirely.
	NewLineDelimitedMaxLength int64 `json:"newLineDelimitedMaxLength,omitempty"`
	// Options for octet_counting framing.
	// Relevant when: method = `octet_counting`
	// The maximum frame length limit. Any frames longer than max_length bytes will be discarded entirely.
	OctetCountingMaxLength int64 `json:"octetCountingMaxLength,omitempty"`
}

type CharacterDelimited struct {
	// The character used to separate frames.
	Delimiter string `json:"delimiter,omitempty"`
	// The maximum frame length limit. Any frames longer than max_length bytes will be discarded entirely.
	MaxLegth int64 `json:"maxLength,omitempty"`
}
