package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "schtasks",
	Short: "schedule commands and programs to run periodically or at a specific time",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/schtasks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"create", "create a new scheduled task",
			"change", "change an existing scheduled task",
			"run", "run a scheduled task immediately",
			"end", "stop a running scheduled task",
			"delete", "delete a scheduled task",
			"query", "display scheduled tasks",
			"show", "display information about a scheduled task",
		),
	)
}
