package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a branch to the workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applyCmd).Standalone()

	applyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(applyCmd)

	carapace.Gen(applyCmd).PositionalCompletion(
		carapace.Batch(
			git.ActionLocalBranches(),
			git.ActionRemoteBranches(""),
		).ToA(),
	)
}
