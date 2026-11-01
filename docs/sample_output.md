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
