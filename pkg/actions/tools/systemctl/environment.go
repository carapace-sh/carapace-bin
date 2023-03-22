package systemctl

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionEnvironmentVariables completes environment variables
//
//	SWAYSOCK (/run/user/1000/sway-ipc.1000.978.sock)
//	WAYLAND_DISPLAY (wayland-1)
func ActionEnvironmentVariables(user bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"show-environment"}
		if user {
			args = append(args, "--user")
		}
		return carapace.ActionExecCommand("systemctl", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				vals = append(vals, strings.SplitN(line, "=", 2)...)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
