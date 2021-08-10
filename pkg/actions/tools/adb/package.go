package adb

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionPackages completes installed packages
//   com.google.android.apps.maps
//   com.google.android.music
func ActionPackages() carapace.Action {
	return carapace.ActionExecCommand("adb", "shell", "pm", "list", "packages")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, strings.TrimPrefix(line, "package:"))
		}
		return carapace.ActionValues(vals...)
	})
}
