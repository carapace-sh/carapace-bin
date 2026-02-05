package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var squashCmd = &cobra.Command{
	Use:   "squash",
	Short: "Squash commits together.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(squashCmd).Standalone()

	squashCmd.Flags().StringP("ai", "i", "", "Generate commit message using AI with optional user summary or instructions. Use --ai by itself or --ai=\"your instructions\" (equals sign required for value)")
	squashCmd.Flags().BoolP("drop-message", "d", false, "Drop source commit messages and keep only the target commit's message")
	squashCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	squashCmd.Flags().StringP("message", "m", "", "Provide a new commit message for the resulting commit")
	squashCmd.Flag("ai").NoOptDefVal = " "
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
