package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/systemctl"
	"github.com/spf13/cobra"
)

func userFlag(cmd *cobra.Command) bool {
	if flag := cmd.Root().Flag("user"); flag != nil && flag.Changed {
		return true
	}
	return false
}

func ActionProperties(cmd *cobra.Command, units ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return systemctl.ActionProperties(systemctl.PropertiesOpts{User: userFlag(cmd), Units: units})
	})
}

func ActionServices(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return systemctl.ActionServices(userFlag(cmd))
	})
}

func ActionTargets(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return systemctl.ActionTargets(userFlag(cmd))
	})
}

func ActionUnits(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return systemctl.ActionUnits(systemctl.UnitOpts{User: userFlag(cmd), Active: true, Inactive: true})
	})
}

func ActionEnvironmentVariables(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return systemctl.ActionEnvironmentVariables(userFlag(cmd))
	})
}
