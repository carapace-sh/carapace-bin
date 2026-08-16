package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/devenv"
	"github.com/spf13/cobra"
)

func globalOpts(cmd *cobra.Command) devenv.GlobalOpts {
	opts := devenv.GlobalOpts{}
	if flag := cmd.Flag("from"); flag != nil && flag.Changed {
		opts.From = flag.Value.String()
	}
	if flag := cmd.Flag("impure"); flag != nil && flag.Changed {
		opts.Impure = true
	}
	if flag := cmd.Flag("option"); flag != nil && flag.Changed {
		opts.Options, _ = cmd.Flags().GetStringSlice("option")
	}
	if flag := cmd.Flag("override-input"); flag != nil && flag.Changed {
		opts.OverrideInputs, _ = cmd.Flags().GetStringSlice("override-input")
	}
	if flag := cmd.Flag("profile"); flag != nil && flag.Changed {
		opts.Profiles, _ = cmd.Flags().GetStringArray("profile")
	}
	if flag := cmd.Flag("system"); flag != nil && flag.Changed {
		opts.System = flag.Value.String()
	}
	return opts
}

// ActionProcesses completes processes defined in devenv.nix
func ActionProcesses(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return devenv.ActionProcesses(globalOpts(cmd))
	})
}

// ActionProfiles completes profiles defined in devenv.nix
func ActionProfiles(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return devenv.ActionProfiles(globalOpts(cmd))
	})
}

// ActionRunningProcesses completes processes managed by the running process manager
func ActionRunningProcesses(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return devenv.ActionRunningProcesses(globalOpts(cmd))
	})
}

// ActionScripts completes scripts defined in devenv.nix
func ActionScripts(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return devenv.ActionScripts(globalOpts(cmd))
	})
}

// ActionTasks completes tasks
func ActionTasks(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return devenv.ActionTasks(devenv.TaskOpts{GlobalOpts: globalOpts(cmd)}.Default())
	})
}
