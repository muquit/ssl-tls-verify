package main

/*
** With assistance from Claude AI Sonnet 4,5
** muquit@muquit.com Sep-27-2025 - first cut
** muquit@muquit.com Oct-03-2025 - v1.0.2
*/

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/muquit/ssl-tls-verify/pkg/version"
)

type Config struct {
	Host           string
	Port           string
	SkipVerify     bool
	UseStartTLS    bool
	ConnectTimeout time.Duration
	CertFile       string
	SaveFormat     string
	OutputFile     string
}

const (
	me  = "ssl-tls-verify"
	url = "https://github.com/muquit/ssl-tls-verify"
)

func main() {
	var config Config
	var showVersion bool

	flag.StringVar(&config.Host, "host", "google.com", "Target host")
	flag.StringVar(&config.Port, "port", "443", "Target port (443 for HTTPS, 587 for SMTP StartTLS)")
	flag.BoolVar(&config.SkipVerify, "skip-verify", false, "Skip certificate verification")
	flag.BoolVar(&config.UseStartTLS, "starttls", false, "Use StartTLS (for SMTP-like protocols)")
	flag.DurationVar(&config.ConnectTimeout, "timeout", 10*time.Second, "Connection timeout")
	flag.StringVar(&config.CertFile, "dump", "", "Read and display certificate file (PEM or DER format)")
	flag.StringVar(&config.SaveFormat, "save", "", "Save certificates from connection (formats: pem, der, text)")
	flag.StringVar(&config.OutputFile, "output", "", "Output file for saved certificates (default: stdout)")
	flag.BoolVar(&showVersion, "version", false, "Show version information and exit")

	flag.Usage = func() {
		out := os.Stdout
		fmt.Fprintf(out, "%s %s\n%s\n", me, version.Get(), url)
		fmt.Fprintf(out, "A cross-platform command line tool to test and verify SSL/TLS connections\n")
		fmt.Fprintf(out, "and display certificate files in a human readable format\n\n")
		fmt.Fprintf(out, "Usage:\n")
		flag.PrintDefaults()
		fmt.Fprintf(out, "\nSave Formats:\n")
		fmt.Fprintf(out, "  pem   - PEM encoded certificates (Base64 with headers)\n")
		fmt.Fprintf(out, "  der   - DER encoded certificates (binary format)\n")
		fmt.Fprintf(out, "  text  - Human-readable text format with full details\n")
		fmt.Fprintf(out, "\nExamples:\n")
		fmt.Fprintf(out, "  Test SSL/TLS connection:\n")
		fmt.Fprintf(out, "    %s --host google.com\n", me)
		fmt.Fprintf(out, "    %s --host smtp.gmail.com --port 587 --starttls\n", me)
		fmt.Fprintf(out, "\n  Display certificate files:\n")
		fmt.Fprintf(out, "    %s --dump certificate.pem\n", me)
		fmt.Fprintf(out, "    %s --dump certificate.der\n", me)
		fmt.Fprintf(out, "\n  Save certificates from connection:\n")
		fmt.Fprintf(out, "    %s --host google.com --save pem --output certs.pem\n", me)
		fmt.Fprintf(out, "    %s --host google.com --save text --output report.txt\n", me)
		fmt.Fprintf(out, "    %s --host smtp.gmail.com --port 587 --starttls --save pem\n", me)
	}

	flag.Parse()

	if showVersion {
		fmt.Printf("%s %s\n%s\n", me, version.Get(), url)
		os.Exit(0)
	}

	// Validate save format if specified
	if config.SaveFormat != "" {
		config.SaveFormat = strings.ToLower(config.SaveFormat)
		if config.SaveFormat != "pem" && config.SaveFormat != "der" && config.SaveFormat != "text" {
			log.Fatalf("Invalid save format: %s (valid formats: pem, der, text)", config.SaveFormat)
		}
		// DER format requires output file (binary data shouldn't go to stdout)
		if config.SaveFormat == "der" && config.OutputFile == "" {
			log.Fatalf("DER format requires --output flag (binary data cannot be written to stdout)")
		}
	}

	// If --dump is specified, read and display the certificate file
	if config.CertFile != "" {
		if err := displayCertificateFile(config.CertFile); err != nil {
			log.Fatalf("Failed to display certificate file: %v", err)
		}
		return
	}

	// Otherwise, test SSL/TLS connection
	fmt.Printf("Connecting to %s:%s\n", config.Host, config.Port)
	if config.SkipVerify {
		fmt.Println("Verify Certificate: No")
	} else {
		fmt.Println("Verify Certificate: Yes")
	}

	if config.UseStartTLS {
		fmt.Println("Use StartTLS: Yes")
	} else {
		fmt.Println("Use StartTLS: No")
	}

	if config.SaveFormat != "" {
		fmt.Printf("Save Format: %s\n", strings.ToUpper(config.SaveFormat))
		if config.OutputFile != "" {
			fmt.Printf("Output File: %s\n", config.OutputFile)
		} else {
			fmt.Println("Output: stdout")
		}
	}

	fmt.Println()

	if config.UseStartTLS {
		err := testStartTLS(config)
		if err != nil {
			log.Fatalf("StartTLS connection failed: %v", err)
		} else if config.SaveFormat == "" {
			fmt.Println("StartTLS connection successful!")
		}
	} else {
		err := testDirectTLS(config)
		if err != nil {
			log.Fatalf("Direct TLS connection failed: %v", err)
		} else if config.SaveFormat == "" {
			fmt.Println("Direct TLS connection successful!")
		}
	}
}

