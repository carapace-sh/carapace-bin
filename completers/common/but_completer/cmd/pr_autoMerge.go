package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var pr_autoMergeCmd = &cobra.Command{
	Use:   "auto-merge",
	Short: "Enable or disable the automatic merging of a review or reviews. If no reviews are specified, you will be prompted to select one or multiple of the review associated with branches in your workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_autoMergeCmd).Standalone()

	pr_autoMergeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pr_autoMergeCmd.Flags().BoolP("off", "d", false, "Whether to disable the automatic merging of the review(s)")
	prCmd.AddCommand(pr_autoMergeCmd)

	carapace.Gen(pr_autoMergeCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCliIds(but.CliIdsOpts{Branches: true, Stacks: true}),
			but.ActionLocalBranches(),
			// TODO pr ids
		).ToA().UniqueList(","),
	)
}
