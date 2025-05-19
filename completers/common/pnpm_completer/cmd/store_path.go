package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Returns the path to the active store directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_pathCmd).Standalone()

	storeCmd.AddCommand(store_pathCmd)
}
