package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var block_storage_resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(block_storage_resourceCmd).Standalone()

	block_storageCmd.AddCommand(block_storage_resourceCmd)
}
