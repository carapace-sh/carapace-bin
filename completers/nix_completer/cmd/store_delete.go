package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var store_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete paths from the Nix store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_deleteCmd).Standalone()

	store_deleteCmd.Flags().Bool("ignore-liveness", false, "Do not check whether the paths are reachable from a root")
	storeCmd.AddCommand(store_deleteCmd)

	addEvaluationFlags(store_deleteCmd)
	addFlakeFlags(store_deleteCmd)
	addLoggingFlags(store_deleteCmd)

	// TODO positional completion
}
