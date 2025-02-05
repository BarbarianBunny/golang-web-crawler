package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		expected  string
		expectErr bool
	}{
		{
			name:      "remove https and slash",
			inputURL:  "https://blog.boot.dev/path/",
			expected:  "blog.boot.dev/path",
			expectErr: false,
		}, {
			name:      "remove https",
			inputURL:  "https://blog.boot.dev/path",
			expected:  "blog.boot.dev/path",
			expectErr: false,
		}, {
			name:      "remove http and slash",
			inputURL:  "http://blog.boot.dev/path/",
			expected:  "blog.boot.dev/path",
			expectErr: false,
		}, {
			name:      "remove http",
			inputURL:  "http://blog.boot.dev/path",
			expected:  "blog.boot.dev/path",
			expectErr: false,
		}, {
			name:      "remove slash",
			inputURL:  "blog.boot.dev/path/",
			expected:  "blog.boot.dev/path",
			expectErr: false,
		}, {
			name:      "leave unchanged",
			inputURL:  "blog.boot.dev/path",
			expected:  "blog.boot.dev/path",
			expectErr: false,
		}, {
			name:      "remove host slash",
			inputURL:  "blog.boot.dev/",
			expected:  "blog.boot.dev",
			expectErr: false,
		}, {
			name:      "leave host unchanged",
			inputURL:  "blog.boot.dev",
			expected:  "blog.boot.dev",
			expectErr: false,
		}, {
			name:      "error no period",
			inputURL:  "dev/path",
			expected:  "dev/path",
			expectErr: true,
		}, {
			name:      "error no domain",
			inputURL:  "http://",
			expected:  "http://",
			expectErr: true,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err == nil && tc.expectErr {
				t.Errorf("Test %v - '%s' FAIL: expected error: nil", i, tc.name)
				return
			}
			if err != nil && !tc.expectErr {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}
