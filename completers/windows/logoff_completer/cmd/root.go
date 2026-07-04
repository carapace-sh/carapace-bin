package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "logoff",
	Short: "logs a user out of a session on a terminal server and deletes the session",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/logoff",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
