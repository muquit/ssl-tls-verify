# Table Of Contents
- [v1.0.4](#v104)
- [v1.0.3](#v103)
- [v1.0.2](#v102)
- [v1.0.1](#v101)

# v1.0.4
> [!NOTE]
> In October 2026, eight intermediate CA certificates operated by
> eMudhra under its emSign PKI were found to have private keys that
> could be derived by anyone.
> go 1.26.10 and 1.27.3 now reject certificate chains containing
> these intermediate CA keys.
> See [golang/go#82345](https://github.com/golang/go/issues/82345).

- Rebuilt the binaries with go 1.27.3 to pick up the crypto/x509 fix
for [CVE-2026-111943](https://nvd.nist.gov/vuln/detail/cve-2026-111943).
Certificate chains containing the blocked intermediate CA keys now will
be blocked.

- `go.mod` updated to require go 1.27.3 or later for building from source.

- Fixed a bug where the host was silently ignored if it was given as an
argument, and `google.com` was tested instead. There is no default host
now.
    - The host can be given as an argument: `ssl-tls-verify example.com`.
      Flags must come before it.
    - Unexpected arguments are reported as an error.
    - If neither `--host` nor `--dump` is specified, an error and the usage
      are printed and the program exits with code 2.

- Added `--example` to print usage examples. The examples are removed from
`--help`. An example of checking `dv-g3b-emudhra.we-have-a-problem.com`,
which should fail verification, is added.

- If certificate verification fails, the certificate chain presented by the
server (subject, issuer, validity and SHA-256 fingerprint of each
certificate) and the result of re-verifying it against the system roots are
printed to `stderr` before the error. Previously only the error was shown.

- Removed the name, version and URL lines from the top of `--help`, they are
shown at the bottom.

(Oct-31-2026)

# v1.0.3
- No code logic changed, code is re-factored from one giant main.go, so that
the code be re-used in other projects.

- Add examples on how to use `ssl-tls-verify` as a libary
- Added API reference

(Oct-05-2025)

# v1.0.2
- Added the following new options:
    - `--save format` to save certificates from connection. The supported format is pem, der or text.
    - `--output file` to save certificates from connection to the file. The default is `stdout`
    - `--dump cert_file` to read and display certificate file from disk (PEM, DER or CERT format)

(Oct-03-2025)

# v1.0.1

- Initial release

(Sep-27-2025)
