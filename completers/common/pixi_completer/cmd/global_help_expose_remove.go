package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_help_expose_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove exposed binaries from the global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_help_expose_removeCmd).Standalone()

	global_help_exposeCmd.AddCommand(global_help_expose_removeCmd)
}
