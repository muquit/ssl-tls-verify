package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"strings"
	"time"
)

type Config struct {
	Host           string
	Port           string
	SkipVerify     bool
	UseStartTLS    bool
	ConnectTimeout time.Duration
}

func main() {
	var config Config

	// Command line flags
	flag.StringVar(&config.Host, "host", "google.com", "Target host")
	flag.StringVar(&config.Port, "port", "443", "Target port (443 for HTTPS, 587 for SMTP StartTLS)")
	flag.BoolVar(&config.SkipVerify, "skip-verify", false, "Skip certificate verification")
	flag.BoolVar(&config.UseStartTLS, "starttls", false, "Use StartTLS (for SMTP-like protocols)")
	flag.DurationVar(&config.ConnectTimeout, "timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	fmt.Printf("Connecting to %s:%s\n", config.Host, config.Port)
	fmt.Printf("Skip certificate verification: %v\n", config.SkipVerify)
	fmt.Printf("Use StartTLS: %v\n", config.UseStartTLS)
	fmt.Println()

	if config.UseStartTLS {
		err := testStartTLS(config)
		if err != nil {
			log.Printf("StartTLS connection failed: %v", err)
		} else {
			fmt.Println("StartTLS connection successful!")
		}
	} else {
		err := testDirectTLS(config)
		if err != nil {
			log.Printf("Direct TLS connection failed: %v", err)
		} else {
			fmt.Println("Direct TLS connection successful!")
		}
	}
}

// testDirectTLS tests a direct TLS connection (like HTTPS)
func testDirectTLS(config Config) error {
	tlsConfig := &tls.Config{
		ServerName:         config.Host,
		InsecureSkipVerify: config.SkipVerify,
	}

	dialer := &net.Dialer{
		Timeout: config.ConnectTimeout,
	}

	conn, err := tls.DialWithDialer(dialer, "tcp", net.JoinHostPort(config.Host, config.Port), tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to establish TLS connection: %w", err)
	}
	defer conn.Close()

	// Get certificate information
	state := conn.ConnectionState()
	printTLSConnectionInfo(state)

	return nil
}

// printTLSConnectionInfo prints detailed TLS connection and certificate information
func printTLSConnectionInfo(state tls.ConnectionState) {
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
		
		printCertificateDetails(cert)
	}
}

// printCertificateDetails prints detailed certificate information in a pretty format
func printCertificateDetails(cert *x509.Certificate) {
	fmt.Printf("Subject:               %s\n", cert.Subject.String())
	fmt.Printf("Issuer:                %s\n", cert.Issuer.String())
	fmt.Printf("Serial Number:         %s\n", cert.SerialNumber.String())
	fmt.Printf("Version:               %d\n", cert.Version)
	
	// Validity period with time until expiry
	fmt.Printf("Valid From:            %s\n", cert.NotBefore.Format("2006-01-02 15:04:05 UTC"))
	fmt.Printf("Valid Until:           %s\n", cert.NotAfter.Format("2006-01-02 15:04:05 UTC"))
	
	// Check if certificate is expired or will expire soon
	now := time.Now()
	if now.After(cert.NotAfter) {
		fmt.Printf("Status:                ⚠️  EXPIRED\n")
	} else if now.Add(30*24*time.Hour).After(cert.NotAfter) {
		daysLeft := int(cert.NotAfter.Sub(now).Hours() / 24)
		fmt.Printf("Status:                ⚠️  EXPIRES IN %d DAYS\n", daysLeft)
	} else {
		daysLeft := int(cert.NotAfter.Sub(now).Hours() / 24)
		fmt.Printf("Status:                ✅ Valid (%d days remaining)\n", daysLeft)
	}
	
	// Subject Alternative Names
	if len(cert.DNSNames) > 0 {
		fmt.Printf("DNS Names:             %s\n", strings.Join(cert.DNSNames, ", "))
	}
	if len(cert.IPAddresses) > 0 {
		ips := make([]string, len(cert.IPAddresses))
		for i, ip := range cert.IPAddresses {
			ips[i] = ip.String()
		}
		fmt.Printf("IP Addresses:          %s\n", strings.Join(ips, ", "))
	}
	if len(cert.EmailAddresses) > 0 {
		fmt.Printf("Email Addresses:       %s\n", strings.Join(cert.EmailAddresses, ", "))
	}
	
	// Key information
	fmt.Printf("Public Key Algorithm:  %s\n", cert.PublicKeyAlgorithm.String())
	fmt.Printf("Signature Algorithm:   %s\n", cert.SignatureAlgorithm.String())
	
	// Key usage
	if cert.KeyUsage != 0 {
		var usages []string
		if cert.KeyUsage&x509.KeyUsageDigitalSignature != 0 {
			usages = append(usages, "Digital Signature")
		}
		if cert.KeyUsage&x509.KeyUsageKeyEncipherment != 0 {
			usages = append(usages, "Key Encipherment")
		}
		if cert.KeyUsage&x509.KeyUsageKeyAgreement != 0 {
			usages = append(usages, "Key Agreement")
		}
		if cert.KeyUsage&x509.KeyUsageCertSign != 0 {
			usages = append(usages, "Certificate Signing")
		}
		if cert.KeyUsage&x509.KeyUsageCRLSign != 0 {
			usages = append(usages, "CRL Signing")
		}
		fmt.Printf("Key Usage:             %s\n", strings.Join(usages, ", "))
	}
	
	// Extended key usage
	if len(cert.ExtKeyUsage) > 0 {
		var extUsages []string
		for _, usage := range cert.ExtKeyUsage {
			switch usage {
			case x509.ExtKeyUsageServerAuth:
				extUsages = append(extUsages, "Server Authentication")
			case x509.ExtKeyUsageClientAuth:
				extUsages = append(extUsages, "Client Authentication")
			case x509.ExtKeyUsageEmailProtection:
				extUsages = append(extUsages, "Email Protection")
			case x509.ExtKeyUsageTimeStamping:
				extUsages = append(extUsages, "Time Stamping")
			case x509.ExtKeyUsageCodeSigning:
				extUsages = append(extUsages, "Code Signing")
			default:
				extUsages = append(extUsages, fmt.Sprintf("Unknown (%v)", usage))
			}
		}
		fmt.Printf("Extended Key Usage:    %s\n", strings.Join(extUsages, ", "))
	}
	
	// Fingerprints
	fmt.Printf("SHA-1 Fingerprint:     %s\n", formatFingerprint(cert.Raw, "sha1"))
	fmt.Printf("SHA-256 Fingerprint:   %s\n", formatFingerprint(cert.Raw, "sha256"))
	
	// Certificate Authority
	if cert.IsCA {
		fmt.Printf("Certificate Authority: Yes\n")
		if cert.MaxPathLen >= 0 {
			fmt.Printf("Max Path Length:       %d\n", cert.MaxPathLen)
		} else if cert.MaxPathLenZero {
			fmt.Printf("Max Path Length:       0\n")
		}
	} else {
		fmt.Printf("Certificate Authority: No\n")
	}
}

