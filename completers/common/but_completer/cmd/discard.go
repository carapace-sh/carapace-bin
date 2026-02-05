package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var discardCmd = &cobra.Command{
	Use:   "discard",
	Short: "Discard uncommitted changes from the worktree.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discardCmd).Standalone()

	discardCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(discardCmd)

	carapace.Gen(discardCmd).PositionalCompletion(
		carapace.Batch(
			git.ActionChanges(git.ChangeOpts{}.Default()),
			but.ActionCliIds(but.CliIdsOpts{Changes: true}),
		).ToA(),
	)
}
