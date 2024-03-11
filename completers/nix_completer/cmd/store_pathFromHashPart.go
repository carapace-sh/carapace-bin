package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_pathFromHashPartCmd = &cobra.Command{
	Use:   "path-from-hash-part",
	Short: "get a store path from its hash part",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_pathFromHashPartCmd).Standalone()

	storeCmd.AddCommand(store_pathFromHashPartCmd)
}
