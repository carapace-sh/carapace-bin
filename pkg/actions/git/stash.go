package git

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionStashes completes stash names
//   stash@{0} (WIP on branchA: ABCDEF1 last commit msg)
//   stash@{1} (WIP on branchB: ABCDEF2 last commit msg)
func ActionStashes() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("git", "stash", "list").Output(); err != nil {
			return carapace.ActionValues(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, (len(lines)-1)*2)

			for index, line := range lines[:len(lines)-1] {
				splitted := strings.SplitN(line, ": ", 2)
				vals[index*2] = splitted[0]
				vals[(index*2)+1] = splitted[1]
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}
