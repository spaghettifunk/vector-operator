package v1alpha1

// +name:"Configuration"
// +weight:"200"
type _hugoConfiguration interface{} //nolint:deadcode,unused

// +name:"Configuration"
// +version:"v1alpha1"
// +description:"SourceSpec TransformerSpec SinkSpec GlobalOptionsSpec"
type _metaConfiguration interface{} //nolint:deadcode,unused

type GlobalOptionsSpec struct {
	// Controls how acknowledgements are handled by all sources. These settings may be overridden in individual sources.
	// +kubebuilder:default:="false"
	AcknowledgementsEnabled bool `json:"acknowledgementsEnabled"`
	// The directory used for persisting Vector state, such as on-disk buffers, file checkpoints, and more.
	// Please make sure the Vector project has write permissions to this directory.
	// +kubebuilder:default:="/var/lib/vector/"
	DataDir string `json:"dataDir"`
	// +optional
	EnrichmentTables `json:"enrichmentaTables"`
	// +optional
	Healthcheck `json:"healthcheck"`
	// +optional
	LogSchema `json:"logSchema"`
	// +optional
	Proxy `json:"proxy"`
	// The name of the time zone to apply to timestamp conversions that do not contain an explicit time zone.
	// The time zone name may be any name in the TZ database, or local to indicate system local time.
	// +kubebuilder:default:="local"
	Timezone string `json:"timezone"`
}

type EnrichmentTables struct {
	// +optional
	File `json:"file"`
}

type File struct {
	// +kubebuilder:default:="csv"
	Type string `json:"type"`
	// +kubebuilder:default:=","
	Delimiter string `json:"delimiter"`
	// +kubebuilder:default:="true"
	IncludeHeaders bool `json:"includeHeaders"`
	// The path of the enrichment table file.
	Path string `json:"path"`
	// ref: https://vector.dev/docs/reference/configuration/global-options/#enrichment_tables.file.schema
	// +kubebuilder:default:="csv"
	Schema map[string]interface{} `json:"schema"`
}

// Configures health checks for all sinks
type Healthcheck struct {
	// Disables all health checks if false, otherwise sink specific option overrides it.
	// +kubebuilder:default:="false"
	Enabled bool
	// Exit on startup if any sinks' health check fails
	// +kubebuilder:default:="false"
	RequireHealthy bool
}

// Configures default log schema for all events. This is used by Vector source components
// to assign the fields on incoming events.
type LogSchema struct {
	// Sets the event key to use for the event host field.
	// +kubebuilder:default:="host"
	HostKey string `json:"hostKey"`
	// Sets the event key to use for the event message field.
	// +kubebuilder:default:="message"
	MessageKey string `json:"messageKey"`
	// Sets the event key to use for event metadata field (e.g. error or abort annotations in the `remap` transform).
	// +kubebuilder:default:="metadata"
	MetadataKey string `json:"metadataKey"`
	// Sets the event key to use for the event source type field that is set by some sources.
	// +kubebuilder:default:="source_type"
	SourceTypeKey string `json:"sourceTypeKey"`
	// Sets the event key to use for the event timestamp field.
	// +kubebuilder:default:="timestamp"
	TimestampKey string `json:"timestampKey"`
}

// Configures an HTTP(S) proxy for Vector to use.
type Proxy struct {
	// Enable the proxy
	// +kubebuilder:default:="false"
	Enabled bool `json:"enabled"`
	// The URL to proxy HTTP requests through.
	// +optional
	HTTPURL string `json:"http"`
	// The URL to proxy HTTPS requests through.
	// +optional
	HTTPSURL string `json:"https"`
	// A list of hosts to avoid proxying. Ref: https://vector.dev/docs/reference/configuration/global-options/#proxy.no_proxy
	// +optional
	NoProxy []string `json:"noProxy"`
}
