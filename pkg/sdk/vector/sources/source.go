package sources

type SourceType string

const (
	ApacheMetrics            SourceType = "apache_metrics"
	AWSECSMetrics            SourceType = "aws_ecs_metrics"
	AWSKinesisFirehose       SourceType = "aws_kinesis_firehose"
	AWSS3                    SourceType = "aws_s3"
	AWSSQS                   SourceType = "aws_sqs"
	DatadogAgent             SourceType = "datadog_agent"
	DemoLogs                 SourceType = "demo_logs"
	DNSStap                  SourceType = "dnstap"
	DockerLogs               SourceType = "docker_logs"
	EventStoreDBMetrics      SourceType = "eventstoredb_metrics"
	Exec                     SourceType = "exec"
	File                     SourceType = "file"
	Fluent                   SourceType = "fluent"
	GCPPubSub                SourceType = "gcp_pubsub"
	HerokuLogplex            SourceType = "heroku_logs"
	HostMetrics              SourceType = "host_metrics"
	HTTP                     SourceType = "http"
	InternalLogs             SourceType = "internal_logs"
	InternalMetrics          SourceType = "internal_metrics"
	JournalD                 SourceType = "journald"
	Kafka                    SourceType = "kafka"
	KubernetesLogs           SourceType = "kubernetes_logs"
	Logstash                 SourceType = "logstash"
	MongoDBMetrics           SourceType = "mongodb_metrics"
	NATS                     SourceType = "nats"
	NGINXMetrics             SourceType = "nginx_metrics"
	PostgreSQLMetrics        SourceType = "postgresql_metrics"
	PrometheusRemoteWrite    SourceType = "prometheus_remote_write"
	PrometheusScrape         SourceType = "prometheus_scrape"
	Redis                    SourceType = "redis"
	Socket                   SourceType = "socket"
	SplunkHTTPEventCollector SourceType = "splunk_hec"
	StatsD                   SourceType = "statsd"
	STDIN                    SourceType = "stdin"
	Syslog                   SourceType = "syslog"
	Vector                   SourceType = "vector"
)

// Controls how acknowledgements are handled by this source. These settings override
// the global acknowledgement settings. This setting is deprecated in favor of
// enabling acknowledgements in the destination sink.
type Acknowledgement struct {
	// Controls if the source will wait for destination sinks to deliver the events before acknowledging receipt.
	Enabled bool `json:"enabled,omitempty"`
}

type TLSConfig struct {
	// Require TLS for incoming connections. If this is set, an identity certificate is also required.
	Enabled bool `json:"enabled,omitempty"`
	// Absolute path to an additional CA certificate file, in DER or PEM format (X.509), or an in-line CA certificate in PEM format.
	CAFile string `json:"caFile,omitempty"`
	// The key name added to each event with the client certificate’s metadata.
	ClientMetadataKey string `json:"clientMetadataKey,omitempty"`
	// Absolute path to a certificate file used to identify this server, in DER or PEM format (X.509) or PKCS#12,
	// or an in-line certificate in PEM format. If this is set, and is not a PKCS#12 archive,
	// key_file must also be set. This is required if enabled is set to true.
	CRTFile string `json:"crtFile,omitempty"`
	// Absolute path to a private key file used to identify this server, in DER or PEM format (PKCS#8),
	// or an in-line private key in PEM format.
	KeyFile string `json:"keyFile,omitempty"`
	// Pass phrase used to unlock the encrypted key file. This has no effect unless key_file is set.
	KeyPass string `json:"keyPass,omitempty"`
	// If true, Vector will require a TLS certificate from the connecting host and terminate
	// the connection if the certificate is not valid. If false (the default), Vector will not
	// request a certificate from the client.
	VerifyCertificate bool `json:"verifyCertificate,omitempty"`
}

// TODO: need to understand how this works first
type Telemetry struct {
}
