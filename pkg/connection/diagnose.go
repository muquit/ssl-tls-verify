package connection

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"

	"github.com/muquit/ssl-tls-verify/pkg/certificate"
)

// PrintVerifyFailure prints details about the certificate chain the server
// presented when err is a certificate verification failure. It returns false
// (and prints nothing) for any other kind of error.
func PrintVerifyFailure(host string, err error) bool {
	var verr *tls.CertificateVerificationError
	if !errors.As(err, &verr) || len(verr.UnverifiedCertificates) == 0 {
		return false
	}
	certs := verr.UnverifiedCertificates

	out := os.Stderr
	fmt.Fprintf(out, "\nCertificate verification FAILED for %s\n", host)
	fmt.Fprintf(out, "Reason: %v\n", verr.Err)

	fmt.Fprintf(out, "\nChain presented by the server (%d certificate(s)):\n", len(certs))
	for i, c := range certs {
		fmt.Fprintf(out, "\n [%d] %s\n", i, role(i, c))
		fmt.Fprintf(out, "     Subject:    %s\n", c.Subject)
		fmt.Fprintf(out, "     Issuer:     %s\n", c.Issuer)
		fmt.Fprintf(out, "     Valid:      %s to %s\n",
			c.NotBefore.Format("2006-01-02"), c.NotAfter.Format("2006-01-02"))
		fmt.Fprintf(out, "     SHA-256:    %s\n", certificate.FormatFingerprint(c.Raw, "sha256"))
	}

	// Re-run verification ourselves to show the verifier's own verdict.
	intermediates := x509.NewCertPool()
	for _, c := range certs[1:] {
		intermediates.AddCert(c)
	}
	_, rerr := certs[0].Verify(x509.VerifyOptions{
		DNSName:       host,
		Intermediates: intermediates,
	})
	if rerr != nil {
		fmt.Fprintf(out, "\nRe-verification against system roots: %v\n", rerr)
	} else {
		fmt.Fprintf(out, "\nRe-verification against system roots: succeeded (the failure may be context dependent)\n")
	}
	fmt.Fprintln(out)
	return true
}

func role(i int, c *x509.Certificate) string {
	switch {
	case i == 0:
		return "Leaf"
	case c.CheckSignatureFrom(c) == nil && c.Subject.String() == c.Issuer.String():
		return "Root (self-signed)"
	default:
		return "Intermediate"
	}
}
