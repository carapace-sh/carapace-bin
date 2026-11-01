package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var block_storage_resource_filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(block_storage_resource_filterCmd).Standalone()

	block_storage_resourceCmd.AddCommand(block_storage_resource_filterCmd)
}
