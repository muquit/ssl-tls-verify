package connection

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/muquit/ssl-tls-verify/pkg/utils"
)

// StartTLSConfig holds StartTLS connection configuration.
type StartTLSConfig struct {
	Host           string
	Port           string
	SkipVerify     bool
	ConnectTimeout time.Duration
}

// TestStartTLS tests a StartTLS connection (like SMTP).
// Returns peer certificates, connection state, and error.
func TestStartTLS(config StartTLSConfig) ([]*x509.Certificate, tls.ConnectionState, error) {
	// First, establish a plain connection
	dialer := &net.Dialer{
		Timeout: config.ConnectTimeout,
	}

	conn, err := dialer.Dial("tcp", net.JoinHostPort(config.Host, config.Port))
	if err != nil {
		return nil, tls.ConnectionState{}, fmt.Errorf("failed to establish plain connection: %w", err)
	}
	defer conn.Close()

	fmt.Println("Plain connection established, attempting StartTLS...")

	// For SMTP, we need to do the protocol handshake manually to maintain control
	// over the connection for certificate inspection

	// Read server greeting
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, tls.ConnectionState{}, fmt.Errorf("failed to read server greeting: %w", err)
	}
	greeting := string(buffer[:n])
	fmt.Printf("Server greeting: %s", utils.MaskIPAddresses(greeting))

	// Send EHLO command
	ehloCmd := fmt.Sprintf("EHLO %s\r\n", config.Host)
	_, err = conn.Write([]byte(ehloCmd))
	if err != nil {
		return nil, tls.ConnectionState{}, fmt.Errorf("failed to send EHLO: %w", err)
	}

	// Read EHLO response
	n, err = conn.Read(buffer)
	if err != nil {
		return nil, tls.ConnectionState{}, fmt.Errorf("failed to read EHLO response: %w", err)
	}
	ehloResponse := string(buffer[:n])
	fmt.Printf("EHLO response: %s", utils.MaskIPAddresses(ehloResponse))

	// Check if STARTTLS is supported
	if !strings.Contains(ehloResponse, "STARTTLS") {
		return nil, tls.ConnectionState{}, fmt.Errorf("server does not support STARTTLS")
	}

	// Send STARTTLS command
	_, err = conn.Write([]byte("STARTTLS\r\n"))
	if err != nil {
		return nil, tls.ConnectionState{}, fmt.Errorf("failed to send STARTTLS: %w", err)
	}

	// Read STARTTLS response
	n, err = conn.Read(buffer)
	if err != nil {
		return nil, tls.ConnectionState{}, fmt.Errorf("failed to read STARTTLS response: %w", err)
	}
	startTLSResponse := string(buffer[:n])
	fmt.Printf("STARTTLS response: %s", utils.MaskIPAddresses(startTLSResponse))

	// Check for successful STARTTLS response (220)
	if !strings.HasPrefix(startTLSResponse, "220") {
		return nil, tls.ConnectionState{}, fmt.Errorf("STARTTLS command failed: %s", startTLSResponse)
	}

	// Now upgrade to TLS
	tlsConfig := &tls.Config{
		ServerName:         config.Host,
		InsecureSkipVerify: config.SkipVerify,
	}

	tlsConn := tls.Client(conn, tlsConfig)

	// Perform TLS handshake
	if err := tlsConn.Handshake(); err != nil {
		return nil, tls.ConnectionState{}, fmt.Errorf("TLS handshake failed: %w", err)
	}

	fmt.Println("StartTLS upgrade successful!")

	// Get certificate information
	state := tlsConn.ConnectionState()

	return state.PeerCertificates, state, nil
}