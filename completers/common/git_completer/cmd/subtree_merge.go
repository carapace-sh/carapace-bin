package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var subtree_mergeCmd = &cobra.Command{
	Use:   "merge <local-commit> [<repository>]",
	Short: "Merge recent changes up to <local-commit> into the <prefix> subtree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subtree_mergeCmd).Standalone()

	subtree_mergeCmd.Flags().StringP("message", "m", "", "Specify <message> as the commit message for the merge commit")
	subtree_mergeCmd.Flags().Bool("squash", false, "produce only a single commit that contains all the differences")
	subtreeCmd.AddCommand(subtree_mergeCmd)

	carapace.Gen(subtree_mergeCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
		git.ActionRepositorySearch(git.SearchOpts{}.Default()),
	)
}
