package tofu

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionWorkspaces completes workspaces
func ActionWorkspaces() carapace.Action {
	return carapace.ActionExecCommand("tofu", "workspace", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^[* ] (?P<name>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				vals = append(vals, r.FindStringSubmatch(line)[1])
			}
		}
		return carapace.ActionValues(vals...)
	})
}
