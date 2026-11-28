package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the language server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_serverCmd).Standalone()

	helpCmd.AddCommand(help_serverCmd)
}
