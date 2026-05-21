package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/systemd"
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
		return systemd.ActionProperties(systemd.PropertiesOpts{User: userFlag(cmd), Units: units})
	})
}

func ActionServices(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return systemd.ActionServices(userFlag(cmd))
	})
}

func ActionTargets(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return systemd.ActionTargets(userFlag(cmd))
	})
}

func ActionUnits(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return systemd.ActionUnits(systemd.UnitOpts{User: userFlag(cmd), Active: true, Inactive: true})
	})
}

func ActionUnitFiles(cmd *cobra.Command, enabled bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		opts := systemd.UnitFileOpts{}.Default()
		opts.User = userFlag(cmd)
		opts.Enabled = enabled
		opts.Disabled = !enabled
		return systemd.ActionUnitFiles(opts)
	})
}

func ActionEnvironmentVariables(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return systemd.ActionEnvironmentVariables(userFlag(cmd))
	})
}
