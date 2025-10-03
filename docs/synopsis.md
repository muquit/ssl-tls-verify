## Synopsis
```
ssl-tls-verify v1.0.2
https://github.com/muquit/ssl-tls-verify
A cross-platform command line tool to test and verify SSL/TLS connections
and display certificate files in a human readable format

Usage:
  -dump string
    	Read and display certificate file (PEM or DER format)
  -host string
    	Target host (default "google.com")
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
  text  - Human-readable text format with full details

Examples:
  Test SSL/TLS connection:
    ssl-tls-verify --host google.com
    ssl-tls-verify --host smtp.gmail.com --port 587 --starttls

  Display certificate files:
    ssl-tls-verify --dump certificate.pem
    ssl-tls-verify --dump certificate.der

  Save certificates from connection:
    ssl-tls-verify --host google.com --save pem --output certs.pem
    ssl-tls-verify --host google.com --save text --output report.txt
    ssl-tls-verify --host smtp.gmail.com --port 587 --starttls --save pem
```
