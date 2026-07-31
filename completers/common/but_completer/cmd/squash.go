package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var squashCmd = &cobra.Command{
	Use:     "squash",
	Short:   "Squash commits, branches, or changes",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "editing commits",
}

func init() {
	carapace.Gen(squashCmd).Standalone()

	squashCmd.Flags().Bool("allow-merged", false, "Allow targeting branches and commits that are already merged upstream")
	squashCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	squashCmd.Flags().StringSliceP("message", "m", nil, "The message to use for the new commit")
	squashCmd.Flags().Bool("no-message", false, "Creates the commit without a commit message")
	squashCmd.Flags().StringP("target", "t", "", "The target to squash into")
	squashCmd.Flags().Bool("use-source-message", false, "Use the message of the source(s)")
	squashCmd.Flags().BoolP("use-target-message", "u", false, "Use the message of the target")
	rootCmd.AddCommand(squashCmd)

	carapace.Gen(squashCmd).FlagCompletion(carapace.ActionMap{
		"target": carapace.Batch(
			but.ActionCommits(),
			but.ActionCliIds(but.CliIdsOpts{Commits: true}),
		).ToA(),
	})

	carapace.Gen(squashCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCommits(),
			but.ActionCliIds(but.CliIdsOpts{Commits: true}),
		).ToA(),
	)
}
