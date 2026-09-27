package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Removes unreferenced packages from the store. Unreferenced packages are packages that are not used by any projects on the system. Packages can become unreferenced after most installation operations, for instance when dependencies are made redundant",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_pruneCmd).Standalone()

	store_pruneCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	storeCmd.AddCommand(store_pruneCmd)
}
