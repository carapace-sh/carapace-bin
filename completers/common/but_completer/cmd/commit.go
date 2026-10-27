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
	commitCmd.Flags().Bool("empty", false, "Create an empty commit even when there are changes")
	commitCmd.Flags().StringP("file", "F", "", "Read the commit message from FILE, or from stdin when FILE is -")
	commitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	commitCmd.Flags().BoolP("interactive", "i", false, "Open the TUI to interactively select what to commit")
	commitCmd.Flags().StringSliceP("message", "m", nil, "The message to use for the commit")
	commitCmd.Flags().Bool("no-message", false, "Create the commit without a message")
	commitCmd.Flags().BoolP("switch", "s", false, "Switch to the target branch instead of remaining in the GitButler workspace")
	rootCmd.AddCommand(commitCmd)

	carapace.Gen(commitCmd).FlagCompletion(carapace.ActionMap{
		"above":  but.ActionTargets(),
		"below":  but.ActionTargets(),
		"branch": but.ActionLocalBranches(),
		"file":   carapace.ActionFiles(),
	})

	carapace.Gen(commitCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCliIds(but.CliIdsOpts{Changes: true}),
		).ToA().UniqueList(","),
	)
}
