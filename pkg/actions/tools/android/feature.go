package android

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionFeatures completes system features
//
//	android.hardware.camera
//	android.hardware.bluetooth
func ActionFeatures() carapace.Action {
	return carapace.ActionExecCommand("pm", "list", "features")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if name, found := strings.CutPrefix(line, "feature:"); found {
				vals = append(vals, name)
			}
		}
		return carapace.ActionValues(vals...)
	})
}
