package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var store_repairCmd = &cobra.Command{
	Use:   "repair",
	Short: "repair store path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_repairCmd).Standalone()

	storeCmd.AddCommand(store_repairCmd)

	addEvaluationFlags(store_repairCmd)
	addFlakeFlags(store_repairCmd)
	addLoggingFlags(store_repairCmd)

	// TODO positional completion
}
