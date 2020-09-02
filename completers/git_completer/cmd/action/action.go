package action

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionRemotes() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("git", "remote").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValues(strings.Split(string(output), "\n")...)
		}
	})
}

// see  https://github.com/fish-shell/fish-shell/blob/93cb0e2abba4e5c2fa2153a18f60f5e196d5efa6/share/completions/git.fish
func ActionBranches() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("git", "log", `--pretty=tformat:%h;%<(64,trunc)%s`, "--max-count=50").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			result := regexp.MustCompile("[\\;\\\n]+").Split(string(output), -1)
			for index, val := range result {
				result[index] = strings.TrimSpace(val)
			}
            // TODO prevent sorting
			return carapace.ActionValuesDescribed(result[:len(result)-1]...)
		}
	})
}
