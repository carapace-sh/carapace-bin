// package xbps contains Advanced Package Tool related actions
package xbps

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPackageSearch completes installed packages
func ActionPackages() carapace.Action {
	return carapace.ActionExecCommand("xbps-query", "-l")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		re := regexp.MustCompile(`\S+\s(\S+)-\S+\s+(.*)`)
		for _, line := range lines {
			match := re.FindStringSubmatch(line)

			if len(match) > 1 {
				vals = append(vals, match[1], match[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionPackageSearch completes installable packages
func ActionPackageSearch() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("xbps-query", "-Rs", "--regex", "^"+c.Value)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			re := regexp.MustCompile(`\[.\] (\S+)-\S+\s+(.*)`)
			for _, line := range lines {
				match := re.FindStringSubmatch(line)

				if len(match) > 1 {
					vals = append(vals, match[1], match[2])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
