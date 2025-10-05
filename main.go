package main

/*
** With assistance from Claude AI Sonnet 4,5
** muquit@muquit.com Sep-27-2025 - first cut
** muquit@muquit.com Oct-03-2025 - v1.0.2
** muquit@muquit.com Oct-04-2025 - v1.0.3 - refactor code
*/

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/muquit/ssl-tls-verify/pkg/certificate"
	"github.com/muquit/ssl-tls-verify/pkg/connection"
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
		fmt.Fprintf(out, "  text  - Human-readable text format with details\n")
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
		handleCertificateFile(config)
		return
	}

	// Otherwise, test SSL/TLS connection
	handleConnection(config)
}

func handleCertificateFile(config Config) {
	fmt.Printf("Reading certificate file: %s\n", config.CertFile)

	certs, format, err := certificate.ParseFile(config.CertFile)
	if err != nil {
		log.Fatalf("Failed to display certificate file: %v", err)
	}

	fmt.Printf("Format: %s\n", format)
	certificate.Display(certs)
}

func handleConnection(config Config) {
	printConnectionInfo(config)

	if config.UseStartTLS {
		certs, state, err := connection.TestStartTLS(
			connection.StartTLSConfig{
				Host:           config.Host,
				Port:           config.Port,
				SkipVerify:     config.SkipVerify,
				ConnectTimeout: config.ConnectTimeout,
			},
		)
		if err != nil {
			log.Fatalf("StartTLS connection failed: %v", err)
		}

		// If save format is specified, save certificates instead of displaying
		if config.SaveFormat != "" {
			err = certificate.Save(certs, certificate.SaveConfig{
				Format:     config.SaveFormat,
				OutputFile: config.OutputFile,
			})
			if err != nil {
				log.Fatalf("Failed to save certificates: %v", err)
			}
		} else {
			connection.PrintTLSInfo(state)
			fmt.Println("StartTLS connection successful!")
		}
	} else {
		certs, state, err := connection.TestDirectTLS(
			connection.TLSConfig{
				Host:           config.Host,
				Port:           config.Port,
				SkipVerify:     config.SkipVerify,
				ConnectTimeout: config.ConnectTimeout,
			},
		)
		if err != nil {
			log.Fatalf("Direct TLS connection failed: %v", err)
		}

		// If save format is specified, save certificates instead of displaying
		if config.SaveFormat != "" {
			err = certificate.Save(certs, certificate.SaveConfig{
				Format:     config.SaveFormat,
				OutputFile: config.OutputFile,
			})
			if err != nil {
				log.Fatalf("Failed to save certificates: %v", err)
			}
		} else {
			connection.PrintTLSInfo(state)
			fmt.Println("Direct TLS connection successful!")
		}
	}
}

func printConnectionInfo(config Config) {
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
}
