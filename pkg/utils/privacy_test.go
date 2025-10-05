package utils

import (
	"os"
	"testing"
)

func TestMaskIPAddresses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maskEnv  string
		expected string
	}{
		{
			name:     "No masking when env not set",
			input:    "Server at 192.168.1.1 is ready",
			maskEnv:  "",
			expected: "Server at 192.168.1.1 is ready",
		},
		{
			name:     "Mask IPv4 address",
			input:    "Server at 192.168.1.1 is ready",
			maskEnv:  "1",
			expected: "Server at xxx.xxx.xxx.xxx is ready",
		},
		{
			name:     "Mask multiple IPv4 addresses",
			input:    "From 10.0.0.1 to 192.168.1.100",
			maskEnv:  "1",
			expected: "From xxx.xxx.xxx.xxx to xxx.xxx.xxx.xxx",
		},
		{
			name:     "Mask IPv6 address",
			input:    "Server at 2001:dead:beef:cafe:0000:8a2e:0370:7334",
			maskEnv:  "1",
			expected: "Server at xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx",
		},
		{
			name:     "Mask compressed IPv6",
			input:    "Server at 2001:db8::1",
			maskEnv:  "1",
			expected: "Server at xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx",
		},
		{
			name:     "No IP addresses in text",
			input:    "Hello World",
			maskEnv:  "1",
			expected: "Hello World",
		},
		{
			name:     "Mixed IPv4 and IPv6",
			input:    "IPv4: 192.168.1.1 and IPv6: 2001:db8::1",
			maskEnv:  "1",
			expected: "IPv4: xxx.xxx.xxx.xxx and IPv6: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set or unset environment variable
			if tt.maskEnv != "" {
				os.Setenv("MAILSEND_MASK_IP", tt.maskEnv)
				defer os.Unsetenv("MAILSEND_MASK_IP")
			} else {
				os.Unsetenv("MAILSEND_MASK_IP")
			}

			result := MaskIPAddresses(tt.input)
			if result != tt.expected {
				t.Errorf("MaskIPAddresses() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestMaskIPSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		maskEnv  string
		expected []string
	}{
		{
			name:     "No masking when env not set",
			input:    []string{"192.168.1.1", "10.0.0.1"},
			maskEnv:  "",
			expected: []string{"192.168.1.1", "10.0.0.1"},
		},
		{
			name:     "Mask multiple addresses",
			input:    []string{"192.168.1.1", "10.0.0.1", "2001:db8::1"},
			maskEnv:  "1",
			expected: []string{"xxx.xxx.xxx.xxx", "xxx.xxx.xxx.xxx", "xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx"},
		},
		{
			name:     "Empty slice",
			input:    []string{},
			maskEnv:  "1",
			expected: []string{},
		},
		{
			name:     "Mixed content",
			input:    []string{"Server: 192.168.1.1", "example.com", "Host at 10.0.0.1"},
			maskEnv:  "1",
			expected: []string{"Server: xxx.xxx.xxx.xxx", "example.com", "Host at xxx.xxx.xxx.xxx"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set or unset environment variable
			if tt.maskEnv != "" {
				os.Setenv("MAILSEND_MASK_IP", tt.maskEnv)
				defer os.Unsetenv("MAILSEND_MASK_IP")
			} else {
				os.Unsetenv("MAILSEND_MASK_IP")
			}

			result := MaskIPSlice(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("MaskIPSlice() length = %d, want %d", len(result), len(tt.expected))
				return
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("MaskIPSlice()[%d] = %q, want %q", i, result[i], tt.expected[i])
				}
			}
		})
	}
}
