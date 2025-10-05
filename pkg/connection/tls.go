package connection

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"time"
)

// TLSConfig holds TLS connection configuration.
type TLSConfig struct {
	Host           string
	Port           string
	SkipVerify     bool
	ConnectTimeout time.Duration
}

// TestDirectTLS tests a direct TLS connection (like HTTPS).
// Returns peer certificates, connection state, and error.
func TestDirectTLS(config TLSConfig) ([]*x509.Certificate, tls.ConnectionState, error) {
	tlsConfig := &tls.Config{
		ServerName:         config.Host,
		InsecureSkipVerify: config.SkipVerify,
	}

	dialer := &net.Dialer{
		Timeout: config.ConnectTimeout,
	}

	conn, err := tls.DialWithDialer(dialer, "tcp", net.JoinHostPort(config.Host, config.Port), tlsConfig)
	if err != nil {
		return nil, tls.ConnectionState{}, fmt.Errorf("failed to establish TLS connection: %w", err)
	}
	defer conn.Close()

	// Get certificate information
	state := conn.ConnectionState()

	return state.PeerCertificates, state, nil
}