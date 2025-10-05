package certificate

import (
	"crypto/x509"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/muquit/ssl-tls-verify/pkg/utils"
)

// Display displays certificates in human-readable format to stdout.
func Display(certs []*x509.Certificate) {
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

		PrintDetails(cert)
	}
}

// PrintDetails prints detailed certificate information to stdout.
func PrintDetails(cert *x509.Certificate) {
	PrintDetailsToWriter(cert, os.Stdout)
}

// PrintDetailsToWriter prints certificate details to a specific writer.
func PrintDetailsToWriter(cert *x509.Certificate, output io.Writer) {
	fmt.Fprintf(output, "Subject:               %s\n", utils.MaskIPAddresses(cert.Subject.String()))
	fmt.Fprintf(output, "Issuer:                %s\n", utils.MaskIPAddresses(cert.Issuer.String()))
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
		maskedDNS := utils.MaskIPSlice(cert.DNSNames)
		fmt.Fprintf(output, "DNS Names:             %s\n", strings.Join(maskedDNS, ", "))
	}
	if len(cert.IPAddresses) > 0 {
		ips := make([]string, len(cert.IPAddresses))
		for i, ip := range cert.IPAddresses {
			ips[i] = ip.String()
		}
		maskedIPs := utils.MaskIPSlice(ips)
		fmt.Fprintf(output, "IP Addresses:          %s\n", strings.Join(maskedIPs, ", "))
	}
	if len(cert.EmailAddresses) > 0 {
		maskedEmails := utils.MaskIPSlice(cert.EmailAddresses)
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
	fmt.Fprintf(output, "SHA-1 Fingerprint:     %s\n", FormatFingerprint(cert.Raw, "sha1"))
	fmt.Fprintf(output, "SHA-256 Fingerprint:   %s\n", FormatFingerprint(cert.Raw, "sha256"))

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