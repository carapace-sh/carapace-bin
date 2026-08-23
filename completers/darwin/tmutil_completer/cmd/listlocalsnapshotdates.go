package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listlocalsnapshotdatesCmd = &cobra.Command{
	Use:   "listlocalsnapshotdates",
	Short: "list creation dates of local snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listlocalsnapshotdatesCmd).Standalone()
	rootCmd.AddCommand(listlocalsnapshotdatesCmd)
}