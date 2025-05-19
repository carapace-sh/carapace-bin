package ghinstance

import (
	"errors"
	"fmt"
	"strings"
)

const defaultHostname = "github.com"

// Default returns the host name of the default GitHub instance
func Default() string {
	return defaultHostname
}

// IsEnterprise reports whether a non-normalized host name looks like a GHE instance
func IsEnterprise(h string) bool {
	return NormalizeHostname(h) != defaultHostname
}

// NormalizeHostname returns the canonical host name of a GitHub instance
func NormalizeHostname(h string) string {
	hostname := strings.ToLower(h)
	if strings.HasSuffix(hostname, "."+defaultHostname) {
		return defaultHostname
	}
	return hostname
}

func HostnameValidator(v interface{}) error {
	hostname, valid := v.(string)
	if !valid {
		return errors.New("hostname is not a string")
	}

	if len(strings.TrimSpace(hostname)) < 1 {
		return errors.New("a value is required")
	}
	if strings.ContainsRune(hostname, '/') || strings.ContainsRune(hostname, ':') {
		return errors.New("invalid hostname")
	}
	return nil
}

func GraphQLEndpoint(hostname string) string {
	if IsEnterprise(hostname) {
		return fmt.Sprintf("https://%s/api/graphql", hostname)
	}
	return "https://api.github.com/graphql"
}

func RESTPrefix(hostname string) string {
	if IsEnterprise(hostname) {
		return fmt.Sprintf("https://%s/api/v3/", hostname)
	}
	return "https://api.github.com/"
}

func GistPrefix(hostname string) string {
	if IsEnterprise(hostname) {
		return fmt.Sprintf("https://%s/gist/", hostname)
	}
	return fmt.Sprintf("https://gist.%s/", hostname)
}
