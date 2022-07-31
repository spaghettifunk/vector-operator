package sources

// https://vector.dev/docs/reference/configuration/sources/vector/
type VectorSpec struct {
	Type SourceType `json:"type,omitempty"`
	// The HTTP address to listen for connections on. It must include a port.
	Address string `json:"address,omitempty"`
	// Source API version. Specifying this version ensures that Vector does not break backward compatibility.
	Version string `json:"version,omitempty"`
	// Configures the TCP keepalive behavior for the connection to the source.
	KeepAlive *KeepAlive `json:"keepAlive,omitempty"`
	// Configures the receive buffer size using the SO_RCVBUF option on the socket.
	ReceiveBufferBytes int64 `json:"receiveBufferBytes,omitempty"`
	// The timeout before a connection is forcefully closed during shutdown.
	ShutdownTimeoutSecs int64 `json:"shutdownTimeoutSecs,omitempty"`
	// Configures the TLS options for incoming connections.
	TLSConfig *TLSConfig `json:"tls,omitempty"`
}

type KeepAlive struct {
	// The time a connection needs to be idle before sending TCP keepalive probes.
	TimeSecs int64 `json:"timeSecs,omitempty"`
}
