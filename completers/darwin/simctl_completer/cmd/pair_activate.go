package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pairActivateCmd = &cobra.Command{
	Use:   "pair_activate",
	Short: "Set a given pair as active",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pairActivateCmd).Standalone()
	rootCmd.AddCommand(pairActivateCmd)
	carapace.Gen(pairActivateCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
