package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "timeout",
	Short: "pause the command processor for a specified number of seconds",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/timeout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("nobreak", false, "ignore key presses")
}
