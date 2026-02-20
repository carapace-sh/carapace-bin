package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_expose_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove exposed binaries from the global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_expose_help_removeCmd).Standalone()

	global_expose_helpCmd.AddCommand(global_expose_help_removeCmd)
}
