package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var subtree_splitCmd = &cobra.Command{
	Use:   "split [<local-commit>] [<repository>]",
	Short: "Extract a new, synthetic project history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subtree_splitCmd).Standalone()

	subtree_splitCmd.Flags().String("annotate", "", "add <annotation> as a prefix to each commit message")
	subtree_splitCmd.Flags().StringP("branch", "b", "", "create a new branch called <branch> that contains the new history")
	subtree_splitCmd.Flags().Bool("ignore-joins", false, "regenerate the entire history")
	subtree_splitCmd.Flags().StringP("message", "m", "", "Specify <message> as the commit message for the merge commit")
	subtree_splitCmd.Flags().String("onto", "", "specify the commit ID <onto> to build history from")
	subtree_splitCmd.Flags().Bool("rejoin", false, "merge the newly created synthetic history back into your main project")
	subtreeCmd.AddCommand(subtree_splitCmd)

	carapace.Gen(subtree_splitCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
		git.ActionRepositorySearch(git.SearchOpts{}.Default()), // TODO verify
	)
}