// displayCertificateFile reads a certificate file and displays it in human
// readable format
func displayCertificateFile(filename string) error {
	fmt.Printf("Reading certificate file: %s\n", filename)

	// Read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Try to parse as PEM first
	certs, err := parsePEMCertificates(data)
	if err == nil && len(certs) > 0 {
		fmt.Printf("Format: PEM\n")
		displayCertificates(certs)
		return nil
	}

	// Try to parse as DER
	cert, err := x509.ParseCertificate(data)
	if err == nil {
		fmt.Printf("Format: DER\n")
		displayCertificates([]*x509.Certificate{cert})
		return nil
	}

	return fmt.Errorf("unable to parse certificate file (tried PEM and DER formats)")
}

// parsePEMCertificates parses PEM-encoded certificates
func parsePEMCertificates(data []byte) ([]*x509.Certificate, error) {
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

// displayCertificates displays certificates in 
// human readable format
func displayCertificates(certs []*x509.Certificate) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("CERTIFICATE DETAILS")
	fmt.Printf("Total Certificates: %d\n", len(certs))
	fmt.Println(strings.Repeat("=", 60))

	for i, cert := range certs {
		fmt.Printf("\n%s\n", strings.Repeat("-", 60))
		if len(certs) == 1 {
			fmt.Printf("CERTIFICATE\n")
		} else if i == 0 {
			fmt.Printf("CERTIFICATE #%d (Leaf Certificate)\n", i+1)
		} else {
			fmt.Printf("CERTIFICATE #%d (Intermediate Certificate)\n", i+1)
		}
		fmt.Printf("%s\n", strings.Repeat("-", 60))

		printCertificateDetails(cert)
	}
}

// saveCertificates saves certificates in the specified format
func saveCertificates(certs []*x509.Certificate, config Config) error {
	var output *os.File
	var err error

	if config.OutputFile != "" {
		output, err = os.Create(config.OutputFile)
		if err != nil {
			return fmt.Errorf("failed to create output file: %w", err)
		}
		defer output.Close()
	} else {
		output = os.Stdout
	}

	switch config.SaveFormat {
	case "pem":
		return savePEM(certs, output)
	case "der":
		return saveDER(certs, output)
	case "text":
		return saveText(certs, output)
	default:
		return fmt.Errorf("unsupported save format: %s", config.SaveFormat)
	}
}

// savePEM saves certificates in PEM format
func savePEM(certs []*x509.Certificate, output *os.File) error {
	for i, cert := range certs {
		if i > 0 {
			fmt.Fprintln(output) // Add blank line between certificates
		}

		certType := "CERTIFICATE"
		if i == 0 {
			fmt.Fprintf(output, "# Certificate #%d - Leaf Certificate\n", i+1)
		} else {
			fmt.Fprintf(output, "# Certificate #%d - Intermediate Certificate\n", i+1)
		}
		fmt.Fprintf(output, "# Subject: %s\n", cert.Subject.String())
		fmt.Fprintf(output, "# Issuer: %s\n", cert.Issuer.String())
		fmt.Fprintln(output)

		block := &pem.Block{
			Type:  certType,
			Bytes: cert.Raw,
		}

		if err := pem.Encode(output, block); err != nil {
			return fmt.Errorf("failed to encode certificate %d: %w", i+1, err)
		}
	}

	if output != os.Stdout {
		fmt.Printf("\n✅ Successfully saved %d certificate(s) in PEM format to %s\n",
			len(certs), output.Name())
	}

	return nil
}

