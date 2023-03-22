package systemctl

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionUnits completes units
//
//	NetworkManager.service (Network Manager)
//	apparmor.service (Load AppArmor profiles)
func ActionUnits(user bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"list-units"}
		if user {
			args = append(args, "--user")
		}
		return carapace.ActionExecCommand("systemctl", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[1:] {
				if line == "" {
					break
				}
				fields := strings.Fields(line)
				vals = append(vals, fields[0], strings.Join(fields[4:], " "))
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
