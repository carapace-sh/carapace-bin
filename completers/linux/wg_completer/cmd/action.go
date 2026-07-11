package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionInterfaces completes WireGuard interface names from `wg show interfaces`.
func ActionInterfaces() carapace.Action {
	return carapace.ActionExecCommand("wg", "show", "interfaces")(func(output []byte) carapace.Action {
		lines := strings.Fields(string(output))
		return carapace.ActionValues(lines...)
	}).Tag("interfaces")
}