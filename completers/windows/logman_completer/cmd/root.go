package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "logman",
	Short: "manage and schedule performance counter and event trace log collections",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/logman",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"create", "create a data collector",
			"delete", "delete a data collector",
			"query", "query data collector properties",
			"start", "start a data collector",
			"stop", "stop a data collector",
			"update", "update a data collector",
		),
	)
}
