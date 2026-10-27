package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var unapplyCmd = &cobra.Command{
	Use:     "unapply",
	Short:   "Remove a branch and the rest of its stack from the workspace, keeping them to apply again later",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "branching and committing",
}

func init() {
	carapace.Gen(unapplyCmd).Standalone()

	unapplyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(unapplyCmd)

	carapace.Gen(unapplyCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
