package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "change",
	Short: "change terminal server settings",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/change",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"logon", "enable/disable logons",
			"port", "change the listening port",
			"user", "change user settings",
		),
	)
}
