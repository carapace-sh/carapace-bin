package upower

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionDevices completes devices
//
//	/org/freedesktop/UPower/devices/line_power_ADP1
//	/org/freedesktop/UPower/devices/DisplayDevice
func ActionDevices() carapace.Action {
	return carapace.ActionExecCommand("upower", "--enumerate")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
