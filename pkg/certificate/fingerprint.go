package certificate

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// FormatFingerprint creates a formatted fingerprint string from certificate raw bytes.
// Supported algorithms: "sha1", "sha256"
// Returns fingerprint in format: XX:XX:XX:XX...
func FormatFingerprint(certRaw []byte, algorithm string) string {
	var hash []byte
	switch algorithm {
	case "sha1":
		h := sha1.Sum(certRaw)
		hash = h[:]
	case "sha256":
		h := sha256.Sum256(certRaw)
		hash = h[:]
	default:
		return "Unknown algorithm"
	}

	hexStr := hex.EncodeToString(hash)
	// Format as XX:XX:XX...
	var formatted []string
	for i := 0; i < len(hexStr); i += 2 {
		formatted = append(formatted, strings.ToUpper(hexStr[i:i+2]))
	}
	return strings.Join(formatted, ":")
}