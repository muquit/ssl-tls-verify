## Examples
### Commands

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
