package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var subtree_pushCmd = &cobra.Command{
	Use:   "push <repository> [+][<local-commit>:]<remote-ref>",
	Short: "Push your subtree to different branches of the remote repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subtree_pushCmd).Standalone()

	subtree_pushCmd.Flags().String("annotate", "", "add <annotation> as a prefix to each commit message")
	subtree_pushCmd.Flags().StringP("branch", "b", "", "create a new branch called <branch> that contains the new history")
	subtree_pushCmd.Flags().Bool("ignore-joins", false, "regenerate the entire history")
	subtree_pushCmd.Flags().StringP("message", "m", "", "Specify <message> as the commit message for the merge commit")
	subtree_pushCmd.Flags().String("onto", "", "specify the commit ID <onto> to build history from")
	subtree_pushCmd.Flags().Bool("rejoin", false, "merge the newly created synthetic history back into your main project")
	subtreeCmd.AddCommand(subtree_pushCmd)

	// TODO check and fix completion
	carapace.Gen(subtree_pushCmd).PositionalCompletion(
		git.ActionRemotes(),
		carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return git.ActionRefs(git.RefOption{
					LocalBranches: true,
					Heads:         true,
					Tags:          true,
				}).NoSpace()
			default:
				return git.ActionRemoteBranchNames(c.Args[0])
			}
		}),
	)
}
