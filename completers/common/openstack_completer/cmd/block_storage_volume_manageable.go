package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var block_storage_volume_manageableCmd = &cobra.Command{
	Use:   "manageable",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(block_storage_volume_manageableCmd).Standalone()

	block_storage_volumeCmd.AddCommand(block_storage_volume_manageableCmd)
}
