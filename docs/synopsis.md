# Synopsis
```
A cross-platform command line tool to test and verify SSL/TLS connections
and display certificate files in a human readable format

Usage:
  -dump string
    	Read and display certificate file (PEM or DER format)
  -example
    	Show usage examples and exit
  -host string
    	Target host (required unless given as an argument)
  -output string
    	Output file for saved certificates (default: stdout)
  -port string
    	Target port (443 for HTTPS, 587 for SMTP StartTLS) (default "443")
  -save string
    	Save certificates from connection (formats: pem, der, text)
  -skip-verify
    	Skip certificate verification
  -starttls
    	Use StartTLS (for SMTP-like protocols)
  -timeout duration
    	Connection timeout (default 10s)
  -version
    	Show version information and exit

Save Formats:
  pem   - PEM encoded certificates (Base64 with headers)
  der   - DER encoded certificates (binary format)
  text  - Human-readable text format with details

Use --example to see usage examples

 ssl-tls-verify v1.0.4
 https://github.com/muquit/ssl-tls-verify
 Compiled with go version: go1.27.3

```
