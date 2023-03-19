package yarn

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
)

func actionYarn(args ...string) func(f func(output []byte) carapace.Action) carapace.Action {
	return func(f func(output []byte) carapace.Action) carapace.Action {
		return carapace.ActionExecCommandE("yarn", args...)(func(output []byte, err error) carapace.Action {
			lines := strings.Split(string(output), "\n")
			if err != nil {
				if _, ok := err.(*exec.ExitError); ok {
					return carapace.ActionMessage(lines[0]) // error is printed to stdout
				}
				return carapace.ActionMessage(err.Error())
			}
			return f(output)
		})
	}

}
