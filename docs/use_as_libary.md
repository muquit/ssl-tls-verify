## How to use ssl-tls-verify as a libary

`ssl-tls-verify` can be use a a libary from other projects. Some examples are
shown below

### Cert checker

#### Create a project directory
```bash
mkdir cert_checker
cd cert_checker
```
#### Create `main.go`

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
#### Initialize go module

```bash
go mod init cert_checker
go mod tidy
```

#### Build
```bash
go build .
```

#### Run
```bash
./cert_checker
```
```bash
OK: google.com - Valid for 63 days
OK: github.com - Valid for 123 days
OK: stackoverflow.com - Valid for 49 days
```

### TLS checker
#### Create a project directory
```bash
mkdir tls_checker
cd tls_checker
```
#### Create `main.go`

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
#### Initialize go module
```bash
go mod init tls_checker
go mod tidy
```

#### Build
```bash
go build .
```

#### Run
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
### StartTLS checker

#### Create a project directory
```bash
mkdir starttls_checker
cd starttls_checker
```
#### Create `main.go`

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
#### Initialize go module
```bash
go mod init starttls_checker
go mod tidy
```

#### Build
```bash
go build .
```

#### Run

```bash
./starttls_checker
```

```bash
Plain connection established, attempting StartTLS...
Server greeting: 220 smtp.gmail.com ESMTP 6a1803df08f44-878bae60582sm97362196d6.7 - gsmtp
EHLO response: 250-smtp.gmail.com at your service, [2600:4040:7800:5500:d0e5:435d:ef7c:6518]
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
### Display certificate files

It looks for `cert.pem` in the current working directory

#### Create a project directory
```bash
mkdir print_certs
cd print_certs
```
#### Create `main.go`

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
#### Initialize go module

```bash
go mod init print_certs
go mod tidy
```

#### Build
```bash
go build .
```

#### Run
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

### Parse PEM Certificates from memory
#### Create a project directory
```bash
mkdir parse_pem
cd parse_pem
```
#### Create `main.go`
```go
package main

import (
    "log"

    "github.com/muquit/ssl-tls-verify/pkg/certificate"
)

func main() {
    pemData := []byte(`-----BEGIN CERTIFICATE-----
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
yr8lqoaAcOSOX+Pt0w==
-----END CERTIFICATE-----`)

    certs, err := certificate.ParsePEMCertificates(pemData)
    if err != nil {
        log.Fatalf("Failed to parse PEM: %v", err)
    }

    certificate.Display(certs)
}
```
#### Initialize go module

```bash
go mod init parse_pem
go mod tidy
```

#### Build
```bash
go build .
```

#### Run
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
### Save certs to different formats
```bash
mkdir save_certs
cd save_certs
```
#### Create `main.go`

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
#### Initialize go module

```bash
go mod init save_certs
go mod tidy
```

#### Build
```bash
go build .
```

#### Run
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

### Generate certificate fingerprints
#### Create a project directory
```bash
mkdir gen_fingerprint
cd gen_fingerprint
```
#### Create `main.go`
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
#### Initialize go module

```bash
go mod init gen_fingerprint
go mod tidy
```

#### Build
```bash
go build .
```

#### Run
```bash
./gen_fingerprint
```
```bash
SHA-1 Fingerprint:   33:09:D4:D3:61:83:44:73:CF:04:5A:44:53:2E:B5:36:64:BB:FC:7F
SHA-256 Fingerprint: 86:F0:16:7D:8B:24:BC:17:6C:2B:06:E6:05:A6:33:43:10:A0:61:A6:BA:81:FB:58:ED:0A:DC:0B:AE:5D:08:C6
```
etc.
