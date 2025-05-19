package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var requestPullCmd = &cobra.Command{
	Use:     "request-pull",
	Short:   "Generates a summary of pending changes",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interaction].ID,
}

func init() {
	carapace.Gen(requestPullCmd).Standalone()

	requestPullCmd.Flags().BoolS("p", "p", false, "show patch text as well")
	rootCmd.AddCommand(requestPullCmd)

	carapace.Gen(requestPullCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
		git.ActionRepositorySearch(git.SearchOpts{}.Default()),
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
