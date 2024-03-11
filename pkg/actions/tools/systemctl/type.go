package systemctl

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionUnitTypes completes unit types
//
//	service
//	mount
func ActionUnitTypes() carapace.Action {
	return carapace.ActionExecCommand("systemctl", "--type=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[1 : len(lines)-1]...)
	})
}
