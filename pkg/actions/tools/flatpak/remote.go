package flatpak

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionRemotes completes remotes
func ActionRemotes() carapace.Action {
	return carapace.ActionExecCommand("flatpak", "remotes", "--columns", "name,description")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines {
			if splitted := strings.SplitN(line, "\t", 2); len(splitted) == 2 {
				vals = append(vals, splitted...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
