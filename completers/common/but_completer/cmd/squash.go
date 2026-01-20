package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var squashCmd = &cobra.Command{
	Use:   "squash",
	Short: "Squash two commits together",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(squashCmd).Standalone()

	squashCmd.Flags().BoolP("drop-message", "d", false, "Drop the first commit's message and keep only the second commit's message")
	squashCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(squashCmd)

	carapace.Gen(squashCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCommits(),
			but.ActionCliIds(but.CliIdsOpts{Commits: true}),
		).ToA(),
		carapace.Batch(
			but.ActionCommits(),
			but.ActionCliIds(but.CliIdsOpts{Commits: true}),
		).ToA(),
	)
}
