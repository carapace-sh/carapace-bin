package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sc",
	Short: "communicates with and controls services",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/sc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"query", "display service status",
			"queryex", "display extended service status",
			"start", "start a service",
			"stop", "stop a service",
			"pause", "pause a service",
			"continue", "continue a paused service",
			"config", "display service configuration",
			"delete", "delete a service",
			"create", "create a service",
			"getdisplayname", "get display name",
			"getkeyname", "get key name",
			"enumdep", "enumerate dependent services",
			"sdshow", "display security descriptor",
			"sdset", "set security descriptor",
			"lock", "lock service database",
			"unlock", "unlock service database",
		),
	)
}
