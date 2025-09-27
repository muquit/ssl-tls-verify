# ssl-tls-verify

## Introduction
`ssl-tls-verify` is a command line tool to test and verify SSL/TLS connections with detailed certificate inspection. This is a
[golang](https://golang.org/) utility that supports both direct TLS connections (like HTTPS) and StartTLS protocol upgrades (like SMTP). 
The tool provides comprehensive certificate chain analysis with pretty-printed output, making it ideal for debugging SSL/TLS issues, 
certificate validation problems, and connection troubleshooting.

Unlike basic tools like `openssl s_client`, this utility focuses on providing human-readable certificate information with 
expiry warnings, fingerprint validation, and detailed TLS connection state analysis. All the certificate parsing and validation 
is handled by golang's robust `crypto/tls` and `crypto/x509` packages.

If you have any question, request or suggestion, please enter it in the
[Issues](https://github.com/your-username/ssl-tls-verify/issues) with appropriate label.

## Background

There was a lingering bug in @MAILSEND_GO@ that `-verifyCert` does not work.
I was not getting time to look at the code and figure out the issue, so I asked @CLAUDE_SONNET@ to
write a simple CLI tool to verify certificates with optional verification. Hence, ssl-tls-verify was born.
After creating this tool, it became clear to me that I had missed ServerName in tls.Config -
a critical field needed for proper certificate hostname validation.
I also love this tool, it's much nicer than openssl with its human-readable output,
clear certificate chain display, and intuitive formatting. What started as a debugging
utility has become my go-to tool for SSL/TLS certificate inspection. Hope you
find that tool useful as well.

## Features

- **Direct TLS connections** - Test HTTPS and other TLS-enabled services
- **StartTLS support** - Test SMTP and other services that upgrade from plain to TLS
- **Certificate chain inspection** - Display complete certificate chains with detailed information
- **Pretty-printed output** - Human-readable certificate details with visual formatting
- **Expiry warnings** - Smart detection of expired or soon-to-expire certificates
- **Fingerprint validation** - SHA-1 and SHA-256 certificate fingerprints
- **Flexible verification** - Option to skip certificate verification for testing

## Installation

```bash
go build -o ssl-tls-verify main.go
```

Or install directly:
```bash
go install github.com/your-username/ssl-tls-verify@latest
```

## Usage

### Basic Examples

Test HTTPS connection with certificate verification:
```bash
./ssl-tls-verify --host google.com --port 443
```

Test HTTPS connection ignoring certificate errors:
```bash
./ssl-tls-verify --host self-signed-site.com --port 443 --skip-verify
```

Test SMTP with StartTLS:
```bash
./ssl-tls-verify --host smtp.gmail.com --port 587 --starttls
```

Test with custom timeout:
```bash
./ssl-tls-verify --host example.com --port 443 --timeout 5s
```

### Command Line Options

```
  -host string
        Target host (default "google.com")
  -port string
        Target port (default "443")
  -skip-verify
        Skip certificate verification
  -starttls
        Use StartTLS (for SMTP-like protocols)
  -timeout duration
        Connection timeout (default 10s)
```

## Output

The tool provides detailed information about:

- **TLS Connection Details**
  - TLS version (1.0, 1.1, 1.2, 1.3)
  - Cipher suite information
  - Handshake status
  - Protocol negotiation results

- **Certificate Information**
  - Subject and Issuer details
  - Validity period with expiry warnings
  - Subject Alternative Names (DNS, IP, Email)
  - Public key and signature algorithms
  - Key usage and extended key usage
  - Certificate Authority status
  - SHA-1 and SHA-256 fingerprints

- **Certificate Chain**
  - Complete certificate chain display
  - Leaf and intermediate certificate identification
  - Individual certificate analysis

## Use Cases

- **SSL/TLS Troubleshooting** - Debug connection issues and certificate problems
- **Certificate Monitoring** - Check certificate expiry dates and validity
- **Security Auditing** - Inspect certificate chains and TLS configurations
- **Development Testing** - Verify SSL/TLS setups in development environments
- **Automation Scripts** - Integrate into monitoring and alerting systems

## Requirements

- Go 1.19 or later
- Network connectivity to target hosts

## License

This project is released under the MIT License. See LICENSE file for details.
