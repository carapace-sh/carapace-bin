package paru

import (
	"os/exec"
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
		return carapace.ActionExecCommandE("paru", "-Ssx", "^"+c.Value)(func(output []byte, err error) carapace.Action {
			if err != nil {
				if exitErr, ok := err.(*exec.ExitError); !ok || exitErr.ExitCode() != 1 { // pacman exits with `1` when there are no search results
					return carapace.ActionMessage(err.Error())
				}
			}

			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^(?P<group>[^/]+)/(?P<name>[^ ]+) (?P<version>[^ ]+)(?P<context>.*)$`)

			vals := make([]string, 0)
			for i := 0; i < len(lines)-1; i += 2 {
				if matches := r.FindStringSubmatch(lines[i]); matches != nil {
					s := style.Default
					if matches[1] == "aur" {
						s = style.Yellow
					}
					if strings.Contains(matches[4], "[Installed]") {
						s = style.Blue
					}
					vals = append(vals, matches[2], strings.TrimSpace(lines[i+1]), s)
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}
