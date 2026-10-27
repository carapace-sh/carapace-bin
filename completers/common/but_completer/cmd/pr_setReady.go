package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var pr_setReadyCmd = &cobra.Command{
	Use:   "set-ready",
	Short: "Mark existing reviews as ready for review",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_setReadyCmd).Standalone()

	pr_setReadyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	prCmd.AddCommand(pr_setReadyCmd)

	carapace.Gen(pr_setReadyCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCliIds(but.CliIdsOpts{Branches: true, Stacks: true}),
			but.ActionLocalBranches(),
			// TODO pr ids
		).ToA().UniqueList(","),
	)
}
