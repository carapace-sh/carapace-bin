package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_recursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Concurrently runs a command in all subdirectory projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_recursiveCmd).Standalone()

	helpCmd.AddCommand(help_recursiveCmd)
}
