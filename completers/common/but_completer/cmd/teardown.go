package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var teardownCmd = &cobra.Command{
	Use:     "teardown",
	Short:   "Exit GitButler mode and return to normal Git workflow.",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "other commands",
}

func init() {
	carapace.Gen(teardownCmd).Standalone()

	teardownCmd.Flags().StringP("checkout-to", "c", "", "Explicit override for which local branch to checkout to")
	teardownCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(teardownCmd)

	carapace.Gen(teardownCmd).FlagCompletion(carapace.ActionMap{
		"checkout-to": but.ActionLocalBranches(),
	})
}
