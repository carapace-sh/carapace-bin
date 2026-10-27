package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionStashIndexes completes stash entries as recorded by `grit stash list`
//
//	stash@{0} (WIP on branchA: ABCDEF1 fix bug)
//	stash@{1} (WIP on branchB: ABCDEF2 work)
func ActionStashIndexes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("grit", "stash", "list")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines[:len(lines)-1] {
				splitted := strings.SplitN(line, ": ", 2)
				fields := strings.SplitN(splitted[0], " ", 2)
				if len(splitted) == 2 {
					vals = append(vals, fields[0], splitted[1])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

// ActionHookNames completes hook scripts known to `grit hooks list`
//
//	pre-commit
//	commit-msg
func ActionHookNames() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("grit", "hooks", "list")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines[:len(lines)-1] {
				if name := strings.SplitN(line, "  ", 2)[0]; name != "" {
					vals = append(vals, name)
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}
