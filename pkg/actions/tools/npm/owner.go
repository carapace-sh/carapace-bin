package npm

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionOwners completes owners
func ActionOwners(pkg string) carapace.Action {
	return carapace.ActionExecCommand("npm", "owner", "ls", pkg)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, strings.SplitN(line, " ", 2)...)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
