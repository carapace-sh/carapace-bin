package pacman

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionPackageSearch completes installable packages
//
//	a2ps (An Any to PostScript filter)
//	a52dec (A free library for decoding ATSC A/52 streams)
func ActionPackageSearch() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("pacman", "-Ss", "^"+c.Value)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^(?P<group>[^/]+)/(?P<name>[^ ]+) (?P<version>[^ ]+)(?P<context>.*)$`)

			vals := make([]string, 0)
			for i := 0; i < len(lines)-1; i += 2 {
				if matches := r.FindStringSubmatch(lines[i]); matches != nil {
					s := style.Default
					if strings.Contains(matches[4], "[installed]") {
						s = style.Blue
					}
					vals = append(vals, matches[2], strings.TrimSpace(lines[i+1]), s)
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}
