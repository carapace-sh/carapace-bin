package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rebaseCmd = &cobra.Command{
	Use:   "rebase [branch]",
	Short: "Rebase a stack of branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebaseCmd).Standalone()

	rebaseCmd.Flags().Bool("abort", false, "Abort rebase and restore all branches")
	rebaseCmd.Flags().Bool("continue", false, "Continue rebase after resolving conflicts")
	rebaseCmd.Flags().Bool("downstack", false, "Only rebase branches from trunk to current branch")
	rebaseCmd.Flags().String("remote", "", "Remote to fetch from (defaults to auto-detected remote)")
	rebaseCmd.Flags().Bool("upstack", false, "Only rebase branches from current branch to top")
	rootCmd.AddCommand(rebaseCmd)

	carapace.Gen(rebaseCmd).FlagCompletion(carapace.ActionMap{
		"remote": git.ActionRemotes(),
	})

	carapace.Gen(rebaseCmd).PositionalCompletion(
		git.ActionLocalBranches(),
	)
}
