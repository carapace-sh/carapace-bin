package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_shimCmd = &cobra.Command{
	Use:   "shim",
	Short: "Manage context-aware shims for packages that are not installed globally, so a project decides which version runs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_shimCmd).Standalone()

	helpCmd.AddCommand(help_shimCmd)
}
