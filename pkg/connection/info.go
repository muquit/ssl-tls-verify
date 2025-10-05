package connection

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/muquit/ssl-tls-verify/pkg/certificate"
)

// PrintTLSInfo prints detailed TLS connection and certificate information.
func PrintTLSInfo(state tls.ConnectionState) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("TLS CONNECTION INFORMATION")
	fmt.Println(strings.Repeat("=", 60))

	// TLS Version mapping
	versionMap := map[uint16]string{
		tls.VersionSSL30: "SSL 3.0",
		tls.VersionTLS10: "TLS 1.0",
		tls.VersionTLS11: "TLS 1.1",
		tls.VersionTLS12: "TLS 1.2",
		tls.VersionTLS13: "TLS 1.3",
	}

	version := fmt.Sprintf("0x%04x (Unknown)", state.Version)
	if v, ok := versionMap[state.Version]; ok {
		version = fmt.Sprintf("%s (0x%04x)", v, state.Version)
	}

	fmt.Printf("TLS Version:           %s\n", version)
	fmt.Printf("Cipher Suite:          0x%04x\n", state.CipherSuite)
	fmt.Printf("Server Certificates:   %d\n", len(state.PeerCertificates))
	fmt.Printf("Handshake Complete:    %v\n", state.HandshakeComplete)
	fmt.Printf("Did Resume:            %v\n", state.DidResume)
	fmt.Printf("Negotiated Protocol:   %s\n", state.NegotiatedProtocol)

	if len(state.PeerCertificates) == 0 {
		fmt.Println("\nNo certificates found")
		return
	}

	// Print certificate chain
	for i, cert := range state.PeerCertificates {
		fmt.Printf("\n%s\n", strings.Repeat("-", 60))
		if i == 0 {
			fmt.Printf("SERVER CERTIFICATE #%d (Leaf Certificate)\n", i+1)
		} else {
			fmt.Printf("INTERMEDIATE CERTIFICATE #%d\n", i+1)
		}
		fmt.Printf("%s\n", strings.Repeat("-", 60))

		certificate.PrintDetails(cert)
	}
}