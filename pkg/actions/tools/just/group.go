package just

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionGroups completes groups
//
//	first
//	second
func ActionGroups(path string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"--groups"}
		if path != "" {
			args = append(args, "--justfile", path)
		}
		return carapace.ActionExecCommand("just", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[1:] {
				vals = append(vals, strings.TrimSpace(line))
			}
			return carapace.ActionValues(vals...)
		}).Tag("groups")
	})
}
