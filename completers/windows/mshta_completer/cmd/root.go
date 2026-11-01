package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mshta",
	Short: "execute HTML Application (HTA) files",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/mshta",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
