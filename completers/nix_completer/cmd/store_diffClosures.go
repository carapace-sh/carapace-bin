package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_diffClosuresCmd = &cobra.Command{
	Use:   "diff-closures",
	Short: "show what packages and versions were added and removed between two closures",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_diffClosuresCmd).Standalone()

	storeCmd.AddCommand(store_diffClosuresCmd)

	addEvaluationFlags(store_diffClosuresCmd)
	addFlakeFlags(store_diffClosuresCmd)
	addLoggingFlags(store_diffClosuresCmd)

	// TODO positional completion
}
