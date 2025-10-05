package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// ParsePEMCertificates parses PEM-encoded certificates from raw bytes.
// Returns a slice of parsed certificates or an error if parsing fails.
func ParsePEMCertificates(data []byte) ([]*x509.Certificate, error) {
	var certs []*x509.Certificate

	for {
		block, rest := pem.Decode(data)
		if block == nil {
			break
		}

		if block.Type == "CERTIFICATE" {
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("failed to parse certificate: %w", err)
			}
			certs = append(certs, cert)
		}

		data = rest
	}

	if len(certs) == 0 {
		return nil, fmt.Errorf("no certificates found in PEM data")
	}

	return certs, nil
}

// ParseFile reads and parses a certificate file in either PEM or DER format.
// Returns certificates, format string ("PEM" or "DER"), and error.
func ParseFile(filename string) ([]*x509.Certificate, string, error) {
	// Read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read file: %w", err)
	}

	// Try to parse as PEM first
	certs, err := ParsePEMCertificates(data)
	if err == nil && len(certs) > 0 {
		return certs, "PEM", nil
	}

	// Try to parse as DER
	cert, err := x509.ParseCertificate(data)
	if err == nil {
		return []*x509.Certificate{cert}, "DER", nil
	}

	return nil, "", fmt.Errorf("unable to parse certificate file (tried PEM and DER formats)")
}