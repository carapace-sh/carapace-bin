package ssh

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/kevinburke/ssh_config"
)

// ActionHosts completes ssh hosts
func ActionHosts(config string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		settings := ssh_config.DefaultUserSettings
		if config != "" {
			// TODO this skips the system config
			settings.ConfigFinder(func() string {
				return config
			})
		}

		vals := make([]string, 0)
		for _, host := range settings.GetHosts() {
			for _, pattern := range host.Patterns {
				if s := strings.ReplaceAll(pattern.String(), `\`, ""); !strings.ContainsAny(s, `+*?()|[]{}^$!`) {
					vals = append(vals, s)
				}
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("ssh hosts")
}
