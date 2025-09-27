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

