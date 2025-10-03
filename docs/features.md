## Features

### SSL/TLS Connection Testing
- **Direct TLS connections** - Test HTTPS and other TLS-enabled services (port 443)
- **StartTLS support** - Test SMTP, IMAP, and other protocols that upgrade to TLS (port 587)
- **Certificate verification** - Option to skip verification with `--skip-verify` flag
- **Configurable timeouts** - Control connection timeout duration
- **Output** - Human-readable certificate chain display with:
  - TLS version and cipher suite information
  - Complete certificate chain (leaf + intermediates)
  - Subject and Issuer details
  - Validity periods with expiry warnings
  - SHA-1 and SHA-256 fingerprints
  - Key usage and extended key usage
  - Certificate Authority information

### Certificate File Display
- **Read certificate files** - Display any certificate file in human readable format
- **Multiple format support**:
  - PEM format (`.pem`)
  - DER format (`.der`)
- **Automatic format detection** - No need to specify the format
- **Certificate chain display** - Shows all certificates in multi-cert files

### Certificate Export
- **Save from live connections** - Export certificates retrieved from SSL/TLS connections
- **Multiple export formats**:
  - **PEM** - Standard text format with Base64 encoding (full chain supported)
  - **DER** - Binary format (leaf certificate only). User `--save pem` to save the complete certificate chain in one file
  - **TEXT** - Human-readable report format with complete details
- **Flexible output** - Save to file or pipe to stdout (except DER)
- **Chain preservation** - PEM and TEXT formats save complete certificate chains

### Security & Privacy
- **IP address masking** - Optional masking IP address in StartTLS EHLO with `MAILSEND_MASK_IP` environment variable
- **Secure connections** - TLS support (TLS 1.0 - 1.3)
- **Certificate validation** - Full certificate chain verification
### Cross Platform
- Pre-compiled binaries are available in @RELEASES@ page. But can be cross-compiled for any
platforms @GO@ supports. Currently it is more than 40 platforms (Look at
        `platforms.txt` or run `go tool dist list` to see the list of
        supported platforms.
