package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var amendCmd = &cobra.Command{
	Use:     "amend",
	Short:   "Amend one or more file changes into a specific commit and rebases any dependent commits",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "editing commits",
}

func init() {
	carapace.Gen(amendCmd).Standalone()

	amendCmd.Flags().StringSliceP("changes", "p", nil, "Uncommitted file or hunk CLI IDs to amend into the commit")
	amendCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(amendCmd)

	carapace.Gen(amendCmd).FlagCompletion(carapace.ActionMap{
		"changes": carapace.Batch(
			git.ActionChanges(git.ChangeOpts{}.Default()),
			but.ActionCliIds(but.CliIdsOpts{Changes: true}),
		).ToA().UniqueList(","),
	})

	carapace.Gen(amendCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCommits(),
			but.ActionCliIds(but.CliIdsOpts{Commits: true}),
		).ToA(),
	)
}
