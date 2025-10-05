package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
	"strings"
)

// SaveConfig holds certificate export configuration.
type SaveConfig struct {
	Format     string // "pem", "der", or "text"
	OutputFile string
}

// Save saves certificates in the specified format.
// Writes to file if OutputFile is specified, otherwise to stdout.
func Save(certs []*x509.Certificate, config SaveConfig) error {
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

	switch config.Format {
	case "pem":
		return SavePEM(certs, output)
	case "der":
		return SaveDER(certs, output)
	case "text":
		return SaveText(certs, output)
	default:
		return fmt.Errorf("unsupported save format: %s", config.Format)
	}
}

// SavePEM saves certificates in PEM format.
func SavePEM(certs []*x509.Certificate, output io.Writer) error {
	outFile, isFile := output.(*os.File)

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

	if isFile && outFile != os.Stdout {
		fmt.Printf("\n✅ Successfully saved %d certificate(s) in PEM format to %s\n",
			len(certs), outFile.Name())
	}

	return nil
}

// SaveDER saves certificate in DER format (leaf certificate only).
// DER format only saves the first (leaf) certificate as multiple DER
// certificates cannot be concatenated in a single file.
func SaveDER(certs []*x509.Certificate, output io.Writer) error {
	if len(certs) == 0 {
		return fmt.Errorf("no certificates to save")
	}

	outFile, isFile := output.(*os.File)

	// DER format: save only the leaf certificate (first one)
	// Multiple DER certificates cannot be concatenated in a single file
	leafCert := certs[0]

	_, err := output.Write(leafCert.Raw)
	if err != nil {
		return fmt.Errorf("failed to write certificate: %w", err)
	}

	if isFile && outFile != os.Stdout {
		if len(certs) > 1 {
			fmt.Printf("\n✅ Successfully saved leaf certificate in DER format to %s\n", outFile.Name())
			fmt.Printf("   Note: Saved leaf certificate only. %d intermediate certificate(s) were not saved.\n", len(certs)-1)
			fmt.Printf("   Use PEM format (--save pem) to save the complete certificate chain in one file.\n")
		} else {
			fmt.Printf("\n✅ Successfully saved certificate in DER format to %s\n", outFile.Name())
		}
	}

	return nil
}

// SaveText saves certificates in human-readable text format.
func SaveText(certs []*x509.Certificate, output io.Writer) error {
	outFile, isFile := output.(*os.File)

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

		PrintDetailsToWriter(cert, output)
	}

	if isFile && outFile != os.Stdout {
		fmt.Printf("\n✅ Successfully saved %d certificate(s) in TEXT format to %s\n",
			len(certs), outFile.Name())
	}

	return nil
}