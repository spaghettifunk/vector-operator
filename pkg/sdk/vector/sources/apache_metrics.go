package sources

// +kubebuilder:object:generate=true

// ref: https://vector.dev/docs/reference/configuration/sources/apache_metrics/
type ApacheMetricsSpec struct {
	Type SourceType `json:"type,omitempty"`
	// mod_status endpoints to scrape metrics from.
	Endpoints []string `json:"endpoints,omitempty"`
	// The namespace of the metric. Disabled if empty.
	Namespace string `json:"namespace,omitempty"`
	// Configures an HTTP(S) proxy for Vector to use. By default, the globally configured proxy is used
	Proxy *Proxy `json:"proxy,omitempty"`
}

// +kubebuilder:object:generate=true

type Proxy struct {
	// If false the proxy will be disabled.
	Enabled bool `json:"enabled,omitempty"`
	// The URL to proxy HTTP requests through.
	HTTP string `json:"http,omitempty"`
	// The URL to proxy HTTPS requests through.
	HTTPS string `json:"https,omitempty"`
	// A list of hosts to avoid proxying
	// see: https://vector.dev/docs/reference/configuration/sources/apache_metrics/#proxy.no_proxy
	NoProxy []string `json:"noProxy,omitempty"`
	// The interval between scrapes.
	ScrapeIntervalSecs int64 `json:"scrapeIntervalSecs,omitempty"`
}
