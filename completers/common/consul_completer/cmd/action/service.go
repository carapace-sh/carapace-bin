package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func ActionServices(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO client flags from cmd
		// TODO filter options (namespace,...)
		return carapace.ActionExecCommand("consul", "catalog", "services")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}

func ActionServiceIdentity(cmd *cobra.Command) carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionServices(cmd).NoSpace()
		case 1:
			return ActionDatacenters(cmd).UniqueList(",")
		default:
			return carapace.ActionValues()
		}
	})
}
