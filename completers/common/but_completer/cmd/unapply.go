package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var unapplyCmd = &cobra.Command{
	Use:   "unapply",
	Short: "Unapply a branch from the workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unapplyCmd).Standalone()

	unapplyCmd.Flags().BoolP("force", "f", false, "Force unapply without confirmation")
	unapplyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(unapplyCmd)

	carapace.Gen(unapplyCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
