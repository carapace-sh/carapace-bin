// package make contains make related actions
package make

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionTargets completes targets
//
//	build
//	check
func ActionTargets(file string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if file == "" {
			file = "Makefile"
		}

		cmd := c.Command("make", "-qp", "--file", file)
		if output, err := cmd.Output(); err != nil && cmd.ProcessState.ExitCode() != 1 {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^(?P<target>[a-zA-Z0-9][^$#\/\t=]*):([^=]|$)`)

			vals := make([]string, 0)
			skip := false
			for _, line := range lines {
				if strings.HasPrefix(line, "# Not a target:") {
					skip = true
					continue
				} else if skip {
					skip = false
					continue
				}

				if r.MatchString(line) {
					vals = append(vals, r.FindStringSubmatch(line)[1])
				}
			}
			return carapace.ActionValues(vals...)
		}
	})
}
