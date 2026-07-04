package defaults

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionDomains completes defaults preference domains
//
//	com.apple.something
//	Apple Global Domain
func ActionDomains() carapace.Action {
	return carapace.ActionExecCommand("defaults", "domains")(func(output []byte) carapace.Action {
		domains := strings.Split(strings.TrimRight(string(output), ",\n"), ",")
		vals := make([]string, 0)
		for _, domain := range domains {
			domain = strings.TrimSpace(domain)
			if domain != "" {
				vals = append(vals, domain)
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("domains").Uid("defaults", "domains")
}
