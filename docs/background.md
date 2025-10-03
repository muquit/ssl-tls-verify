## Background

There was a lingering bug in @MAILSEND_GO@ that `-verifyCert` was broken (Bug
        #71)
I was not getting time to look at the code and figure out the issue, so I asked @CLAUDE@ to
write a simple CLI tool in @GO@ to verify certificates from connection. Hence, `ssl-tls-verify` was born.
After creating this tool, I realized I had missed `ServerName` in `tls.Config`.
I also love this tool, it's much nicer than openssl with its human-readable output,
clear certificate chain display, and intuitive formatting. What started as a debugging
utility has become my go-to tool for SSL/TLS certificate inspection. 

Hope you find the tool useful as well.
