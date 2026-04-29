package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionPlugins() carapace.Action {
	return carapace.ActionExecCommand("gatsby", "plugin", "ls")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			if after, ok := strings.CutPrefix(line, "- "); ok {
				vals = append(vals, after)
			}
		}
		return carapace.ActionValues(vals...)
	})
}
