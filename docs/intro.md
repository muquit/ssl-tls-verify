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
the certificate parsing and validation is handled by @GO@'s 
`crypto/tls` and `crypto/x509` packages.

If you have any question, request or suggestion, please enter it in the
[Issues](https://github.com/muquit/ssl-tls-verify/issues) with appropriate label.
