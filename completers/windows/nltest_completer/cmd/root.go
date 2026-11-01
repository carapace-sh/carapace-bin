package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nltest",
	Short: "test the secure channel and trust relationships",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/nltest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
