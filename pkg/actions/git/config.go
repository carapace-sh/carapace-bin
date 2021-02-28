package git

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionConfigs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO support different git folder
		if output, err := exec.Command("git", "help", "--config").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)
			for _, line := range strings.Split(string(output), "\n") {
				if line != "" && !strings.ContainsAny(line, "<>*") && !strings.Contains(line, "git help config") {
					vals = append(vals, line)
				}
			}
			return carapace.ActionValues(vals...).Invoke(c).ToMultiPartsA(".")
		}
	})
}
