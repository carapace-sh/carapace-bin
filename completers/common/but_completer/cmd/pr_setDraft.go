package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var pr_setDraftCmd = &cobra.Command{
	Use:   "set-draft",
	Short: "Mark existing reviews as draft",
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
