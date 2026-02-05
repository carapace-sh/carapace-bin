package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var pickCmd = &cobra.Command{
	Use:   "pick",
	Short: "Cherry-pick a commit from an unapplied branch into an applied virtual branch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pickCmd).Standalone()

	pickCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(pickCmd)

	carapace.Gen(pickCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCliIds(but.CliIdsOpts{}.Default()),
			but.ActionTargets(), // TODO targets correct?
		).ToA(),
		but.ActionLocalBranches(),
	)
}
