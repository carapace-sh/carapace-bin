package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "net",
	Short: "manage network resources, services, and accounts",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/net",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"accounts", "update user accounts and password policies",
			"computer", "add or remove a computer from the domain",
			"config", "display configurable services",
			"continue", "continue a paused service",
			"file", "display open shared files",
			"group", "manage global groups",
			"help", "display help",
			"helpmsg", "display information about a network message",
			"localgroup", "manage local groups",
			"name", "manage messaging names",
			"pause", "pause a service",
			"print", "display print jobs",
			"send", "send messages to other users",
			"session", "manage sessions",
			"share", "manage shared resources",
			"start", "start a service",
			"statistics", "display network statistics",
			"stop", "stop a service",
			"time", "display or set the time",
			"use", "connect to or disconnect from shared resources",
			"user", "manage user accounts",
			"view", "display network resources",
		),
	)
}
