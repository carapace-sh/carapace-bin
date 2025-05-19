package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Removes unreferenced packages from the store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_pruneCmd).Standalone()

	storeCmd.AddCommand(store_pruneCmd)
}
