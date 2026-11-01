package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionContainers() carapace.Action {
	return carapace.ActionExecCommand("nixos-container", "list")(func(output []byte) carapace.Action {
		return carapace.ActionValues(strings.Fields(string(output))...)
	})
}