// saveDER saves certificate in DER format (leaf certificate only)
func saveDER(certs []*x509.Certificate, output *os.File) error {
	if len(certs) == 0 {
		return fmt.Errorf("no certificates to save")
	}

	// DER format: save only the leaf certificate (first one)
	// Multiple DER certificates cannot be concatenated in a single file
	leafCert := certs[0]
	
	_, err := output.Write(leafCert.Raw)
	if err != nil {
		return fmt.Errorf("failed to write certificate: %w", err)
	}

	if output != os.Stdout {
		if len(certs) > 1 {
			fmt.Printf("\n✅ Successfully saved leaf certificate in DER format to %s\n", output.Name())
			fmt.Printf("   Note: Saved leaf certificate only. %d intermediate certificate(s) were not saved.\n", len(certs)-1)
			fmt.Printf("   Use PEM format (--save pem) to save the complete certificate chain in one file.\n")
		} else {
			fmt.Printf("\n✅ Successfully saved certificate in DER format to %s\n", output.Name())
		}
	}

	return nil
}

// saveText saves certificates in human-readable text format
func saveText(certs []*x509.Certificate, output *os.File) error {
	fmt.Fprintln(output, strings.Repeat("=", 60))
	fmt.Fprintln(output, "CERTIFICATE CHAIN DETAILS")
	fmt.Fprintf(output, "Total Certificates: %d\n", len(certs))
	fmt.Fprintln(output, strings.Repeat("=", 60))

	for i, cert := range certs {
		fmt.Fprintf(output, "\n%s\n", strings.Repeat("-", 60))
		if i == 0 {
			fmt.Fprintf(output, "CERTIFICATE #%d (Leaf Certificate)\n", i+1)
		} else {
			fmt.Fprintf(output, "CERTIFICATE #%d (Intermediate Certificate)\n", i+1)
		}
		fmt.Fprintf(output, "%s\n", strings.Repeat("-", 60))

		printCertificateDetailsToWriter(cert, output)
	}

	if output != os.Stdout {
		fmt.Printf("\n✅ Successfully saved %d certificate(s) in TEXT format to %s\n",
			len(certs), output.Name())
	}

	return nil
}

// printCertificateDetailsToWriter prints certificate details to a specific writer
func printCertificateDetailsToWriter(cert *x509.Certificate, output *os.File) {
	fmt.Fprintf(output, "Subject:               %s\n", maskIPAddresses(cert.Subject.String()))
	fmt.Fprintf(output, "Issuer:                %s\n", maskIPAddresses(cert.Issuer.String()))
	fmt.Fprintf(output, "Serial Number:         %s\n", cert.SerialNumber.String())
	fmt.Fprintf(output, "Version:               %d\n", cert.Version)

	// Validity period with time until expiry
	fmt.Fprintf(output, "Valid From:            %s\n", cert.NotBefore.Format("2006-01-02 15:04:05 UTC"))
	fmt.Fprintf(output, "Valid Until:           %s\n", cert.NotAfter.Format("2006-01-02 15:04:05 UTC"))

	// Check if certificate is expired or will expire soon
	now := time.Now()
	if now.After(cert.NotAfter) {
		fmt.Fprintf(output, "Status:                ⚠️  EXPIRED\n")
	} else if now.Add(30*24*time.Hour).After(cert.NotAfter) {
		daysLeft := int(cert.NotAfter.Sub(now).Hours() / 24)
		fmt.Fprintf(output, "Status:                ⚠️  EXPIRES IN %d DAYS\n", daysLeft)
	} else {
		daysLeft := int(cert.NotAfter.Sub(now).Hours() / 24)
		fmt.Fprintf(output, "Status:                ✅ Valid (%d days remaining)\n", daysLeft)
	}

	// Subject Alternative Names with IP masking
	if len(cert.DNSNames) > 0 {
		maskedDNS := maskIPSlice(cert.DNSNames)
		fmt.Fprintf(output, "DNS Names:             %s\n", strings.Join(maskedDNS, ", "))
	}
	if len(cert.IPAddresses) > 0 {
		ips := make([]string, len(cert.IPAddresses))
		for i, ip := range cert.IPAddresses {
			ips[i] = ip.String()
		}
		maskedIPs := maskIPSlice(ips)
		fmt.Fprintf(output, "IP Addresses:          %s\n", strings.Join(maskedIPs, ", "))
	}
	if len(cert.EmailAddresses) > 0 {
		maskedEmails := maskIPSlice(cert.EmailAddresses)
		fmt.Fprintf(output, "Email Addresses:       %s\n", strings.Join(maskedEmails, ", "))
	}

	// Key information
	fmt.Fprintf(output, "Public Key Algorithm:  %s\n", cert.PublicKeyAlgorithm.String())
	fmt.Fprintf(output, "Signature Algorithm:   %s\n", cert.SignatureAlgorithm.String())

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
		fmt.Fprintf(output, "Key Usage:             %s\n", strings.Join(usages, ", "))
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
		fmt.Fprintf(output, "Extended Key Usage:    %s\n", strings.Join(extUsages, ", "))
	}

	// Fingerprints
	fmt.Fprintf(output, "SHA-1 Fingerprint:     %s\n", formatFingerprint(cert.Raw, "sha1"))
	fmt.Fprintf(output, "SHA-256 Fingerprint:   %s\n", formatFingerprint(cert.Raw, "sha256"))

	// Certificate Authority
	if cert.IsCA {
		fmt.Fprintf(output, "Certificate Authority: Yes\n")
		if cert.MaxPathLen >= 0 {
			fmt.Fprintf(output, "Max Path Length:       %d\n", cert.MaxPathLen)
		} else if cert.MaxPathLenZero {
			fmt.Fprintf(output, "Max Path Length:       0\n")
		}
	} else {
		fmt.Fprintf(output, "Certificate Authority: No\n")
	}
}

