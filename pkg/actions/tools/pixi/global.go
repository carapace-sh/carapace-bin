package pixi

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionGlobalEnvironments completes global environment names
//
//	gh (2.86.0)
//	jj (0.38.0)
func ActionGlobalEnvironments() carapace.Action {
	return actionExecPixi("global", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^[├└]── (?P<name>[^:]+): (?P<version>\S+)`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("global environments")
}
