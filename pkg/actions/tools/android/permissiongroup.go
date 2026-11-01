package android

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPermissionGroups completes permission group names
//
//	android.permission-group.CONTACTS
//	android.permission-group.LOCATION
func ActionPermissionGroups() carapace.Action {
	return carapace.ActionExecCommand("pm", "list", "permission-groups")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if name, found := strings.CutPrefix(line, "permissionGroup:"); found {
				vals = append(vals, name)
			}
		}
		return carapace.ActionValues(vals...)
	})
}
