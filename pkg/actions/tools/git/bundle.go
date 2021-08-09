package git

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionBundleHeads(file string) carapace.Action {
	return carapace.ActionExecCommand("git", "bundle", "list-heads", file)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			splitted := strings.SplitN(line, " ", 2)
			vals = append(vals, splitted[1], splitted[0])
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