// formatFingerprint creates a formatted fingerprint string
func formatFingerprint(certRaw []byte, algorithm string) string {
	var hash []byte
	switch algorithm {
	case "sha1":
		h := sha1.Sum(certRaw)
		hash = h[:]
	case "sha256":
		h := sha256.Sum256(certRaw)
		hash = h[:]
	default:
		return "Unknown algorithm"
	}
	
	hexStr := hex.EncodeToString(hash)
	// Format as XX:XX:XX...
	var formatted []string
	for i := 0; i < len(hexStr); i += 2 {
		formatted = append(formatted, strings.ToUpper(hexStr[i:i+2]))
	}
	return strings.Join(formatted, ":")
}

// testStartTLS tests a StartTLS connection (like SMTP)
func testStartTLS(config Config) error {
	// First, establish a plain connection
	dialer := &net.Dialer{
		Timeout: config.ConnectTimeout,
	}

	conn, err := dialer.Dial("tcp", net.JoinHostPort(config.Host, config.Port))
	if err != nil {
		return fmt.Errorf("failed to establish plain connection: %w", err)
	}
	defer conn.Close()

	fmt.Println("Plain connection established, attempting StartTLS...")

	// For demonstration, we'll use SMTP as an example of StartTLS
	// This assumes the server supports SMTP protocol
	client, err := smtp.NewClient(conn, config.Host)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Quit()

	// Check if the server supports StartTLS
	if ok, _ := client.Extension("STARTTLS"); !ok {
		return fmt.Errorf("server does not support StartTLS")
	}

	// Configure TLS
	tlsConfig := &tls.Config{
		ServerName:         config.Host,
		InsecureSkipVerify: config.SkipVerify,
	}

	// Upgrade to TLS
	if err = client.StartTLS(tlsConfig); err != nil {
		return fmt.Errorf("StartTLS failed: %w", err)
	}

	fmt.Println("StartTLS upgrade successful!")

	fmt.Println("TLS connection is now active")
	// Note: smtp.Client doesn't expose the underlying TLS connection state directly
	// For detailed TLS info, see the testGenericStartTLS function below

	return nil
}

// Alternative generic StartTLS implementation
func testGenericStartTLS(config Config) error {
	// Establish plain connection
	dialer := &net.Dialer{
		Timeout: config.ConnectTimeout,
	}

	conn, err := dialer.Dial("tcp", net.JoinHostPort(config.Host, config.Port))
	if err != nil {
		return fmt.Errorf("failed to establish connection: %w", err)
	}
	defer conn.Close()

	fmt.Println("Plain connection established")

	// Configure TLS
	tlsConfig := &tls.Config{
		ServerName:         config.Host,
		InsecureSkipVerify: config.SkipVerify,
	}

	// Upgrade to TLS (generic approach)
	tlsConn := tls.Client(conn, tlsConfig)
	
	// Perform TLS handshake
	if err := tlsConn.Handshake(); err != nil {
		return fmt.Errorf("TLS handshake failed: %w", err)
	}

	fmt.Println("TLS handshake successful!")

	// Get certificate information
	state := tlsConn.ConnectionState()
	printTLSConnectionInfo(state)

	return nil
}