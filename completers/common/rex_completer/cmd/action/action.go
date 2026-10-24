package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rex"
	"github.com/spf13/cobra"
)

// ActionSessions completes session targets
func ActionSessions(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return rex.ActionSessions(flagValue(cmd, "server"))
	})
}

// ActionWindows completes window targets
func ActionWindows(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return rex.ActionWindows(rex.WindowsOpts{
			Server:  flagValue(cmd, "server"),
			Session: flagValue(cmd, "session"),
		})
	})
}

// ActionBlocks completes block targets
func ActionBlocks(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return rex.ActionBlocks(rex.BlockOpts{
			Server:  flagValue(cmd, "server"),
			Session: flagValue(cmd, "session"),
			Window:  flagValue(cmd, "window"),
		})
	})
}

// ActionClients completes client targets
func ActionClients(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return rex.ActionClients(flagValue(cmd, "server"))
	})
}

// ActionClientActions completes actions of a client or the server
func ActionClientActions(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return rex.ActionClientActions(flagValue(cmd, "client"))
	})
}

// ActionClientActionArgs completes argument names of the action given as first argument
func ActionClientActionArgs(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		name := ""
		if len(c.Args) > 0 {
			name = c.Args[0]
		}
		return rex.ActionClientActionArgs(rex.ClientActionArgsOpts{
			Action: name,
			Client: flagValue(cmd, "client"),
		})
	})
}

func flagValue(cmd *cobra.Command, name string) string {
	if f := cmd.Flag(name); f != nil && f.Changed {
		return f.Value.String()
	}
	return ""
}
