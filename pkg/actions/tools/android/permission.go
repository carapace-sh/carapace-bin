package android

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPermissions completes permission names
//
//	android.permission.INTERNET
//	android.permission.CAMERA
func ActionPermissions() carapace.Action {
	return carapace.ActionExecCommand("pm", "list", "permissions")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if name, found := strings.CutPrefix(line, "permission:"); found {
				vals = append(vals, name)
			}
		}
		return carapace.ActionValues(vals...)
	})
}
