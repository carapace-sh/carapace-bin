package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shimCmd = &cobra.Command{
	Use:   "shim",
	Short: "Manage context-aware shims for packages that are not installed globally, so a project decides which version runs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shimCmd).Standalone()

	shimCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(shimCmd)
}
