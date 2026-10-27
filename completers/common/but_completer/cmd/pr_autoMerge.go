package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var pr_autoMergeCmd = &cobra.Command{
	Use:   "auto-merge",
	Short: "Enable or disable the automatic merging of reviews",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_autoMergeCmd).Standalone()

	pr_autoMergeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pr_autoMergeCmd.Flags().BoolP("off", "d", false, "Disable automatic merging instead of enabling it")
	prCmd.AddCommand(pr_autoMergeCmd)

	carapace.Gen(pr_autoMergeCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCliIds(but.CliIdsOpts{Branches: true, Stacks: true}),
			but.ActionLocalBranches(),
			// TODO pr ids
		).ToA().UniqueList(","),
	)
}
