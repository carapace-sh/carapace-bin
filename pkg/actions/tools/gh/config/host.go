package config

import "strings"

const defaultHostname = "github.com"

// DefaultHost returns the host name of the default GitHub instance.
func DefaultHost() string {
	return defaultHostname
}

// IsEnterprise reports whether a non-normalized host name looks like a GHE instance.
func IsEnterprise(h string) bool {
	return NormalizeHostname(h) != defaultHostname
}

// NormalizeHostname returns the canonical host name of a GitHub instance.
func NormalizeHostname(h string) string {
	hostname := strings.ToLower(h)
	if strings.HasSuffix(hostname, "."+defaultHostname) {
		return defaultHostname
	}
	return hostname
}