// maskIPAddresses masks IPv4 and IPv6 addresses in text
func maskIPAddresses(text string) string {
	if os.Getenv("MAILSEND_MASK_IP") == "" {
		return text
	}

	// IPv4 pattern: matches xxx.xxx.xxx.xxx format
	ipv4Pattern := regexp.MustCompile(`\b(\d{1,3}\.){3}\d{1,3}\b`)

	// IPv6 pattern: matches various IPv6 formats
	ipv6Pattern := regexp.MustCompile(`\b([0-9a-fA-F]{0,4}:){2,7}[0-9a-fA-F]{0,4}\b`)

	// Replace IPv4 addresses
	maskedText := ipv4Pattern.ReplaceAllString(text, "xxx.xxx.xxx.xxx")

	// Replace IPv6 addresses
	maskedText = ipv6Pattern.ReplaceAllString(maskedText, "xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx")

	return maskedText
}

// maskIPSlice masks IP addresses in a string slice
func maskIPSlice(addresses []string) []string {
	if os.Getenv("MAILSEND_MASK_IP") == "" {
		return addresses
	}

	masked := make([]string, len(addresses))
	for i, addr := range addresses {
		masked[i] = maskIPAddresses(addr)
	}
	return masked
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

	// If save format is specified, save certificates instead of displaying
	if config.SaveFormat != "" {
		return saveCertificates(state.PeerCertificates, config)
	}

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
	printCertificateDetailsToWriter(cert, os.Stdout)
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

	// For SMTP, we need to do the protocol handshake manually to maintain control
	// over the connection for certificate inspection

	// Read server greeting
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return fmt.Errorf("failed to read server greeting: %w", err)
	}
	greeting := string(buffer[:n])
	fmt.Printf("Server greeting: %s", maskIPAddresses(greeting))

	// Send EHLO command
	ehloCmd := fmt.Sprintf("EHLO %s\r\n", config.Host)
	_, err = conn.Write([]byte(ehloCmd))
	if err != nil {
		return fmt.Errorf("failed to send EHLO: %w", err)
	}

	// Read EHLO response
	n, err = conn.Read(buffer)
	if err != nil {
		return fmt.Errorf("failed to read EHLO response: %w", err)
	}
	ehloResponse := string(buffer[:n])
	fmt.Printf("EHLO response: %s", maskIPAddresses(ehloResponse))

	// Check if STARTTLS is supported
	if !strings.Contains(ehloResponse, "STARTTLS") {
		return fmt.Errorf("server does not support STARTTLS")
	}

	// Send STARTTLS command
	_, err = conn.Write([]byte("STARTTLS\r\n"))
	if err != nil {
		return fmt.Errorf("failed to send STARTTLS: %w", err)
	}

	// Read STARTTLS response
	n, err = conn.Read(buffer)
	if err != nil {
		return fmt.Errorf("failed to read STARTTLS response: %w", err)
	}
	startTLSResponse := string(buffer[:n])
	fmt.Printf("STARTTLS response: %s", maskIPAddresses(startTLSResponse))

	// Check for successful STARTTLS response (220)
	if !strings.HasPrefix(startTLSResponse, "220") {
		return fmt.Errorf("STARTTLS command failed: %s", startTLSResponse)
	}

	// Now upgrade to TLS
	tlsConfig := &tls.Config{
		ServerName:         config.Host,
		InsecureSkipVerify: config.SkipVerify,
	}

	tlsConn := tls.Client(conn, tlsConfig)

	// Perform TLS handshake
	if err := tlsConn.Handshake(); err != nil {
		return fmt.Errorf("TLS handshake failed: %w", err)
	}

	fmt.Println("StartTLS upgrade successful!")

	// Get certificate information
	state := tlsConn.ConnectionState()

	// If save format is specified, save certificates instead of displaying
	if config.SaveFormat != "" {
		return saveCertificates(state.PeerCertificates, config)
	}

	printTLSConnectionInfo(state)

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
