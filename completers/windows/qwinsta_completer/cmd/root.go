package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qwinsta",
	Short: "display information about sessions on a terminal server",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/qwinsta",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
