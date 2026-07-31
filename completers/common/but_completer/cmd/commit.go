package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:     "commit",
	Short:   "Create a commit",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "branching and committing",
}

func init() {
	carapace.Gen(commitCmd).Standalone()

	commitCmd.Flags().StringP("above", "A", "", "Place the commit above BRANCH_OR_COMMIT, which must be an applied branch or commit")
	commitCmd.Flags().Bool("allow-merged", false, "Allow targeting branches and commits that are already merged upstream")
	commitCmd.Flags().StringP("below", "B", "", "Place the commit below BRANCH_OR_COMMIT, which must be an applied branch or commit")
	commitCmd.Flags().StringP("branch", "b", "", "Place the commit on the branch BRANCH")
	commitCmd.Flags().Bool("empty", false, "Forces the commit to be empty regardless of repository state")
	commitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	commitCmd.Flags().BoolP("interactive", "i", false, "Open the TUI to interactively select what to commit")
	commitCmd.Flags().StringSliceP("message", "m", nil, "The message to use for the commit")
	commitCmd.Flags().Bool("no-message", false, "Creates the commit without a commit message")
	rootCmd.AddCommand(commitCmd)

	carapace.Gen(commitCmd).FlagCompletion(carapace.ActionMap{
		"above":  but.ActionTargets(),
		"below":  but.ActionTargets(),
		"branch": but.ActionLocalBranches(),
	})

	carapace.Gen(commitCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCliIds(but.CliIdsOpts{Changes: true}),
		).ToA().UniqueList(","),
	)
}
