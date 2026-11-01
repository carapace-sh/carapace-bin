package tofu

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionProviders completes provider source addresses
//
//	registry.opentofu.org/hashicorp/aws
//	registry.opentofu.org/hashicorp/random
func ActionProviders() carapace.Action {
	return carapace.ActionExecCommand("tofu", "providers", "-no-color")(func(output []byte) carapace.Action {
		r := regexp.MustCompile(`provider\[([^\]]+)\]`)

		vals := make([]string, 0)
		for _, line := range strings.Split(string(output), "\n") {
			if match := r.FindStringSubmatch(line); match != nil {
				vals = append(vals, match[1])
			}
		}
		return carapace.ActionValues(vals...).Unique()
	})
}
