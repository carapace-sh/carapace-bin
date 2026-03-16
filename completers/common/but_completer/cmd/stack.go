package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Move branches on top of others to stack them.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stackCmd).Standalone()

	stackCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(stackCmd)

	carapace.Gen(stackCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionLocalBranches(),
			but.ActionCliIds(but.CliIdsOpts{Branches: true}),
		).ToA(),
		carapace.Batch(
			but.ActionLocalBranches(),
			but.ActionCliIds(but.CliIdsOpts{Branches: true}),
		).ToA().FilterArgs(),
	)
}
