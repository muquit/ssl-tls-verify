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
Serial Number:         166700031870418575880302779524424518776
Version:               3
Valid From:            2025-09-15 08:34:18 UTC
Valid Until:           2025-12-08 08:34:17 UTC
Status:                ✅ Valid (66 days remaining)
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
Status:                ✅ Valid (1236 days remaining)
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
Status:                ✅ Valid (846 days remaining)
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

2025/10/02 22:17:17 Direct TLS connection failed: failed to establish TLS connection: tls: failed to verify certificate: x509: “example.com” certificate is not standards compliant
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
Status:                ✅ Valid (3644 days remaining)
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
Server greeting: 220 smtp.gmail.com ESMTP 6a1803df08f44-878bd785963sm28393576d6.35 - gsmtp
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
Serial Number:         198559767591087662403673263944902778968
Version:               3
Valid From:            2025-09-15 08:36:10 UTC
Valid Until:           2025-12-08 08:36:09 UTC
Status:                ✅ Valid (66 days remaining)
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
Status:                ✅ Valid (1236 days remaining)
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
Status:                ✅ Valid (846 days remaining)
Public Key Algorithm:  RSA
Signature Algorithm:   SHA256-RSA
Key Usage:             Digital Signature, Certificate Signing, CRL Signing
SHA-1 Fingerprint:     08:74:54:87:E8:91:C1:9E:30:78:C1:F2:A0:7E:45:29:50:EF:36:F6
SHA-256 Fingerprint:   3E:E0:27:8D:F7:1F:A3:C1:25:C4:CD:48:7F:01:D7:74:69:4E:6F:C5:7E:0C:D9:4C:24:EF:D7:69:13:39:18:E5
Certificate Authority: Yes
StartTLS connection successful!
```
