package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goto",
	Short: "direct cmd.exe to a labeled line in a batch program",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/goto",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
