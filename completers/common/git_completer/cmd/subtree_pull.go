package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var subtree_pullCmd = &cobra.Command{
	Use:   "pull <repository> <remote-ref>",
	Short: "Exactly like merge, but parallels git pull",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subtree_pullCmd).Standalone()

	subtree_pullCmd.Flags().StringP("message", "m", "", "Specify <message> as the commit message for the merge commit")
	subtree_pullCmd.Flags().Bool("squash", false, "produce only a single commit that contains all the differences")
	subtreeCmd.AddCommand(subtree_pullCmd)

	carapace.Gen(subtree_pullCmd).PositionalCompletion(
		git.ActionRepositorySearch(git.SearchOpts{}.Default()),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: c.Args[0], Branches: true, Tags: true})
		}),
	)
}
