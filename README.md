# Table Of Contents
- [Introduction](#introduction)
- [Background](#background)
- [Features](#features)
  - [SSL/TLS Connection Testing](#ssltls-connection-testing)
  - [Certificate File Display](#certificate-file-display)
  - [Certificate Export](#certificate-export)
  - [Security & Privacy](#security-privacy)
  - [Cross Platform](#cross-platform)
- [Synopsis](#synopsis)
- [Latest Version (v1.0.4)](#latest-version-v104)
- [Installation](#installation)
  - [Download pre-compiled binaries](#download-pre-compiled-binaries)
  - [Build from source](#build-from-source)
  - [Install directly](#install-directly)
  - [Installing using Homebrew on Mac/Linux](#installing-using-homebrew-on-maclinux)
    - [Install](#install)
    - [Upgrade](#upgrade)
    - [Uninstall](#uninstall)
    - [Remove the tap](#remove-the-tap)
- [Examples](#examples)
  - [Commands](#commands)
- [Sample output](#sample-output)
  - [google.com HTTPS port 443](#googlecom-https-port-443)
  - [Server using self signed certificate](#server-using-self-signed-certificate)
  - [Skip Verification](#skip-verification)
  - [Certificate chain containing a blocked intermediate CA (eMudhra)](#certificate-chain-containing-a-blocked-intermediate-ca-emudhra)
  - [smtp.gmail.com with StartTLS](#smtpgmailcom-with-starttls)
  - [Save certificates in connections in PEM format](#save-certificates-in-connections-in-pem-format)
  - [Dump the saved PEM certificates to stdout](#dump-the-saved-pem-certificates-to-stdout)
  - [Save leaf certificate in DER format](#save-leaf-certificate-in-der-format)
  - [Dump the saved DER certificate to stdout](#dump-the-saved-der-certificate-to-stdout)
  - [Save to stdout and pipe to other tools e.g. openssl](#save-to-stdout-and-pipe-to-other-tools-eg-openssl)
  - [Saving certs works with StartTLSS too](#saving-certs-works-with-starttlss-too)
- [How to use ssl-tls-verify as a library](#how-to-use-ssl-tls-verify-as-a-library)
  - [Cert checker](#cert-checker)
    - [Create a project directory](#create-a-project-directory)
    - [Create `main.go`](#create-maingo)
    - [Initialize go module](#initialize-go-module)
    - [Build](#build)
    - [Run](#run)
  - [TLS checker](#tls-checker)
    - [Create a project directory](#create-a-project-directory)
    - [Create `main.go`](#create-maingo)
    - [Initialize go module](#initialize-go-module)
    - [Build](#build)
    - [Run](#run)
  - [StartTLS checker](#starttls-checker)
    - [Create a project directory](#create-a-project-directory)
    - [Create `main.go`](#create-maingo)
    - [Initialize go module](#initialize-go-module)
    - [Build](#build)
    - [Run](#run)
  - [Display certificate files](#display-certificate-files)
    - [Create a project directory](#create-a-project-directory)
    - [Create `main.go`](#create-maingo)
    - [Initialize go module](#initialize-go-module)
    - [Build](#build)
    - [Run](#run)
  - [Parse PEM Certificates from memory](#parse-pem-certificates-from-memory)
    - [Create a project directory](#create-a-project-directory)
    - [Create `main.go`](#create-maingo)
    - [Initialize go module](#initialize-go-module)
    - [Build](#build)
    - [Run](#run)
  - [Save certs to different formats](#save-certs-to-different-formats)
    - [Create `main.go`](#create-maingo)
    - [Initialize go module](#initialize-go-module)
    - [Build](#build)
    - [Run](#run)
  - [Generate certificate fingerprints](#generate-certificate-fingerprints)
    - [Create a project directory](#create-a-project-directory)
    - [Create `main.go`](#create-maingo)
    - [Initialize go module](#initialize-go-module)
    - [Build](#build)
    - [Run](#run)
- [API Reference](#api-reference)
  - [`pkg/certificate`](#pkgcertificate)
    - [`ParseFile(filename string) ([]*x509.Certificate, string, error)`](#parsefilefilename-string-x509certificate-string-error)
    - [`ParsePEMCertificates(data []byte) ([]*x509.Certificate, error)`](#parsepemcertificatesdata-byte-x509certificate-error)
    - [`Display(certs []*x509.Certificate)`](#displaycerts-x509certificate)
    - [`PrintDetails(cert *x509.Certificate)`](#printdetailscert-x509certificate)
    - [`PrintDetailsToWriter(cert *x509.Certificate, output io.Writer)`](#printdetailstowritercert-x509certificate-output-iowriter)
    - [`FormatFingerprint(certRaw []byte, algorithm string) string`](#formatfingerprintcertraw-byte-algorithm-string-string)
    - [`Save(certs []*x509.Certificate, config SaveConfig) error`](#savecerts-x509certificate-config-saveconfig-error)
    - [`SavePEM(certs []*x509.Certificate, output io.Writer) error`](#savepemcerts-x509certificate-output-iowriter-error)
    - [`SaveDER(certs []*x509.Certificate, output io.Writer) error`](#savedercerts-x509certificate-output-iowriter-error)
    - [`SaveText(certs []*x509.Certificate, output io.Writer) error`](#savetextcerts-x509certificate-output-iowriter-error)
  - [`pkg/connection`](#pkgconnection)
    - [`TestDirectTLS(config TLSConfig) ([]*x509.Certificate, tls.ConnectionState, error)`](#testdirecttlsconfig-tlsconfig-x509certificate-tlsconnectionstate-error)
    - [`TestStartTLS(config StartTLSConfig) ([]*x509.Certificate, tls.ConnectionState, error)`](#teststarttlsconfig-starttlsconfig-x509certificate-tlsconnectionstate-error)
    - [`PrintTLSInfo(state tls.ConnectionState)`](#printtlsinfostate-tlsconnectionstate)
    - [`PrintVerifyFailure(host string, err error) bool`](#printverifyfailurehost-string-err-error-bool)
  - [`pkg/utils`](#pkgutils)
    - [`MaskIPAddresses(text string) string`](#maskipaddressestext-string-string)
    - [`MaskIPSlice(addresses []string) []string`](#maskipsliceaddresses-string-string)
  - [Best Practices](#best-practices)
  - [Error Handling](#error-handling)
- [License](#license)
- [Authors](#authors)

# Introduction
`ssl-tls-verify` is a simple cross-platform command line tool to test and 
verify SSL/TLS connections with detailed certificate inspection.
It supports both direct TLS connections (like HTTPS) and StartTLS protocol 
upgrades (like SMTP). 
The tool provides comprehensive certificate chain analysis with 
pretty-printed output, making it ideal for debugging SSL/TLS issues, 
certificate validation problems, and connection troubleshooting.

Unlike basic tools like `openssl s_client`, this utility focuses on 
providing human-readable certificate information with expiry warnings, 
fingerprint validation, and detailed TLS connection state analysis. All 
the certificate parsing and validation is handled by [go](https://go.dev/)'s 
`crypto/tls` and `crypto/x509` packages.

If you have any question, request or suggestion, please enter it in the
[Issues](https://github.com/muquit/ssl-tls-verify/issues) with appropriate label.

# Background

There was a lingering bug in [mailsend-go](https://github.com/muquit/mailsend-go) that `-verifyCert` was broken (Bug
        #71)
I was not getting time to look at the code and figure out the issue, so I asked [Claude AI Sonnet 4.5](https://claude.ai) to
write a simple CLI tool in [go](https://go.dev/) to verify certificates from connection. Hence, `ssl-tls-verify` was born.
After creating this tool, I realized I had missed `ServerName` in `tls.Config`.

I also love this tool now, it's much nicer than openssl with its human-readable output,
clear certificate chain display, and intuitive formatting. What started as a debugging
utility has become my go-to tool for SSL/TLS certificate inspection. 

Hope you find the tool useful as well.

# Features

## SSL/TLS Connection Testing
- **Direct TLS connections** - Test HTTPS and other TLS-enabled services (port 443)
- **StartTLS support** - Test SMTP, IMAP, and other protocols that upgrade to TLS (port 587)
- **Certificate verification** - Option to skip verification with `--skip-verify` flag
- **Configurable timeouts** - Control connection timeout duration
- **Verification failure details** - If certificate verification fails, the
certificate chain presented by the server (subject, issuer, validity and
SHA-256 fingerprint) and the result of re-verifying it against the system
roots are printed, in addition to the error
- **Output** - Human-readable certificate chain display with:
  - TLS version and cipher suite information
  - Complete certificate chain (leaf + intermediates)
  - Subject and Issuer details
  - Validity periods with expiry warnings
  - SHA-1 and SHA-256 fingerprints
  - Key usage and extended key usage
  - Certificate Authority information

## Certificate File Display
- **Read certificate files** - Display any certificate file in human readable format
- **Multiple format support**:
  - PEM format (`.pem`)
  - DER format (`.der`)
- **Automatic format detection** - No need to specify the format
- **Certificate chain display** - Shows all certificates in multi-cert files

## Certificate Export
- **Save from live connections** - Export certificates retrieved from SSL/TLS connections
- **Multiple export formats**:
  - **PEM** - Standard text format with Base64 encoding (full chain supported)
  - **DER** - Binary format (leaf certificate only). User `--save pem` to save the complete certificate chain in one file
  - **TEXT** - Human-readable report format with complete details
- **Flexible output** - Save to file or pipe to stdout (except DER)
- **Chain preservation** - PEM and TEXT formats save complete certificate chains

## Security & Privacy
- **IP address masking** - Optional masking IP address in StartTLS EHLO with `MAILSEND_MASK_IP` environment variable
- **Secure connections** - TLS support (TLS 1.0 - 1.3)
- **Certificate validation** - Full certificate chain verification
## Cross Platform
- Pre-compiled binaries are available in [Releases](https://github.com/muquit/ssl-tls-verify/releases) page. But can be cross-compiled for any
platforms [go](https://go.dev/) supports. Currently it is more than 40 platforms (Look at
        `platforms.txt` or run `go tool dist list` to see the list of
        supported platforms.

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

# Latest Version (v1.0.4)
The current version is v1.0.4
Please look at [ChangeLog](ChangeLog.md) for what has changed in the current version.

# Installation

## Download pre-compiled binaries

Please download pre-compiled binaries from [Releases](https://github.com/muquit/ssl-tls-verify/releases) page. After extracting
the archive, copy/rename the binary to somewhere in your path.

## Build from source
```bash
go build -o ssl-tls-verify main.go
# or
make build
```

If you're on Linux/MacOS:
```bash
make install
```
It will compile and do the following:

```bash
    sudo cp ssl-tls-verify /usr/local/bin
    sudo cp docs/ssl-tls-verify.1 /usr/local/share/man/man1
```

## Install directly
```bash
go install github.com/muquit/ssl-tls-verify@latest
```

## Installing using Homebrew on Mac/Linux

You will need to install [Homebrew](https://brew.sh/) first.

### Install

```
brew tap muquit/ssl-tls-verify https://github.com/muquit/ssl-tls-verify.git
brew install muquit/ssl-tls-verify/ssl-tls-verify
```

### Upgrade
```
brew upgrade ssl-tls-verify
```

### Uninstall
```
brew uninstall ssl-tls-verify
```

### Remove the tap
```
brew untap muquit/ssl-tls-verify
```

<hr width="200" align="left">

<sub>Brew install instructions and formula automatically generated by [go-xbuild-go](https://github.com/muquit/go-xbuild-go) v1.0.13 on Oct-31-2026</sub>

<sub>[go-xbuild-go](https://github.com/muquit/go-xbuild-go) writes the formula into the project's own `Formula/project.rb` rather than a central Homebrew tap repo. This keeps the formula version-controlled and committed alongside the code.</sub>

<hr width="200" align="left">

# Examples
## Commands

Test HTTPS connection with certificate verification:
```bash
./ssl-tls-verify --host google.com --port 443
```

The host can also be given as an argument. Flags must come before it:
```bash
./ssl-tls-verify google.com
./ssl-tls-verify --port 8443 example.com
```

Check a host whose certificate chain contains a blocked intermediate CA. It
should fail verification and print the certificate chain presented by the server:
```bash
./ssl-tls-verify dv-g3b-emudhra.we-have-a-problem.com
```

Show usage examples:
```bash
./ssl-tls-verify --example
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
Save in PEM format:
```bash
./ssl-tls-verify --host google.com --save pem --output certs.pem
```

Save in DER format: 
```bash
./ssl-tls-verify --host google.com --save der --output cert.der
```
Note: only left cert is saved in cert.der. If you would like to get all the
certs in DER format, you can convert the certs in pem file using openssl.
Example:
```bash
openssl x509 -in certs.pem -outform DER -out certs.der
```

Save in TEXT format (human-readable):
```bash
./ssl-tls-verify --host google.com --save text --output report.txt
```

Save to stdout (pipe to other tools)
```bash
./ssl-tls-verify --host google.com --save pem | openssl x509 -text
```

Works with StartTLS too
```bash
./ssl-tls-verify --host smtp.gmail.com --port 587 \
        --starttls --save pem --output smtp-certs.pem
```       

# Sample output
## google.com HTTPS port 443
```bash
./ssl-tls-verify --host google.com --port 443
```
```bash
Connecting to google.com:443
Verify Certificate: Yes
Use StartTLS: No


============================================================
TLS CONNECTION INFORMATION
============================================================
TLS Version:           TLS 1.3 (0x0304)
Cipher Suite:          0x1301
Server Certificates:   3
Handshake Complete:    true
Did Resume:            false
Negotiated Protocol:   

------------------------------------------------------------
SERVER CERTIFICATE #1 (Leaf Certificate)
------------------------------------------------------------
Subject:               CN=*.google.com
Issuer:                CN=WE2,O=Google Trust Services,C=US
Serial Number:         301270412142668962639679896844227265037
Version:               3
Valid From:            2026-10-07 19:07:42 UTC
Valid Until:           2026-12-30 19:07:41 UTC
Status:                ✅ Valid (59 days remaining)
DNS Names:             *.google.com, *.appengine.google.com, *.bdn.dev, *.origin-test.bdn.dev, *.cloud.google.com, *.crowdsource.google.com, *.datacompute.google.com, *.google.ca, *.google.cl, *.google.co.in, *.google.co.jp, *.google.co.uk, *.google.com.ar, *.google.com.au, *.google.com.br, *.google.com.co, *.google.com.mx, *.google.com.tr, *.google.com.vn, *.google.de, *.google.es, *.google.fr, *.google.hu, *.google.it, *.google.nl, *.google.pl, *.google.pt, *.gemini.cloud.google.com, *.gstatic.com, *.metric.gstatic.com, *.gvt1.com, *.gcpcdn.gvt1.com, *.gvt2.com, *.gcp.gvt2.com, *.url.google.com, *.youtube-nocookie.com, *.ytimg.com, ai.android, android.com, *.android.com, *.flash.android.com, g.co, *.g.co, goo.gl, www.goo.gl, google-analytics.com, *.google-analytics.com, google.com, googlecommerce.com, *.googlecommerce.com, urchin.com, *.urchin.com, youtu.be, youtube.com, *.youtube.com, music.youtube.com, *.music.youtube.com, youtubeeducation.com, *.youtubeeducation.com, youtubekids.com, *.youtubekids.com, yt.be, *.yt.be, android.clients.google.com, *.aistudio.google.com
Public Key Algorithm:  ECDSA
Signature Algorithm:   ECDSA-SHA256
Key Usage:             Digital Signature
Extended Key Usage:    Server Authentication
SHA-1 Fingerprint:     EA:BC:AB:CD:E0:C7:EA:5E:B3:51:4E:91:8E:95:E3:50:7D:C2:04:D4
SHA-256 Fingerprint:   B0:C4:7F:C7:5D:77:57:CB:45:5D:9F:3F:D7:19:50:0D:36:4D:C6:F4:07:38:BA:A7:0F:AD:F9:57:20:B0:70:B7
Certificate Authority: No

------------------------------------------------------------
INTERMEDIATE CERTIFICATE #2
------------------------------------------------------------
Subject:               CN=WE2,O=Google Trust Services,C=US
Issuer:                CN=GTS Root R4,O=Google Trust Services LLC,C=US
Serial Number:         170074604807150777415902075947543458528
Version:               3
Valid From:            2023-12-13 09:00:00 UTC
Valid Until:           2029-02-20 14:00:00 UTC
Status:                ✅ Valid (842 days remaining)
Public Key Algorithm:  ECDSA
Signature Algorithm:   ECDSA-SHA384
Key Usage:             Digital Signature, Certificate Signing, CRL Signing
Extended Key Usage:    Server Authentication, Client Authentication
SHA-1 Fingerprint:     4D:9A:CB:31:3D:73:DA:2E:9E:B4:51:A7:CF:63:09:AF:45:C9:93:C1
SHA-256 Fingerprint:   9C:3F:2F:D1:1C:57:D7:C6:49:AD:5A:09:32:C0:F0:D2:97:56:F6:A0:A1:C7:4C:43:E1:E8:9A:62:D6:4C:D3:20
Certificate Authority: Yes
Max Path Length:       0

------------------------------------------------------------
INTERMEDIATE CERTIFICATE #3
------------------------------------------------------------
Subject:               CN=GTS Root R4,O=Google Trust Services LLC,C=US
Issuer:                CN=GlobalSign Root CA,OU=Root CA,O=GlobalSign nv-sa,C=BE
Serial Number:         170001980149335831901244157168837298715
Version:               3
Valid From:            2023-11-15 03:43:21 UTC
Valid Until:           2028-01-28 00:00:42 UTC
Status:                ✅ Valid (453 days remaining)
Public Key Algorithm:  ECDSA
Signature Algorithm:   SHA256-RSA
Key Usage:             Digital Signature, Certificate Signing, CRL Signing
Extended Key Usage:    Server Authentication, Client Authentication
SHA-1 Fingerprint:     93:2B:ED:33:9A:A6:92:12:C8:93:75:B7:93:04:B4:75:49:0B:89:A0
SHA-256 Fingerprint:   76:B2:7B:80:A5:80:27:DC:3C:F1:DA:68:DA:C1:70:10:ED:93:99:7D:0B:60:3E:2F:AD:BE:85:01:24:93:B5:A7
Certificate Authority: Yes
Direct TLS connection successful!
```
## Server using self signed certificate
```bash
# Server using self-signed certificate
./ssl-tls-verify -host 127.0.0.1 -port 8881
```
```bash
Connecting to 127.0.0.1:8881
Verify Certificate: Yes
Use StartTLS: No

2026/10/31 19:49:13 Direct TLS connection failed: failed to establish TLS connection: dial tcp 127.0.0.1:8881: connect: connection refused
```
## Skip Verification
```bash
# Skip verification
./ssl-tls-verify -host 127.0.0.1 -port 8881 -skip-verify
```
```bash
Connecting to 127.0.0.1:8881
Verify Certificate: No
Use StartTLS: No

2026/10/31 19:49:13 Direct TLS connection failed: failed to establish TLS connection: dial tcp 127.0.0.1:8881: connect: connection refused
```
## Certificate chain containing a blocked intermediate CA (eMudhra)
```bash
./ssl-tls-verify dv-g3b-emudhra.we-have-a-problem.com
```
```bash
Connecting to dv-g3b-emudhra.we-have-a-problem.com:443
Verify Certificate: Yes
Use StartTLS: No


Certificate verification FAILED for dv-g3b-emudhra.we-have-a-problem.com
Reason: x509: “dv-g3b-emudhra.we-have-a-problem.com” certificate is blocked

Chain presented by the server (3 certificate(s)):

 [0] Leaf
     Subject:    CN=dv-g3b-emudhra.we-have-a-problem.com
     Issuer:     CN=EM DV TLS CA - G3B,O=eMudhra BV,C=NL
     Valid:      2026-10-20 to 2027-05-03
     SHA-256:    9D:CB:21:B1:84:1D:75:31:3B:3C:88:8F:4B:B9:CC:01:F8:F5:13:79:A7:6F:BA:A2:41:CA:49:2B:66:88:39:E7

 [1] Intermediate
     Subject:    CN=EM DV TLS CA - G3B,O=eMudhra BV,C=NL
     Issuer:     CN=emSign Root TLS CA - G3,O=eMudhra Technologies Limited,C=IN
     Valid:      2025-04-30 to 2040-04-28
     SHA-256:    42:C4:AC:CD:12:C4:F5:95:97:92:5E:1A:08:A7:A9:01:02:60:85:43:80:93:08:D5:0E:C4:27:E1:18:BA:14:93

 [2] Intermediate
     Subject:    CN=emSign Root TLS CA - G3,O=eMudhra Technologies Limited,C=IN
     Issuer:     CN=emSign ECC Root CA - G3,OU=emSign PKI,O=eMudhra Technologies Limited,C=IN
     Valid:      2024-07-11 to 2039-07-09
     SHA-256:    30:94:02:57:92:03:9A:68:AF:6E:55:36:86:41:0B:DA:B7:C0:D4:A3:6B:90:84:CB:A6:42:02:E3:5D:6E:BB:DD

Re-verification against system roots: x509: “dv-g3b-emudhra.we-have-a-problem.com” certificate is blocked

2026/10/31 19:49:14 Direct TLS connection failed: failed to establish TLS connection: tls: failed to verify certificate: x509: “dv-g3b-emudhra.we-have-a-problem.com” certificate is blocked
```
## smtp.gmail.com with StartTLS
```bash
./ssl-tls-verify --host smtp.gmail.com --port 587 --starttls
```
```bash
Connecting to smtp.gmail.com:587
Verify Certificate: Yes
Use StartTLS: Yes

Plain connection established, attempting StartTLS...
Server greeting: 220 smtp.gmail.com ESMTP 6a1803df08f44-91ca804dc9fsm53104106d6.24 - gsmtp
EHLO response: 250-smtp.gmail.com at your service, [xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx]
250-SIZE 35882577
250-8BITMIME
250-STARTTLS
250-ENHANCEDSTATUSCODES
250-PIPELINING
250-CHUNKING
250 SMTPUTF8
STARTTLS response: 220 2.0.0 Ready to start TLS
StartTLS upgrade successful!

============================================================
TLS CONNECTION INFORMATION
============================================================
TLS Version:           TLS 1.3 (0x0304)
Cipher Suite:          0x1301
Server Certificates:   3
Handshake Complete:    true
Did Resume:            false
Negotiated Protocol:   

------------------------------------------------------------
SERVER CERTIFICATE #1 (Leaf Certificate)
------------------------------------------------------------
Subject:               CN=smtp.gmail.com
Issuer:                CN=WE2,O=Google Trust Services,C=US
Serial Number:         242554843314789844639913615394864740594
Version:               3
Valid From:            2026-10-07 19:09:08 UTC
Valid Until:           2026-12-30 19:09:07 UTC
Status:                ✅ Valid (59 days remaining)
DNS Names:             smtp.gmail.com
Public Key Algorithm:  ECDSA
Signature Algorithm:   ECDSA-SHA256
Key Usage:             Digital Signature
Extended Key Usage:    Server Authentication
SHA-1 Fingerprint:     A4:84:AE:63:CD:C8:0A:4B:34:D6:F8:05:72:BF:6A:3A:92:AD:5C:94
SHA-256 Fingerprint:   4D:AC:A5:3C:17:7F:CE:DF:D7:BE:D7:F4:4B:60:3E:F8:AA:67:B1:DC:08:36:05:31:61:3F:68:72:EE:B0:66:AF
Certificate Authority: No

------------------------------------------------------------
INTERMEDIATE CERTIFICATE #2
------------------------------------------------------------
Subject:               CN=WE2,O=Google Trust Services,C=US
Issuer:                CN=GTS Root R4,O=Google Trust Services LLC,C=US
Serial Number:         170074604807150777415902075947543458528
Version:               3
Valid From:            2023-12-13 09:00:00 UTC
Valid Until:           2029-02-20 14:00:00 UTC
Status:                ✅ Valid (842 days remaining)
Public Key Algorithm:  ECDSA
Signature Algorithm:   ECDSA-SHA384
Key Usage:             Digital Signature, Certificate Signing, CRL Signing
Extended Key Usage:    Server Authentication, Client Authentication
SHA-1 Fingerprint:     4D:9A:CB:31:3D:73:DA:2E:9E:B4:51:A7:CF:63:09:AF:45:C9:93:C1
SHA-256 Fingerprint:   9C:3F:2F:D1:1C:57:D7:C6:49:AD:5A:09:32:C0:F0:D2:97:56:F6:A0:A1:C7:4C:43:E1:E8:9A:62:D6:4C:D3:20
Certificate Authority: Yes
Max Path Length:       0

------------------------------------------------------------
INTERMEDIATE CERTIFICATE #3
------------------------------------------------------------
Subject:               CN=GTS Root R4,O=Google Trust Services LLC,C=US
Issuer:                CN=GlobalSign Root CA,OU=Root CA,O=GlobalSign nv-sa,C=BE
Serial Number:         170001980149335831901244157168837298715
Version:               3
Valid From:            2023-11-15 03:43:21 UTC
Valid Until:           2028-01-28 00:00:42 UTC
Status:                ✅ Valid (453 days remaining)
Public Key Algorithm:  ECDSA
Signature Algorithm:   SHA256-RSA
Key Usage:             Digital Signature, Certificate Signing, CRL Signing
Extended Key Usage:    Server Authentication, Client Authentication
SHA-1 Fingerprint:     93:2B:ED:33:9A:A6:92:12:C8:93:75:B7:93:04:B4:75:49:0B:89:A0
SHA-256 Fingerprint:   76:B2:7B:80:A5:80:27:DC:3C:F1:DA:68:DA:C1:70:10:ED:93:99:7D:0B:60:3E:2F:AD:BE:85:01:24:93:B5:A7
Certificate Authority: Yes
StartTLS connection successful!
```
## Save certificates in connections in PEM format
```bash
./ssl-tls-verify --host google.com --save pem --output certs.pem
```
```bash
Connecting to google.com:443
Verify Certificate: Yes
Use StartTLS: No
Save Format: PEM
Output File: certs.pem


✅ Successfully saved 3 certificate(s) in PEM format to certs.pem
```
## Dump the saved PEM certificates to stdout 
```bash
./ssl-tls-verify --host google.com --dump ./certs.pem
```
```bash
Reading certificate file: ./certs.pem
Format: PEM

============================================================
CERTIFICATE DETAILS
Total Certificates: 3
============================================================

------------------------------------------------------------
CERTIFICATE #1 (Leaf Certificate)
------------------------------------------------------------
Subject:               CN=*.google.com
Issuer:                CN=WE2,O=Google Trust Services,C=US
Serial Number:         301270412142668962639679896844227265037
Version:               3
Valid From:            2026-10-07 19:07:42 UTC
Valid Until:           2026-12-30 19:07:41 UTC
Status:                ✅ Valid (59 days remaining)
DNS Names:             *.google.com, *.appengine.google.com, *.bdn.dev, *.origin-test.bdn.dev, *.cloud.google.com, *.crowdsource.google.com, *.datacompute.google.com, *.google.ca, *.google.cl, *.google.co.in, *.google.co.jp, *.google.co.uk, *.google.com.ar, *.google.com.au, *.google.com.br, *.google.com.co, *.google.com.mx, *.google.com.tr, *.google.com.vn, *.google.de, *.google.es, *.google.fr, *.google.hu, *.google.it, *.google.nl, *.google.pl, *.google.pt, *.gemini.cloud.google.com, *.gstatic.com, *.metric.gstatic.com, *.gvt1.com, *.gcpcdn.gvt1.com, *.gvt2.com, *.gcp.gvt2.com, *.url.google.com, *.youtube-nocookie.com, *.ytimg.com, ai.android, android.com, *.android.com, *.flash.android.com, g.co, *.g.co, goo.gl, www.goo.gl, google-analytics.com, *.google-analytics.com, google.com, googlecommerce.com, *.googlecommerce.com, urchin.com, *.urchin.com, youtu.be, youtube.com, *.youtube.com, music.youtube.com, *.music.youtube.com, youtubeeducation.com, *.youtubeeducation.com, youtubekids.com, *.youtubekids.com, yt.be, *.yt.be, android.clients.google.com, *.aistudio.google.com
Public Key Algorithm:  ECDSA
Signature Algorithm:   ECDSA-SHA256
Key Usage:             Digital Signature
Extended Key Usage:    Server Authentication
SHA-1 Fingerprint:     EA:BC:AB:CD:E0:C7:EA:5E:B3:51:4E:91:8E:95:E3:50:7D:C2:04:D4
SHA-256 Fingerprint:   B0:C4:7F:C7:5D:77:57:CB:45:5D:9F:3F:D7:19:50:0D:36:4D:C6:F4:07:38:BA:A7:0F:AD:F9:57:20:B0:70:B7
Certificate Authority: No

------------------------------------------------------------
CERTIFICATE #2 (Intermediate Certificate)
------------------------------------------------------------
Subject:               CN=WE2,O=Google Trust Services,C=US
Issuer:                CN=GTS Root R4,O=Google Trust Services LLC,C=US
Serial Number:         170074604807150777415902075947543458528
Version:               3
Valid From:            2023-12-13 09:00:00 UTC
Valid Until:           2029-02-20 14:00:00 UTC
Status:                ✅ Valid (842 days remaining)
Public Key Algorithm:  ECDSA
Signature Algorithm:   ECDSA-SHA384
Key Usage:             Digital Signature, Certificate Signing, CRL Signing
Extended Key Usage:    Server Authentication, Client Authentication
SHA-1 Fingerprint:     4D:9A:CB:31:3D:73:DA:2E:9E:B4:51:A7:CF:63:09:AF:45:C9:93:C1
SHA-256 Fingerprint:   9C:3F:2F:D1:1C:57:D7:C6:49:AD:5A:09:32:C0:F0:D2:97:56:F6:A0:A1:C7:4C:43:E1:E8:9A:62:D6:4C:D3:20
Certificate Authority: Yes
Max Path Length:       0

------------------------------------------------------------
CERTIFICATE #3 (Intermediate Certificate)
------------------------------------------------------------
Subject:               CN=GTS Root R4,O=Google Trust Services LLC,C=US
Issuer:                CN=GlobalSign Root CA,OU=Root CA,O=GlobalSign nv-sa,C=BE
Serial Number:         170001980149335831901244157168837298715
Version:               3
Valid From:            2023-11-15 03:43:21 UTC
Valid Until:           2028-01-28 00:00:42 UTC
Status:                ✅ Valid (453 days remaining)
Public Key Algorithm:  ECDSA
Signature Algorithm:   SHA256-RSA
Key Usage:             Digital Signature, Certificate Signing, CRL Signing
Extended Key Usage:    Server Authentication, Client Authentication
SHA-1 Fingerprint:     93:2B:ED:33:9A:A6:92:12:C8:93:75:B7:93:04:B4:75:49:0B:89:A0
SHA-256 Fingerprint:   76:B2:7B:80:A5:80:27:DC:3C:F1:DA:68:DA:C1:70:10:ED:93:99:7D:0B:60:3E:2F:AD:BE:85:01:24:93:B5:A7
Certificate Authority: Yes
```
## Save leaf certificate in DER format
```bash
./ssl-tls-verify --host google.com --save der --output ./cert.der
```
```bash
Connecting to google.com:443
Verify Certificate: Yes
Use StartTLS: No
Save Format: DER
Output File: ./cert.der


✅ Successfully saved leaf certificate in DER format to ./cert.der
   Note: Saved leaf certificate only. 2 intermediate certificate(s) were not saved.
   Use PEM format (--save pem) to save the complete certificate chain in one file.
```
## Dump the saved DER certificate to stdout 
```bash
Reading certificate file: ./cert.der
Format: DER

============================================================
CERTIFICATE DETAILS
Total Certificates: 1
============================================================

------------------------------------------------------------
CERTIFICATE
------------------------------------------------------------
Subject:               CN=*.google.com
Issuer:                CN=WE2,O=Google Trust Services,C=US
Serial Number:         301270412142668962639679896844227265037
Version:               3
Valid From:            2026-10-07 19:07:42 UTC
Valid Until:           2026-12-30 19:07:41 UTC
Status:                ✅ Valid (59 days remaining)
DNS Names:             *.google.com, *.appengine.google.com, *.bdn.dev, *.origin-test.bdn.dev, *.cloud.google.com, *.crowdsource.google.com, *.datacompute.google.com, *.google.ca, *.google.cl, *.google.co.in, *.google.co.jp, *.google.co.uk, *.google.com.ar, *.google.com.au, *.google.com.br, *.google.com.co, *.google.com.mx, *.google.com.tr, *.google.com.vn, *.google.de, *.google.es, *.google.fr, *.google.hu, *.google.it, *.google.nl, *.google.pl, *.google.pt, *.gemini.cloud.google.com, *.gstatic.com, *.metric.gstatic.com, *.gvt1.com, *.gcpcdn.gvt1.com, *.gvt2.com, *.gcp.gvt2.com, *.url.google.com, *.youtube-nocookie.com, *.ytimg.com, ai.android, android.com, *.android.com, *.flash.android.com, g.co, *.g.co, goo.gl, www.goo.gl, google-analytics.com, *.google-analytics.com, google.com, googlecommerce.com, *.googlecommerce.com, urchin.com, *.urchin.com, youtu.be, youtube.com, *.youtube.com, music.youtube.com, *.music.youtube.com, youtubeeducation.com, *.youtubeeducation.com, youtubekids.com, *.youtubekids.com, yt.be, *.yt.be, android.clients.google.com, *.aistudio.google.com
Public Key Algorithm:  ECDSA
Signature Algorithm:   ECDSA-SHA256
Key Usage:             Digital Signature
Extended Key Usage:    Server Authentication
SHA-1 Fingerprint:     EA:BC:AB:CD:E0:C7:EA:5E:B3:51:4E:91:8E:95:E3:50:7D:C2:04:D4
SHA-256 Fingerprint:   B0:C4:7F:C7:5D:77:57:CB:45:5D:9F:3F:D7:19:50:0D:36:4D:C6:F4:07:38:BA:A7:0F:AD:F9:57:20:B0:70:B7
Certificate Authority: No
```
## Save to stdout and pipe to other tools e.g. openssl 
```
./ssl-tls-verify --host google.com --save pem | openssl x509 -text -noout
```
```bash
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            e2:a6:92:20:a0:74:04:f5:12:83:53:c5:0f:ef:42:0d
        Signature Algorithm: ecdsa-with-SHA256
        Issuer: C=US, O=Google Trust Services, CN=WE2
        Validity
            Not Before: Oct  7 19:07:42 2026 GMT
            Not After : Dec 30 19:07:41 2026 GMT
        Subject: CN=*.google.com
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit field, 128 bit security level)
                pub:
                    04:ef:99:8c:9b:59:60:51:b3:9b:f0:f0:d3:2f:d1:76:
                    bc:22:41:35:de:3f:eb:d8:7d:f8:f4:42:ef:9d:b3:24:
                    40:ce:71:81:15:56:3f:bc:28:1a:f4:07:f2:82:63:3e:
                    2e:3b:44:79:6b:ae:e6:72:a0:f6:a0:17:07:e2:ab:19:
                    b6
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature
            X509v3 Extended Key Usage: 
                TLS Web Server Authentication
            X509v3 Basic Constraints: critical
                CA:FALSE
            X509v3 Subject Key Identifier: 
                16:36:71:DA:BB:5F:DF:02:9E:61:04:30:DB:8C:5F:C6:5D:89:3F:16
            X509v3 Authority Key Identifier: 
                75:BE:C4:77:AE:89:F6:44:37:7D:CF:B1:68:1F:1D:1A:EB:DC:34:59
            Authority Information Access: 
                CA Issuers - URI:http://i.pki.goog/we2.crt
            X509v3 Subject Alternative Name: 
                DNS:*.google.com, DNS:*.appengine.google.com, DNS:*.bdn.dev, DNS:*.origin-test.bdn.dev, DNS:*.cloud.google.com, DNS:*.crowdsource.google.com, DNS:*.datacompute.google.com, DNS:*.google.ca, DNS:*.google.cl, DNS:*.google.co.in, DNS:*.google.co.jp, DNS:*.google.co.uk, DNS:*.google.com.ar, DNS:*.google.com.au, DNS:*.google.com.br, DNS:*.google.com.co, DNS:*.google.com.mx, DNS:*.google.com.tr, DNS:*.google.com.vn, DNS:*.google.de, DNS:*.google.es, DNS:*.google.fr, DNS:*.google.hu, DNS:*.google.it, DNS:*.google.nl, DNS:*.google.pl, DNS:*.google.pt, DNS:*.gemini.cloud.google.com, DNS:*.gstatic.com, DNS:*.metric.gstatic.com, DNS:*.gvt1.com, DNS:*.gcpcdn.gvt1.com, DNS:*.gvt2.com, DNS:*.gcp.gvt2.com, DNS:*.url.google.com, DNS:*.youtube-nocookie.com, DNS:*.ytimg.com, DNS:ai.android, DNS:android.com, DNS:*.android.com, DNS:*.flash.android.com, DNS:g.co, DNS:*.g.co, DNS:goo.gl, DNS:www.goo.gl, DNS:google-analytics.com, DNS:*.google-analytics.com, DNS:google.com, DNS:googlecommerce.com, DNS:*.googlecommerce.com, DNS:urchin.com, DNS:*.urchin.com, DNS:youtu.be, DNS:youtube.com, DNS:*.youtube.com, DNS:music.youtube.com, DNS:*.music.youtube.com, DNS:youtubeeducation.com, DNS:*.youtubeeducation.com, DNS:youtubekids.com, DNS:*.youtubekids.com, DNS:yt.be, DNS:*.yt.be, DNS:android.clients.google.com, DNS:*.aistudio.google.com
            X509v3 Certificate Policies: 
                Policy: 2.23.140.1.2.1
            X509v3 CRL Distribution Points: 
                Full Name:
                  URI:http://c.pki.goog/we2/dTM3-0hpWfE.crl

            CT Precertificate SCTs: 
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : D8:09:55:3B:94:4F:7A:FF:C8:16:19:6F:94:4F:85:AB:
                                B0:F8:FC:5E:87:55:26:0F:15:D1:2E:72:BB:45:4B:14
                    Timestamp : Oct  7 20:07:44.021 2026 GMT
                    Extensions: none
                    Signature : ecdsa-with-SHA256
                                30:45:02:20:4D:E0:01:03:DC:C0:96:FC:AD:96:80:8A:
                                FD:C7:D3:DE:89:06:66:2E:F4:99:5C:7E:0D:18:48:56:
                                F5:E5:46:D3:02:21:00:9F:24:C9:D9:28:69:28:70:77:
                                44:28:E0:EA:FB:46:8E:40:FD:F5:91:85:DB:9E:31:9A:
                                7F:C4:84:EE:B1:4C:17
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : 8E:CA:47:0B:AC:DE:6A:F3:A2:06:B0:A4:7A:84:B7:46:
                                FE:1F:C6:BF:95:3E:25:E6:9B:4E:E4:02:48:F3:C6:E8
                    Timestamp : Oct  7 20:07:44.862 2026 GMT
                    Extensions: 00:00:05:00:1E:B5:E1:B0
                    Signature : ecdsa-with-SHA256
                                30:45:02:21:00:D3:D7:55:EA:7F:59:70:0F:BB:A2:2F:
                                A2:85:A0:E5:6D:08:B2:7D:6E:53:15:FA:CD:CC:02:03:
                                42:9E:C7:89:FD:02:20:32:9C:44:3A:04:65:E0:D3:B4:
                                E7:EF:A8:CC:6A:DF:BE:9C:97:3E:70:1B:53:C6:99:5A:
                                81:38:02:57:93:BF:E1
    Signature Algorithm: ecdsa-with-SHA256
    Signature Value:
        30:46:02:21:00:96:2b:8b:0d:6d:a6:fa:83:98:45:6b:5a:3a:a7:0e:38:0c:c0:fa:
        d0:56:ce:23:8d:2e:52:e6:b6:c1:16:60:1f:02:21:00:fa:f5:9f:93:ce:fb:7c:62:
        1b:d7:cb:06:99:ca:a4:af:45:d7:fc:7f:be:e6:7d:e2:ff:9c:16:f8:6a:98:e5:c9
```
## Saving certs works with StartTLSS too
```
./ssl-tls-verify --host smtp.gmail.com --port 587 \
    --starttls --save pem --output smtp-certs.pem
```
```bash
Connecting to smtp.gmail.com:587
Verify Certificate: Yes
Use StartTLS: Yes
Save Format: PEM
Output File: smtp-certs.pem

Plain connection established, attempting StartTLS...
Server greeting: 220 smtp.gmail.com ESMTP 6a1803df08f44-91ca7f4367dsm52161066d6.20 - gsmtp
EHLO response: 250-smtp.gmail.com at your service, [xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx]
250-SIZE 35882577
250-8BITMIME
250-STARTTLS
250-ENHANCEDSTATUSCODES
250-PIPELINING
250-CHUNKING
250 SMTPUTF8
STARTTLS response: 220 2.0.0 Ready to start TLS
StartTLS upgrade successful!

✅ Successfully saved 3 certificate(s) in PEM format to smtp-certs.pem
```

# How to use ssl-tls-verify as a library

`ssl-tls-verify` can be used a libary with other projects. Some examples are
shown below

## Cert checker

### Create a project directory
```bash
mkdir cert_checker
cd cert_checker
```
### Create `main.go`

```go
package main

// Example of using ssl-tls-verify as a library
// requires ssl-tls-verify v1.0.3+
// Oct-05-2025 
import (
    "fmt"
    "time"

    "github.com/muquit/ssl-tls-verify/pkg/connection"
)

func main() {
    hosts := []string{"google.com", "github.com", "stackoverflow.com"}

    for _, host := range hosts {
        config := connection.TLSConfig{
            Host:           host,
            Port:           "443",
            SkipVerify:     false,
            ConnectTimeout: 5 * time.Second,
        }

        certs, _, err := connection.TestDirectTLS(config)
        if err != nil {
			fmt.Printf("ERROR: %s - Failed: %v\n", host, err)
            continue
        }

        leafCert := certs[0]
        daysUntilExpiry := int(time.Until(leafCert.NotAfter).Hours() / 24)

        if daysUntilExpiry < 30 {
			fmt.Printf("WARING: %s - Expires in %d days\n", host, daysUntilExpiry)
        } else {
			fmt.Printf("OK: %s - Valid for %d days\n", host, daysUntilExpiry)
        }
    }
}
```
### Initialize go module

```bash
go mod init cert_checker
go mod tidy
```

### Build
```bash
go build .
```

### Run
```bash
./cert_checker
```
```bash
OK: google.com - Valid for 63 days
OK: github.com - Valid for 123 days
OK: stackoverflow.com - Valid for 49 days
```

## TLS checker
### Create a project directory
```bash
mkdir tls_checker
cd tls_checker
```
### Create `main.go`

```go
package main

// Check TLS connection
import (
    "fmt"
    "log"
    "time"

    "github.com/muquit/ssl-tls-verify/pkg/connection"
)

func main() {
    config := connection.TLSConfig{
        Host:           "google.com",
        Port:           "443",
        SkipVerify:     false,
        ConnectTimeout: 10 * time.Second,
    }

    certs, state, err := connection.TestDirectTLS(config)
    if err != nil {
        log.Fatalf("Connection failed: %v", err)
    }

    fmt.Printf("Connected successfully!\n")
    fmt.Printf("TLS Version: 0x%04x\n", state.Version)
    fmt.Printf("Retrieved %d certificate(s)\n", len(certs))

    connection.PrintTLSInfo(state)
}
```
### Initialize go module
```bash
go mod init tls_checker
go mod tidy
```

### Build
```bash
go build .
```

### Run
```bash
./tls_checker
```

```bash
Connected successfully!
TLS Version: 0x0304
Retrieved 3 certificate(s)

============================================================
TLS CONNECTION INFORMATION
============================================================
TLS Version:           TLS 1.3 (0x0304)
Cipher Suite:          0x1301
Server Certificates:   3
Handshake Complete:    true
Did Resume:            false
Negotiated Protocol:

------------------------------------------------------------
SERVER CERTIFICATE #1 (Leaf Certificate)
------------------------------------------------------------
Subject:               CN=*.google.com
Issuer:                CN=WR2,O=Google Trust Services,C=US
Serial Number:         166700031870418575880302779524424518776
Version:               3
Valid From:            2025-09-15 08:34:18 UTC
Valid Until:           2025-12-08 08:34:17 UTC
Status:                ✅ Valid (63 days remaining)
DNS Names:             *.google.com, *.appengine.google.com, *.bdn.dev, *.origin-test.bdn.dev, *.cloud.google.com, *.crowdsource.google.com, *.datacompute.google.com, *.google.ca, *.google.cl, *.google.co.in, *.google.co.jp, *.google.co.uk, *.google.com.ar, *.google.com.au, *.google.com.br, *.google.com.co, *.google.com.mx, *.google.com.tr, *.google.com.vn, *.google.de, *.google.es, *.google.fr, *.google.hu, *.google.it, *.google.nl, *.google.pl, *.google.pt, *.googleapis.cn, *.googlevideo.com, *.gstatic.cn, *.gstatic-cn.com, googlecnapps.cn, *.googlecnapps.cn, googleapps-cn.com, *.googleapps-cn.com, gkecnapps.cn, *.gkecnapps.cn, googledownloads.cn, *.googledownloads.cn, recaptcha.net.cn, *.recaptcha.net.cn, recaptcha-cn.net, *.recaptcha-cn.net, widevine.cn, *.widevine.cn, ampproject.org.cn, *.ampproject.org.cn, ampproject.net.cn, *.ampproject.net.cn, google-analytics-cn.com, *.google-analytics-cn.com, googleadservices-cn.com, *.googleadservices-cn.com, googlevads-cn.com, *.googlevads-cn.com, googleapis-cn.com, *.googleapis-cn.com, googleoptimize-cn.com, *.googleoptimize-cn.com, doubleclick-cn.net, *.doubleclick-cn.net, *.fls.doubleclick-cn.net, *.g.doubleclick-cn.net, doubleclick.cn, *.doubleclick.cn, *.fls.doubleclick.cn, *.g.doubleclick.cn, dartsearch-cn.net, *.dartsearch-cn.net, googletraveladservices-cn.com, *.googletraveladservices-cn.com, googletagservices-cn.com, *.googletagservices-cn.com, googletagmanager-cn.com, *.googletagmanager-cn.com, googlesyndication-cn.com, *.googlesyndication-cn.com, *.safeframe.googlesyndication-cn.com, app-measurement-cn.com, *.app-measurement-cn.com, gvt1-cn.com, *.gvt1-cn.com, gvt2-cn.com, *.gvt2-cn.com, 2mdn-cn.net, *.2mdn-cn.net, googleflights-cn.net, *.googleflights-cn.net, admob-cn.com, *.admob-cn.com, *.gemini.cloud.google.com, googlesandbox-cn.com, *.googlesandbox-cn.com, *.safenup.googlesandbox-cn.com, *.gstatic.com, *.metric.gstatic.com, *.gvt1.com, *.gcpcdn.gvt1.com, *.gvt2.com, *.gcp.gvt2.com, *.url.google.com, *.youtube-nocookie.com, *.ytimg.com, ai.android, android.com, *.android.com, *.flash.android.com, g.cn, *.g.cn, g.co, *.g.co, goo.gl, www.goo.gl, google-analytics.com, *.google-analytics.com, google.com, googlecommerce.com, *.googlecommerce.com, ggpht.cn, *.ggpht.cn, urchin.com, *.urchin.com, youtu.be, youtube.com, *.youtube.com, music.youtube.com, *.music.youtube.com, youtubeeducation.com, *.youtubeeducation.com, youtubekids.com, *.youtubekids.com, yt.be, *.yt.be, android.clients.google.com, *.android.google.cn, *.chrome.google.cn, *.developers.google.cn, *.aistudio.google.com
Public Key Algorithm:  ECDSA
Signature Algorithm:   SHA256-RSA
Key Usage:             Digital Signature
Extended Key Usage:    Server Authentication
SHA-1 Fingerprint:     33:09:D4:D3:61:83:44:73:CF:04:5A:44:53:2E:B5:36:64:BB:FC:7F
SHA-256 Fingerprint:   86:F0:16:7D:8B:24:BC:17:6C:2B:06:E6:05:A6:33:43:10:A0:61:A6:BA:81:FB:58:ED:0A:DC:0B:AE:5D:08:C6
Certificate Authority: No

------------------------------------------------------------
INTERMEDIATE CERTIFICATE #2
------------------------------------------------------------
Subject:               CN=WR2,O=Google Trust Services,C=US
Issuer:                CN=GTS Root R1,O=Google Trust Services LLC,C=US
Serial Number:         170058220837755766831192027518741805976
Version:               3
Valid From:            2023-12-13 09:00:00 UTC
Valid Until:           2029-02-20 14:00:00 UTC
Status:                ✅ Valid (1233 days remaining)
Public Key Algorithm:  RSA
Signature Algorithm:   SHA256-RSA
Key Usage:             Digital Signature, Certificate Signing, CRL Signing
Extended Key Usage:    Server Authentication, Client Authentication
SHA-1 Fingerprint:     66:E4:16:12:60:B1:00:FE:E0:DE:28:7A:9A:52:93:B4:C2:22:4A:E6
SHA-256 Fingerprint:   E6:FE:22:BF:45:E4:F0:D3:B8:5C:59:E0:2C:0F:49:54:18:E1:EB:8D:32:10:F7:88:D4:8C:D5:E1:CB:54:7C:D4
Certificate Authority: Yes
Max Path Length:       0

------------------------------------------------------------
INTERMEDIATE CERTIFICATE #3
------------------------------------------------------------
Subject:               CN=GTS Root R1,O=Google Trust Services LLC,C=US
Issuer:                CN=GlobalSign Root CA,OU=Root CA,O=GlobalSign nv-sa,C=BE
Serial Number:         159159747900478145820483398898491642637
Version:               3
Valid From:            2020-06-19 00:00:42 UTC
Valid Until:           2028-01-28 00:00:42 UTC
Status:                ✅ Valid (844 days remaining)
Public Key Algorithm:  RSA
Signature Algorithm:   SHA256-RSA
Key Usage:             Digital Signature, Certificate Signing, CRL Signing
SHA-1 Fingerprint:     08:74:54:87:E8:91:C1:9E:30:78:C1:F2:A0:7E:45:29:50:EF:36:F6
SHA-256 Fingerprint:   3E:E0:27:8D:F7:1F:A3:C1:25:C4:CD:48:7F:01:D7:74:69:4E:6F:C5:7E:0C:D9:4C:24:EF:D7:69:13:39:18:E5
Certificate Authority: Yes
```
## StartTLS checker

### Create a project directory
```bash
mkdir starttls_checker
cd starttls_checker
```
### Create `main.go`

```go
package main

// StartTLS checker

import (
    "fmt"
    "log"
    "time"

    "github.com/muquit/ssl-tls-verify/pkg/connection"
)

func main() {
    config := connection.StartTLSConfig{
        Host:           "smtp.gmail.com",
        Port:           "587",
        SkipVerify:     false,
        ConnectTimeout: 10 * time.Second,
    }

    certs, state, err := connection.TestStartTLS(config)
    if err != nil {
        log.Fatalf("StartTLS failed: %v", err)
    }

    fmt.Printf("StartTLS successful!\n")
    fmt.Printf("Retrieved %d certificate(s)\n", len(certs))

    // Print connection info
    connection.PrintTLSInfo(state)
}
```
### Initialize go module
```bash
go mod init starttls_checker
go mod tidy
```

### Build
```bash
go build .
```

### Run

```bash
./starttls_checker
```

```bash
Plain connection established, attempting StartTLS...
Server greeting: 220 smtp.gmail.com ESMTP d75a77b69052e-4e55aa44901sm96789721cf.18 - gsmtp
EHLO response: 250-smtp.gmail.com at your service, [xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx]
250-SIZE 35882577
250-8BITMIME
250-STARTTLS
250-ENHANCEDSTATUSCODES
250-PIPELINING
250-CHUNKING
250 SMTPUTF8
STARTTLS response: 220 2.0.0 Ready to start TLS
StartTLS upgrade successful!
StartTLS successful!
Retrieved 3 certificate(s)

============================================================
TLS CONNECTION INFORMATION
============================================================
TLS Version:           TLS 1.3 (0x0304)
Cipher Suite:          0x1301
Server Certificates:   3
Handshake Complete:    true
Did Resume:            false
Negotiated Protocol:

------------------------------------------------------------
SERVER CERTIFICATE #1 (Leaf Certificate)
------------------------------------------------------------
Subject:               CN=smtp.gmail.com
Issuer:                CN=WR2,O=Google Trust Services,C=US
Serial Number:         198559767591087662403673263944902778968
Version:               3
Valid From:            2025-09-15 08:36:10 UTC
Valid Until:           2025-12-08 08:36:09 UTC
Status:                ✅ Valid (63 days remaining)
DNS Names:             smtp.gmail.com
Public Key Algorithm:  ECDSA
Signature Algorithm:   SHA256-RSA
Key Usage:             Digital Signature
Extended Key Usage:    Server Authentication
SHA-1 Fingerprint:     0F:FA:38:88:60:F8:DA:63:5B:39:7E:33:2D:43:AA:4E:4E:4F:0F:01
SHA-256 Fingerprint:   5F:BB:2E:00:C3:48:A0:B9:9E:F0:17:71:91:82:C1:A2:B9:0C:16:08:CE:86:ED:DB:11:96:09:FE:FE:36:2D:1F
Certificate Authority: No

------------------------------------------------------------
INTERMEDIATE CERTIFICATE #2
------------------------------------------------------------
Subject:               CN=WR2,O=Google Trust Services,C=US
Issuer:                CN=GTS Root R1,O=Google Trust Services LLC,C=US
Serial Number:         170058220837755766831192027518741805976
Version:               3
Valid From:            2023-12-13 09:00:00 UTC
Valid Until:           2029-02-20 14:00:00 UTC
Status:                ✅ Valid (1233 days remaining)
Public Key Algorithm:  RSA
Signature Algorithm:   SHA256-RSA
Key Usage:             Digital Signature, Certificate Signing, CRL Signing
Extended Key Usage:    Server Authentication, Client Authentication
SHA-1 Fingerprint:     66:E4:16:12:60:B1:00:FE:E0:DE:28:7A:9A:52:93:B4:C2:22:4A:E6
SHA-256 Fingerprint:   E6:FE:22:BF:45:E4:F0:D3:B8:5C:59:E0:2C:0F:49:54:18:E1:EB:8D:32:10:F7:88:D4:8C:D5:E1:CB:54:7C:D4
Certificate Authority: Yes
Max Path Length:       0

------------------------------------------------------------
INTERMEDIATE CERTIFICATE #3
------------------------------------------------------------
Subject:               CN=GTS Root R1,O=Google Trust Services LLC,C=US
Issuer:                CN=GlobalSign Root CA,OU=Root CA,O=GlobalSign nv-sa,C=BE
Serial Number:         159159747900478145820483398898491642637
Version:               3
Valid From:            2020-06-19 00:00:42 UTC
Valid Until:           2028-01-28 00:00:42 UTC
Status:                ✅ Valid (844 days remaining)
Public Key Algorithm:  RSA
Signature Algorithm:   SHA256-RSA
Key Usage:             Digital Signature, Certificate Signing, CRL Signing
SHA-1 Fingerprint:     08:74:54:87:E8:91:C1:9E:30:78:C1:F2:A0:7E:45:29:50:EF:36:F6
SHA-256 Fingerprint:   3E:E0:27:8D:F7:1F:A3:C1:25:C4:CD:48:7F:01:D7:74:69:4E:6F:C5:7E:0C:D9:4C:24:EF:D7:69:13:39:18:E5
Certificate Authority: Yes
```
## Display certificate files

It looks for `cert.pem` in the current working directory

### Create a project directory
```bash
mkdir print_certs
cd print_certs
```
### Create `main.go`

```go
package main

import (
    "log"

    "github.com/muquit/ssl-tls-verify/pkg/certificate"
)

func main() {
    // parse a certificate file (auto-detects PEM or DER)
    certs, format, err := certificate.ParseFile("cert.pem")
    if err != nil {
        log.Fatalf("Failed to parse certificate: %v", err)
    }

    // display the certificates
    println("Format:", format)
    certificate.Display(certs)
}

```
### Initialize go module

```bash
go mod init print_certs
go mod tidy
```

### Build
```bash
go build .
```

### Run
```bash
./print_certs
```

```bash
============================================================
CERTIFICATE DETAILS
Total Certificates: 1
============================================================

------------------------------------------------------------
CERTIFICATE
------------------------------------------------------------
Subject:               CN=example.com,OU=example.com,O=Fo Bar LLC,L=Foo Bar,ST=Example,C=AU,1.2.840.113549.1.9.1=#0c106a646f65406578616d706c652e636f6d
Issuer:                CN=example.com,OU=example.com,O=Fo Bar LLC,L=Foo Bar,ST=Example,C=AU,1.2.840.113549.1.9.1=#0c106a646f65406578616d706c652e636f6d
Serial Number:         454408322901924448778008546212275614144524718803
Version:               3
Valid From:            2025-09-27 18:16:21 UTC
Valid Until:           2035-09-25 18:16:21 UTC
Status:                ✅ Valid (3641 days remaining)
Public Key Algorithm:  RSA
Signature Algorithm:   SHA256-RSA
SHA-1 Fingerprint:     D7:4B:21:59:30:4C:84:B0:7C:5D:2E:80:15:66:32:88:7D:62:65:BB
SHA-256 Fingerprint:   43:B8:D4:CF:36:CF:3D:B7:6A:12:06:C2:74:C8:8B:A0:D4:CE:8A:AD:7B:6C:A5:1C:C8:6D:2B:DB:D5:A9:13:8F
Certificate Authority: Yes
```

## Parse PEM Certificates from memory
### Create a project directory
```bash
mkdir parse_pem
cd parse_pem
```
### Create `main.go`
```go
package main

import (
    "log"

    "github.com/muquit/ssl-tls-verify/pkg/certificate"
)

func main() {
    pemData := []byte(`-----BEGIN CERTIFICATE-----
MIIGCTCCA/GgAwIBAgIUT5hfuVg4C+vkn1Tt1/jtRDOIEtMwDQYJKoZIhvcNAQEL
BQAwgZMxCzAJBgNVBAYTAkFVMRAwDgYDVQQIDAdFeGFtcGxlMRAwDgYDVQQHDAdG
b28gQmFyMRMwEQYDVQQKDApGbyBCYXIgTExDMRQwEgYDVQQLDAtleGFtcGxlLmNv
bTEUMBIGA1UEAwwLZXhhbXBsZS5jb20xHzAdBgkqhkiG9w0BCQEWEGpkb2VAZXhh
bXBsZS5jb20wHhcNMjUwOTI3MTgxNjIxWhcNMzUwOTI1MTgxNjIxWjCBkzELMAkG
A1UEBhMCQVUxEDAOBgNVBAgMB0V4YW1wbGUxEDAOBgNVBAcMB0ZvbyBCYXIxEzAR
BgNVBAoMCkZvIEJhciBMTEMxFDASBgNVBAsMC2V4YW1wbGUuY29tMRQwEgYDVQQD
DAtleGFtcGxlLmNvbTEfMB0GCSqGSIb3DQEJARYQamRvZUBleGFtcGxlLmNvbTCC
AiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBALb+lLe9HXXJGq28ELlJk1L0
r1R6mziR07ogfChm4te3WdCAwwvvUjaYnnEb2Lr6VHR3eupS68i8NF/0sclJhEVz
g5a4H3MCLVraX6j65afmQJRQRB8FunbBWVFwONEbpOK5kHK2FEhINspGmU1gYzKK
aX7SpqoXGWyT3S6TugL6Y8ZZFVsDiN2zesAlcjbgFWOuEG62BGPmWh+FqyUAOpix
a/f6ZcmjBBJ8TxEj9WD3syaMUjhSnwojmVH0yNOqbcujxJrK0sBj0M641Qcif33m
+za5arkGONc4VKUhnxaklTvF6UjJ250Fz3FffdlZMbdufRJDJk9H+9rofJfIjUbj
GjWdQu0QaZP/mVc4esDqGowq+siZ8fB3vbbEsNOxA0/h+oyhQ+kStnZgwcoDRXjs
0Pow2lBslI6FDoKDoLNJauR3PgqjBFTbNHblGMgw0IP0fv87EYPtZNz4+T6HnzN4
Lj5Mt0C+VkAd9Zfew/YNNAsTE648l6PHVG/bCiFDL4rgbI8LW61khESPmZ9pmUWC
Ad7b0o/AYK3uhJhSZJ5gv/5SatRDtBm0BvCN+fdXBMTeoWcsWR+lTL2nS2CXujH4
xFC4yNM4Kj5Wcp0rJUOwTcPl7mqxVVabWXVR2Gp9HcM4sciv26WdPVf/ZsU8TNh0
DqFDPX0sShkJIFWi8EAnAgMBAAGjUzBRMB0GA1UdDgQWBBRWZPxLUgTN2+31rMIl
3pIPLyol0DAfBgNVHSMEGDAWgBRWZPxLUgTN2+31rMIl3pIPLyol0DAPBgNVHRMB
Af8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4ICAQBILynM9khfiR3uTpnchaYXkDJ+
rn9y47s710c0yiojOr4F9uedDwNayfK8INLjn/IqE1oVVvpd/XzwIC0VMdRQdnPi
tc1BDLccHkz3IU8gCpn6wG6LeIOBAG7E9mY8laDubdDm+wZcs2RpGBm0EfczcPyE
kbRQDXKEApinmRKr22hZ/jpjjxCmgVKvkRkYzfUVoTfc11rOZf3uKuIep96B5/yS
xVRilaa1uhEYcAR/FE8KuFKJ+pdFLrgZ1SwYBHwfZJYB0ZFYIGrMxtfHkLCmqNev
o0cq8JIqLmgX1FBxaYYEakLDvSh8Xe14w5s2/rhFP+DFduHIxWo9U7XZlv5u0LAo
bDLjUJUvCtNdyhHf3JUWSv6nPNtKSLdjeezvz6ru4uiMPwVS6B96/iWDzgRg6rCV
LADwl63zYKsludU4ba//QGr88W5QzxCbEToyY5GOG0IJj9S1MmmllIgm7ZOq0t6v
9AHzTdyBYkTpMm4/8TBaSF6F3qQ39fZMk+qPAqpsoANf4/TCWgFM7pZPp8r9mvJE
K2Ob64RNqz0BAth6Be5dLV3jFA+IL3wd6YwdGQlNHO96WM0eaXx0xGJXX7sGEXX4
lxozlOeipnPdKJkeH8VeQlIq8I1rpZnFO4A85Rr5F1dPRa4/cU3W93Ixc4RYEAc/
yr8lqoaAcOSOX+Pt0w==
-----END CERTIFICATE-----`)

    certs, err := certificate.ParsePEMCertificates(pemData)
    if err != nil {
        log.Fatalf("Failed to parse PEM: %v", err)
    }

    certificate.Display(certs)
}
```
### Initialize go module

```bash
go mod init parse_pem
go mod tidy
```

### Build
```bash
go build .
```

### Run
```bash
./parse_pem
```

```bash
============================================================
CERTIFICATE DETAILS
Total Certificates: 1
============================================================

------------------------------------------------------------
CERTIFICATE
------------------------------------------------------------
Subject:               CN=example.com,OU=example.com,O=Fo Bar LLC,L=Foo Bar,ST=Example,C=AU,1.2.840.113549.1.9.1=#0c106a646f65406578616d706c652e636f6d
Issuer:                CN=example.com,OU=example.com,O=Fo Bar LLC,L=Foo Bar,ST=Example,C=AU,1.2.840.113549.1.9.1=#0c106a646f65406578616d706c652e636f6d
Serial Number:         454408322901924448778008546212275614144524718803
Version:               3
Valid From:            2025-09-27 18:16:21 UTC
Valid Until:           2035-09-25 18:16:21 UTC
Status:                ✅ Valid (3641 days remaining)
Public Key Algorithm:  RSA
Signature Algorithm:   SHA256-RSA
SHA-1 Fingerprint:     D7:4B:21:59:30:4C:84:B0:7C:5D:2E:80:15:66:32:88:7D:62:65:BB
SHA-256 Fingerprint:   43:B8:D4:CF:36:CF:3D:B7:6A:12:06:C2:74:C8:8B:A0:D4:CE:8A:AD:7B:6C:A5:1C:C8:6D:2B:DB:D5:A9:13:8F
Certificate Authority: Yes
```
## Save certs to different formats
```bash
mkdir save_certs
cd save_certs
```
### Create `main.go`

```go
package main

import (
    "log"
    "time"

    "github.com/muquit/ssl-tls-verify/pkg/certificate"
    "github.com/muquit/ssl-tls-verify/pkg/connection"
)

func main() {
    // Get certificates from a connection
    config := connection.TLSConfig{
        Host:           "example.com",
        Port:           "443",
        SkipVerify:     false,
        ConnectTimeout: 10 * time.Second,
    }

    certs, _, err := connection.TestDirectTLS(config)
    if err != nil {
        log.Fatalf("Connection failed: %v", err)
    }

    // Save as PEM
    err = certificate.Save(certs, certificate.SaveConfig{
        Format:     "pem",
        OutputFile: "certificates.pem",
    })
    if err != nil {
        log.Fatalf("Failed to save PEM: %v", err)
    }

    // Save as DER (leaf certificate only)
    err = certificate.Save(certs, certificate.SaveConfig{
        Format:     "der",
        OutputFile: "certificate.der",
    })
    if err != nil {
        log.Fatalf("Failed to save DER: %v", err)
    }

    // Save as human-readable text
    err = certificate.Save(certs, certificate.SaveConfig{
        Format:     "text",
        OutputFile: "certificate-details.txt",
    })
    if err != nil {
        log.Fatalf("Failed to save text: %v", err)
    }
}
```
### Initialize go module

```bash
go mod init save_certs
go mod tidy
```

### Build
```bash
go build .
```

### Run
```bash
./save_certs
```

```bash
✅ Successfully saved 2 certificate(s) in PEM format to certificates.pem

✅ Successfully saved leaf certificate in DER format to certificate.der
   Note: Saved leaf certificate only. 1 intermediate certificate(s) were not saved.
   Use PEM format (--save pem) to save the complete certificate chain in one file.

✅ Successfully saved 2 certificate(s) in TEXT format to certificate-details.txt
```

## Generate certificate fingerprints
### Create a project directory
```bash
mkdir gen_fingerprint
cd gen_fingerprint
```
### Create `main.go`
```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/muquit/ssl-tls-verify/pkg/certificate"
    "github.com/muquit/ssl-tls-verify/pkg/connection"
)

func main() {
    config := connection.TLSConfig{
        Host:           "google.com",
        Port:           "443",
        SkipVerify:     false,
        ConnectTimeout: 10 * time.Second,
    }

    certs, _, err := connection.TestDirectTLS(config)
    if err != nil {
        log.Fatalf("Connection failed: %v", err)
    }

    // Get fingerprints for the leaf certificate
    leafCert := certs[0]

    sha1FP := certificate.FormatFingerprint(leafCert.Raw, "sha1")
    sha256FP := certificate.FormatFingerprint(leafCert.Raw, "sha256")

    fmt.Printf("SHA-1 Fingerprint:   %s\n", sha1FP)
    fmt.Printf("SHA-256 Fingerprint: %s\n", sha256FP)
}
```
### Initialize go module

```bash
go mod init gen_fingerprint
go mod tidy
```

### Build
```bash
go build .
```

### Run
```bash
./gen_fingerprint
```
```bash
SHA-1 Fingerprint:   33:09:D4:D3:61:83:44:73:CF:04:5A:44:53:2E:B5:36:64:BB:FC:7F
SHA-256 Fingerprint: 86:F0:16:7D:8B:24:BC:17:6C:2B:06:E6:05:A6:33:43:10:A0:61:A6:BA:81:FB:58:ED:0A:DC:0B:AE:5D:08:C6
```
etc.

# API Reference

## `pkg/certificate`

### `ParseFile(filename string) ([]*x509.Certificate, string, error)`
Parse a certificate file (auto-detects PEM or DER format).

**Returns:** certificates, format ("PEM" or "DER"), error

### `ParsePEMCertificates(data []byte) ([]*x509.Certificate, error)`
Parse PEM-encoded certificates from raw bytes.

### `Display(certs []*x509.Certificate)`
Display certificates in human-readable format to stdout.

### `PrintDetails(cert *x509.Certificate)`
Print detailed certificate information to stdout.

### `PrintDetailsToWriter(cert *x509.Certificate, output io.Writer)`
Print certificate details to a specific writer.

### `FormatFingerprint(certRaw []byte, algorithm string) string`
Generate certificate fingerprint. Algorithms: "sha1", "sha256"

### `Save(certs []*x509.Certificate, config SaveConfig) error`
Save certificates in specified format (pem, der, text).

### `SavePEM(certs []*x509.Certificate, output io.Writer) error`
Save certificates in PEM format.

### `SaveDER(certs []*x509.Certificate, output io.Writer) error`
Save certificate in DER format (leaf only).

### `SaveText(certs []*x509.Certificate, output io.Writer) error`
Save certificates in human-readable text format.

## `pkg/connection`

### `TestDirectTLS(config TLSConfig) ([]*x509.Certificate, tls.ConnectionState, error)`
Test a direct TLS connection (HTTPS, IMAPS, etc.).

**TLSConfig fields:**
- `Host string` - Target hostname
- `Port string` - Target port
- `SkipVerify bool` - Skip certificate verification
- `ConnectTimeout time.Duration` - Connection timeout

### `TestStartTLS(config StartTLSConfig) ([]*x509.Certificate, tls.ConnectionState, error)`
Test a StartTLS connection (SMTP, IMAP, POP3, etc.).

**StartTLSConfig fields:**
- `Host string` - Target hostname
- `Port string` - Target port
- `SkipVerify bool` - Skip certificate verification
- `ConnectTimeout time.Duration` - Connection timeout

### `PrintTLSInfo(state tls.ConnectionState)`
Print detailed TLS connection and certificate information.

### `PrintVerifyFailure(host string, err error) bool`
If `err` is a certificate verification failure, print the certificate chain
presented by the server and the result of re-verifying it against the system
roots to `stderr`, and return `true`. Returns `false` and prints nothing for
any other error. Call it with the error returned by `TestDirectTLS` or
`TestStartTLS`.

## `pkg/utils`

### `MaskIPAddresses(text string) string`
Mask IPv4 and IPv6 addresses in text (requires `MAILSEND_MASK_IP` env var).

### `MaskIPSlice(addresses []string) []string`
Mask IP addresses in a string slice (requires `MAILSEND_MASK_IP` env var).

---

## Best Practices

1. **Always set timeouts** when making connections to prevent hanging
2. **Handle errors properly** - connection failures are common
3. **Use `SkipVerify: false`** in production unless you have a specific reason
4. **Check certificate expiration** dates if you're monitoring certificates
5. **Use PEM format** for saving full certificate chains
6. **Enable IP masking** when sharing logs or screenshots

## Error Handling

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

# License

The license is MIT. See [LICENSE](LICENSE) file for details.

# Authors
* Developed with [Claude AI Sonnet 4.5](https://claude.ai), working under my guidance and instructions.


---
<sub>TOC/glossary expansion by https://github.com/muquit/markdown-toc-go v1.0.6 on Oct-31-2026</sub>
