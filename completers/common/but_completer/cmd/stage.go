package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var stageCmd = &cobra.Command{
	Use:   "stage",
	Short: "Stages a file or hunk to a specific branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stageCmd).Standalone()

	stageCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(stageCmd)

	carapace.Gen(stageCmd).PositionalCompletion(
		carapace.Batch(
			git.ActionChanges(git.ChangeOpts{}.Default()),
			but.ActionCliIds(but.CliIdsOpts{Changes: true}),
		).ToA(),
		carapace.Batch(
			but.ActionLocalBranches(),
			but.ActionCliIds(but.CliIdsOpts{Branches: true}),
		).ToA(),
	)
}
