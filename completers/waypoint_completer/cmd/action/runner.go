package action

import (
	"github.com/rsteube/carapace"
	"strings"
)

func ActionRunnerProfiles() carapace.Action {
	return carapace.ActionExecCommand("waypoint", "runner", "profile", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if splitted := strings.Split(line, "|"); len(splitted) == 4 {
				name := strings.TrimSpace(splitted[0])
				if name != "NAME" {
					vals = append(vals, name, strings.TrimSpace(splitted[1]))
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
