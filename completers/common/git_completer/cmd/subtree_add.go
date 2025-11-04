package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var subtree_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create the <prefix> subtree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subtree_addCmd).Standalone()

	subtree_addCmd.Flags().StringP("message", "m", "", "Specify <message> as the commit message for the merge commit")
	subtree_addCmd.Flags().Bool("squash", false, "produce only a single commit that contains all the differences")
	subtreeCmd.AddCommand(subtree_addCmd)

	carapace.Gen(subtree_addCmd).PositionalCompletion(
		carapace.Batch(
			git.ActionRefs(git.RefOption{}.Default()),
			carapace.ActionValues("https://").NoSpace('/').UnlessF(func(c carapace.Context) bool {
				return strings.HasPrefix(c.Value, "https://")
			}),
			git.ActionRepositorySearch(git.SearchOpts{}.Default()).UnlessF(func(c carapace.Context) bool {
				return !strings.HasPrefix(c.Value, "https://")
			}),
		).ToA(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !strings.HasPrefix(c.Args[0], "https://") {
				return carapace.ActionValues()
			}
			return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: c.Args[0], Branches: true, Tags: true})
		}),
	)
}
