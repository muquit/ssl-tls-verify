## API Reference

### `pkg/certificate`

#### `ParseFile(filename string) ([]*x509.Certificate, string, error)`
Parse a certificate file (auto-detects PEM or DER format).

**Returns:** certificates, format ("PEM" or "DER"), error

#### `ParsePEMCertificates(data []byte) ([]*x509.Certificate, error)`
Parse PEM-encoded certificates from raw bytes.

#### `Display(certs []*x509.Certificate)`
Display certificates in human-readable format to stdout.

#### `PrintDetails(cert *x509.Certificate)`
Print detailed certificate information to stdout.

#### `PrintDetailsToWriter(cert *x509.Certificate, output io.Writer)`
Print certificate details to a specific writer.

#### `FormatFingerprint(certRaw []byte, algorithm string) string`
Generate certificate fingerprint. Algorithms: "sha1", "sha256"

#### `Save(certs []*x509.Certificate, config SaveConfig) error`
Save certificates in specified format (pem, der, text).

#### `SavePEM(certs []*x509.Certificate, output io.Writer) error`
Save certificates in PEM format.

#### `SaveDER(certs []*x509.Certificate, output io.Writer) error`
Save certificate in DER format (leaf only).

#### `SaveText(certs []*x509.Certificate, output io.Writer) error`
Save certificates in human-readable text format.

### `pkg/connection`

#### `TestDirectTLS(config TLSConfig) ([]*x509.Certificate, tls.ConnectionState, error)`
Test a direct TLS connection (HTTPS, IMAPS, etc.).

**TLSConfig fields:**
- `Host string` - Target hostname
- `Port string` - Target port
- `SkipVerify bool` - Skip certificate verification
- `ConnectTimeout time.Duration` - Connection timeout

#### `TestStartTLS(config StartTLSConfig) ([]*x509.Certificate, tls.ConnectionState, error)`
Test a StartTLS connection (SMTP, IMAP, POP3, etc.).

**StartTLSConfig fields:**
- `Host string` - Target hostname
- `Port string` - Target port
- `SkipVerify bool` - Skip certificate verification
- `ConnectTimeout time.Duration` - Connection timeout

#### `PrintTLSInfo(state tls.ConnectionState)`
Print detailed TLS connection and certificate information.

### `pkg/utils`

#### `MaskIPAddresses(text string) string`
Mask IPv4 and IPv6 addresses in text (requires `MAILSEND_MASK_IP` env var).

#### `MaskIPSlice(addresses []string) []string`
Mask IP addresses in a string slice (requires `MAILSEND_MASK_IP` env var).

---

### Best Practices

1. **Always set timeouts** when making connections to prevent hanging
2. **Handle errors properly** - connection failures are common
3. **Use `SkipVerify: false`** in production unless you have a specific reason
4. **Check certificate expiration** dates if you're monitoring certificates
5. **Use PEM format** for saving full certificate chains
6. **Enable IP masking** when sharing logs or screenshots

### Error Handling

All functions return errors that should be checked:

```go
certs, state, err := connection.TestDirectTLS(config)
if err != nil {
    // Handle different error types
    if strings.Contains(err.Error(), "timeout") {
        log.Println("Connection timeout")
    } else if strings.Contains(err.Error(), "certificate") {
        log.Println("Certificate validation failed")
    } else {
        log.Printf("Connection error: %v", err)
    }
    return
}
```
