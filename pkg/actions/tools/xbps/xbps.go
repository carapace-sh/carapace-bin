// package xbps contains Advanced Package Tool related actions
package xbps

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPackages completes installed packages
//
//	expat (XML parser library written in C)
//	gawk (GNU awk utility)
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
//
//	3mux (Terminal multiplexer inspired by i3)
//	3proxy (3proxy tiny proxy server)
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
