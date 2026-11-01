package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var block_storage_volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(block_storage_volumeCmd).Standalone()

	block_storageCmd.AddCommand(block_storage_volumeCmd)
}
