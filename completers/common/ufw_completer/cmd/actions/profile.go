package actions

import (
	"github.com/carapace-sh/carapace"
	"strings"
)

func ActionUfwProfiles() carapace.Action {
	return carapace.ActionExecCommand("ufw", "app", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")[1:]
		var profiles []string

		for _, line := range lines {
			if profile := strings.TrimSpace(line); profile != "" {
				profiles = append(profiles, profile, "")
			}
		}

		return carapace.ActionValues(profiles...)
	})
}
