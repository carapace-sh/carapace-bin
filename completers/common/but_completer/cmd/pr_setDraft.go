package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var pr_setDraftCmd = &cobra.Command{
	Use:   "set-draft",
	Short: "Set an existing review (or set of reviews) as draft. If no reviews are specified, you will be prompted to select one or multiple of the review associated with branches in your workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_setDraftCmd).Standalone()

	pr_setDraftCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	prCmd.AddCommand(pr_setDraftCmd)

	carapace.Gen(pr_setDraftCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCliIds(but.CliIdsOpts{Branches: true, Stacks: true}),
			but.ActionLocalBranches(),
			// TODO pr ids
		).ToA().UniqueList(","),
	)
}
