## Table Of Contents
  - [Introduction](#introduction)
  - [Background](#background)
  - [Synopsis](#synopsis)
  - [Latest Version (v1.0.1)](#latest-version-v101)
  - [Installation](#installation)
    - [Build from source](#build-from-source)
    - [Install directly](#install-directly)
  - [Examples](#examples)
    - [Sample output](#sample-output)
      - [google.com HTTPS port 443](#googlecom-https-port-443)
      - [Server using self signed certificate](#server-using-self-signed-certificate)
      - [Skip Verification](#skip-verification)
      - [smtp.gmail.com with StartTLS](#smtpgmailcom-with-starttls)
  - [License](#license)
  - [Authors](#authors)

## Introduction
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
the certificate parsing and validation is handled by [Go](https://go.dev/)'s 
`crypto/tls` and `crypto/x509` packages.

If you have any question, request or suggestion, please enter it in the
[Issues](https://github.com/muquit/ssl-tls-verify/issues) with appropriate label.

## Background

There was a lingering bug in [mailsend-go](https://github.com/muquit/mailsend-go) that `-verifyCert` does not work.
I was not getting time to look at the code and figure out the issue, so I asked [Claude AI Sonnet 4](https://claude.ai) to
write a simple CLI tool to verify certificates with optional verification. Hence, `ssl-tls-verify` was born.
After creating this tool, it became clear to me that I had missed `ServerName` in `tls.Config` -
a critical field needed for proper certificate hostname validation.
I also love this tool, it's much nicer than openssl with its human-readable output,
clear certificate chain display, and intuitive formatting. What started as a debugging
utility has become my go-to tool for SSL/TLS certificate inspection. Hope you
find that tool useful as well.

## Synopsis
```
ssl-tls-verify v1.0.1
https://github.com/muquit/ssl-tls-verify
A cross-platform command line tool to test and verify SSL/TLS connections
  -host string
    	Target host (default "google.com")
  -port string
    	Target port (443 for HTTPS, 587 for SMTP StartTLS) (default "443")
  -skip-verify
    	Skip certificate verification
  -starttls
    	Use StartTLS (for SMTP-like protocols)
  -timeout duration
    	Connection timeout (default 10s)
  -version
    	Show version information and exit
```

## Latest Version (v1.0.1)
The current version is v1.0.1
Please look at [ChangeLog](ChangeLog.md) for what has changed in the current version.

## Installation

Please download pre-compiled binaries from [Releases](https://github.com/muquit/ssl-tls-verify/releases) page

### Build from source
```bash
go build -o ssl-tls-verify main.go
# or
make build
```

### Install directly
```bash
go install github.com/muquit/ssl-tls-verify@latest
```

## Examples

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


### Sample output
#### google.com HTTPS port 443
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
Issuer:                CN=WR2,O=Google Trust Services,C=US
Serial Number:         133813938596851887989668938256071211179
Version:               3
Valid From:            2025-09-08 08:34:53 UTC
Valid Until:           2025-12-01 08:34:52 UTC
Status:                ✅ Valid (64 days remaining)
DNS Names:             *.google.com, *.appengine.google.com, *.bdn.dev, *.origin-test.bdn.dev, *.cloud.google.com, *.crowdsource.google.com, *.datacompute.google.com, *.google.ca, *.google.cl, *.google.co.in, *.google.co.jp, *.google.co.uk, *.google.com.ar, *.google.com.au, *.google.com.br, *.google.com.co, *.google.com.mx, *.google.com.tr, *.google.com.vn, *.google.de, *.google.es, *.google.fr, *.google.hu, *.google.it, *.google.nl, *.google.pl, *.google.pt, *.googleapis.cn, *.googlevideo.com, *.gstatic.cn, *.gstatic-cn.com, googlecnapps.cn, *.googlecnapps.cn, googleapps-cn.com, *.googleapps-cn.com, gkecnapps.cn, *.gkecnapps.cn, googledownloads.cn, *.googledownloads.cn, recaptcha.net.cn, *.recaptcha.net.cn, recaptcha-cn.net, *.recaptcha-cn.net, widevine.cn, *.widevine.cn, ampproject.org.cn, *.ampproject.org.cn, ampproject.net.cn, *.ampproject.net.cn, google-analytics-cn.com, *.google-analytics-cn.com, googleadservices-cn.com, *.googleadservices-cn.com, googlevads-cn.com, *.googlevads-cn.com, googleapis-cn.com, *.googleapis-cn.com, googleoptimize-cn.com, *.googleoptimize-cn.com, doubleclick-cn.net, *.doubleclick-cn.net, *.fls.doubleclick-cn.net, *.g.doubleclick-cn.net, doubleclick.cn, *.doubleclick.cn, *.fls.doubleclick.cn, *.g.doubleclick.cn, dartsearch-cn.net, *.dartsearch-cn.net, googletraveladservices-cn.com, *.googletraveladservices-cn.com, googletagservices-cn.com, *.googletagservices-cn.com, googletagmanager-cn.com, *.googletagmanager-cn.com, googlesyndication-cn.com, *.googlesyndication-cn.com, *.safeframe.googlesyndication-cn.com, app-measurement-cn.com, *.app-measurement-cn.com, gvt1-cn.com, *.gvt1-cn.com, gvt2-cn.com, *.gvt2-cn.com, 2mdn-cn.net, *.2mdn-cn.net, googleflights-cn.net, *.googleflights-cn.net, admob-cn.com, *.admob-cn.com, *.gemini.cloud.google.com, googlesandbox-cn.com, *.googlesandbox-cn.com, *.safenup.googlesandbox-cn.com, *.gstatic.com, *.metric.gstatic.com, *.gvt1.com, *.gcpcdn.gvt1.com, *.gvt2.com, *.gcp.gvt2.com, *.url.google.com, *.youtube-nocookie.com, *.ytimg.com, ai.android, android.com, *.android.com, *.flash.android.com, g.cn, *.g.cn, g.co, *.g.co, goo.gl, www.goo.gl, google-analytics.com, *.google-analytics.com, google.com, googlecommerce.com, *.googlecommerce.com, ggpht.cn, *.ggpht.cn, urchin.com, *.urchin.com, youtu.be, youtube.com, *.youtube.com, music.youtube.com, *.music.youtube.com, youtubeeducation.com, *.youtubeeducation.com, youtubekids.com, *.youtubekids.com, yt.be, *.yt.be, android.clients.google.com, *.android.google.cn, *.chrome.google.cn, *.developers.google.cn, *.aistudio.google.com
Public Key Algorithm:  ECDSA
Signature Algorithm:   SHA256-RSA
Key Usage:             Digital Signature
Extended Key Usage:    Server Authentication
SHA-1 Fingerprint:     13:4A:0B:81:A8:A4:37:A9:D7:31:CB:DD:A6:76:53:21:8A:1B:2E:0C
SHA-256 Fingerprint:   2D:8F:A1:B5:9A:60:F4:14:AD:1C:29:44:92:C7:8B:AF:4F:27:CD:EE:15:0F:A8:E4:E8:11:CD:41:8C:75:45:4B
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
Status:                ✅ Valid (1241 days remaining)
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
Status:                ✅ Valid (852 days remaining)
Public Key Algorithm:  RSA
Signature Algorithm:   SHA256-RSA
Key Usage:             Digital Signature, Certificate Signing, CRL Signing
SHA-1 Fingerprint:     08:74:54:87:E8:91:C1:9E:30:78:C1:F2:A0:7E:45:29:50:EF:36:F6
SHA-256 Fingerprint:   3E:E0:27:8D:F7:1F:A3:C1:25:C4:CD:48:7F:01:D7:74:69:4E:6F:C5:7E:0C:D9:4C:24:EF:D7:69:13:39:18:E5
Certificate Authority: Yes
Direct TLS connection successful!
```
#### Server using self signed certificate
```bash
# Server using self-signed certificate
./ssl-tls-verify -host 127.0.0.1 -port 8881
```
```bash
Connecting to 127.0.0.1:8881
Verify Certificate: Yes
Use StartTLS: No

2025/09/27 16:37:52 Direct TLS connection failed: failed to establish TLS connection: tls: failed to verify certificate: x509: “example.com” certificate is not standards compliant
```
#### Skip Verification
```bash
# Skip verification
./ssl-tls-verify -host 127.0.0.1 -port 8881 -skip-verify
```
```bash
Connecting to 127.0.0.1:8881
Verify Certificate: No
Use StartTLS: No


============================================================
TLS CONNECTION INFORMATION
============================================================
TLS Version:           TLS 1.3 (0x0304)
Cipher Suite:          0x1301
Server Certificates:   1
Handshake Complete:    true
Did Resume:            false
Negotiated Protocol:   

------------------------------------------------------------
SERVER CERTIFICATE #1 (Leaf Certificate)
------------------------------------------------------------
Subject:               CN=example.com,OU=example.com,O=Fo Bar LLC,L=Foo Bar,ST=Example,C=AU,1.2.840.113549.1.9.1=#0c106a646f65406578616d706c652e636f6d
Issuer:                CN=example.com,OU=example.com,O=Fo Bar LLC,L=Foo Bar,ST=Example,C=AU,1.2.840.113549.1.9.1=#0c106a646f65406578616d706c652e636f6d
Serial Number:         454408322901924448778008546212275614144524718803
Version:               3
Valid From:            2025-09-27 18:16:21 UTC
Valid Until:           2035-09-25 18:16:21 UTC
Status:                ✅ Valid (3649 days remaining)
Public Key Algorithm:  RSA
Signature Algorithm:   SHA256-RSA
SHA-1 Fingerprint:     D7:4B:21:59:30:4C:84:B0:7C:5D:2E:80:15:66:32:88:7D:62:65:BB
SHA-256 Fingerprint:   43:B8:D4:CF:36:CF:3D:B7:6A:12:06:C2:74:C8:8B:A0:D4:CE:8A:AD:7B:6C:A5:1C:C8:6D:2B:DB:D5:A9:13:8F
Certificate Authority: Yes
Direct TLS connection successful!
```
#### smtp.gmail.com with StartTLS
```bash
./ssl-tls-verify --host smtp.gmail.com --port 587 --starttls
```
```bash
Connecting to smtp.gmail.com:587
Verify Certificate: Yes
Use StartTLS: Yes

Plain connection established, attempting StartTLS...
Server greeting: 220 smtp.gmail.com ESMTP af79cd13be357-85c2737869esm489698085a.11 - gsmtp
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
Issuer:                CN=WR2,O=Google Trust Services,C=US
Serial Number:         14461562026188826353951632455228095006
Version:               3
Valid From:            2025-09-08 08:36:45 UTC
Valid Until:           2025-12-01 08:36:44 UTC
Status:                ✅ Valid (64 days remaining)
DNS Names:             smtp.gmail.com
Public Key Algorithm:  ECDSA
Signature Algorithm:   SHA256-RSA
Key Usage:             Digital Signature
Extended Key Usage:    Server Authentication
SHA-1 Fingerprint:     28:88:45:90:10:20:88:BA:87:2E:0E:7C:3A:12:D6:35:EC:26:AE:90
SHA-256 Fingerprint:   6F:F8:E2:F5:D4:AE:5A:FF:92:4A:5F:AC:88:80:14:3A:30:33:7A:CF:EE:33:94:82:EF:2A:93:47:80:E4:18:EF
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
Status:                ✅ Valid (1241 days remaining)
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
Status:                ✅ Valid (852 days remaining)
Public Key Algorithm:  RSA
Signature Algorithm:   SHA256-RSA
Key Usage:             Digital Signature, Certificate Signing, CRL Signing
SHA-1 Fingerprint:     08:74:54:87:E8:91:C1:9E:30:78:C1:F2:A0:7E:45:29:50:EF:36:F6
SHA-256 Fingerprint:   3E:E0:27:8D:F7:1F:A3:C1:25:C4:CD:48:7F:01:D7:74:69:4E:6F:C5:7E:0C:D9:4C:24:EF:D7:69:13:39:18:E5
Certificate Authority: Yes
StartTLS connection successful!
```

## License

This project is released under the MIT License. See LICENSE file for details.

## Authors
* Developed with [Claude AI Sonnet 4](https://claude.ai), working under my guidance and instructions.


---
<sub>TOC is created by https://github.com/muquit/markdown-toc-go on Sep-27-2025</sub>
