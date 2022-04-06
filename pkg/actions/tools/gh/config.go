package gh

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action/config"
)

func ActionConfigHosts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if config, err := config.ParseDefaultConfig(); err != nil {
			return carapace.ActionMessage("failed to parse DefaultConfig: " + err.Error())
		} else {
			if hosts, err := config.Hosts(); err != nil {
				return carapace.ActionMessage("failed ot loadd hosts: " + err.Error())
			} else {
				return carapace.ActionValues(hosts...)
			}
		}
	})
}
