package gh

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh/config"
)

// ActionConfigHosts completes configured hosts
//
//	github.com
//	another.com
func ActionConfigHosts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		conf, err := config.ParseDefaultConfig()
		if err != nil {
			return carapace.ActionMessage("failed to parse DefaultConfig: " + err.Error())
		}
		hosts, err := conf.Hosts()
		if err != nil {
			return carapace.ActionMessage("failed to load hosts: " + err.Error())
		}
		return carapace.ActionValues(hosts...).Tag("config hosts")
	})
}

// ActionConfigUsers completes configured users
//
//	userA
//	userB
func ActionConfigUsers(host string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		conf, err := config.ParseDefaultConfig()
		if err != nil {
			return carapace.ActionMessage("failed to parse DefaultConfig: " + err.Error())
		}
		users, err := conf.Users(host)
		if err != nil {
			return carapace.ActionMessage("failed to load users: " + err.Error())
		}
		return carapace.ActionValues(users...)
	})
}
