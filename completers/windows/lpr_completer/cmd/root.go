package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lpr",
	Short: "send a file to a computer running an LPD server",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/lpr",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
