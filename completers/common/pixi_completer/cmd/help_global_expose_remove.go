package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_global_expose_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove exposed binaries from the global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_global_expose_removeCmd).Standalone()

	help_global_exposeCmd.AddCommand(help_global_expose_removeCmd)
}
