package utils

import (
	"os"
	"regexp"
)

// MaskIPAddresses masks IPv4 and IPv6 addresses in text for privacy.
// Masking only occurs when the MAILSEND_MASK_IP environment variable is set.
// IPv4 addresses are replaced with "xxx.xxx.xxx.xxx"
// IPv6 addresses are replaced with "xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx"
func MaskIPAddresses(text string) string {
	if os.Getenv("MAILSEND_MASK_IP") == "" {
		return text
	}

	// IPv4 pattern: matches xxx.xxx.xxx.xxx format
	ipv4Pattern := regexp.MustCompile(`\b(\d{1,3}\.){3}\d{1,3}\b`)

	// IPv6 pattern: matches various IPv6 formats
	ipv6Pattern := regexp.MustCompile(`\b([0-9a-fA-F]{0,4}:){2,7}[0-9a-fA-F]{0,4}\b`)

	// Replace IPv4 addresses
	maskedText := ipv4Pattern.ReplaceAllString(text, "xxx.xxx.xxx.xxx")

	// Replace IPv6 addresses
	maskedText = ipv6Pattern.ReplaceAllString(maskedText, "xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx")

	return maskedText
}

// MaskIPSlice masks IP addresses in a string slice for privacy.
// Masking only occurs when the MAILSEND_MASK_IP environment variable is set.
// Each string in the slice is processed through MaskIPAddresses.
func MaskIPSlice(addresses []string) []string {
	if os.Getenv("MAILSEND_MASK_IP") == "" {
		return addresses
	}

	masked := make([]string, len(addresses))
	for i, addr := range addresses {
		masked[i] = MaskIPAddresses(addr)
	}
	return masked
}