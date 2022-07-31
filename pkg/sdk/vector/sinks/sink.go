package sinks

type SinkType string

const (
	AWSCloudwatchLogs               SinkType = "aws_cloudwatch_logs"
	AWSCloudwatchMetrics            SinkType = "aws_cloudwatch_metrics"
	AWSKinesisFirehose              SinkType = "aws_kinesis_firehose"
	AWSKinesisStream                SinkType = "aws_kinesis_streams"
	AWSS3                           SinkType = "aws_s3"
	AWSSQS                          SinkType = "aws_sqs"
	AzureBlobStorage                SinkType = "azure_blob"
	AzureMonitorLogs                SinkType = "azure_monitor_logs"
	Blackhole                       SinkType = "blackhole"
	Clickhouse                      SinkType = "clickhouse"
	Console                         SinkType = "console"
	DatadogEvents                   SinkType = "datadog_events"
	DatadogLogs                     SinkType = "datadog_logs"
	DatadogMetrics                  SinkType = "datadog_metrics"
	DatadogTraces                   SinkType = "datadog_traces"
	Elasticsearch                   SinkType = "elasticsearch"
	File                            SinkType = "file"
	GCPCloudMonitoring              SinkType = "gcp_stackdriver_metrics"
	GCPCloudStorage                 SinkType = "gcp_cloud_storage"
	GCPOperationsLogs               SinkType = "gcp_stackdriver_logs"
	GCPPubSub                       SinkType = "gcp_pubsub"
	Honeycomb                       SinkType = "honeycomb"
	HTTP                            SinkType = "http"
	HumioLogs                       SinkType = "humio_logs"
	HumioMetrics                    SinkType = "humio_metrics"
	InfluxDBLogs                    SinkType = "influxdb_logs"
	InfluxDBMetrics                 SinkType = "influxdb_metrics"
	Kafka                           SinkType = "kafka"
	LogDNA                          SinkType = "logdna"
	Loki                            SinkType = "loki"
	NATS                            SinkType = "nats"
	NewRelic                        SinkType = "new_relic"
	NewRelicLogs                    SinkType = "new_relic_logs"
	Papertrail                      SinkType = "papertrail"
	PrometheusExporter              SinkType = "prometheus_exporter"
	PrometheusRemoteWrite           SinkType = "prometheus_remote_write"
	Pulsar                          SinkType = "pulsar"
	Redis                           SinkType = "redis"
	SematextLogs                    SinkType = "sematext_logs"
	SematextMetrics                 SinkType = "sematext_metrics"
	Socket                          SinkType = "socket"
	SplunkHTTPEventCollectorLogs    SinkType = "splunk_hec_logs"
	SplunkHTTPEventCollectorMetrics SinkType = "splunk_hec_metrics"
	StatsD                          SinkType = "statsd"
	Vector                          SinkType = "vector"
	Websocket                       SinkType = "websocket"
)

// Controls how acknowledgements are handled by this sink. When enabled, all connected sources that
// support end-to-end acknowledgements will wait for the destination of this sink to acknowledge
// receipt of events before providing acknowledgement to the sending source. These settings override
// the global acknowledgement settings.
type Acknowledgement struct {
	// Controls if all connected sources will wait for this sink to deliver the events before acknowledging receipt.
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
