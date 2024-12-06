package saw

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionGroups completes log groups.
//
//	lambda/one
//	ecs/two
func ActionGroups() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("saw", "groups", "--prefix", c.Value)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines...)
		})
	}).Tag("groups")
}
