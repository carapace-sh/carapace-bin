package ssh

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-ssh/config"
)

// ActionHosts completes ssh hosts
func ActionHosts(configPath string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		hosts, _ := config.HostCandidates(configPath)
		return carapace.ActionValues(hosts...)
	}).Tag("ssh hosts")
}
