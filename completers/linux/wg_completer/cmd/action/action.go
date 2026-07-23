package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionInterfaces() carapace.Action {
	return carapace.ActionExecCommand("wg", "show", "interfaces")(func(output []byte) carapace.Action {
		lines := strings.Fields(string(output))
		return carapace.ActionValues(lines...)
	}).Tag("interfaces")
}
