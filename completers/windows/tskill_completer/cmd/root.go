package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tskill",
	Short: "ends a process running in a session on a terminal server",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/tskill",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
