package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var block_storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(block_storageCmd).Standalone()

	blockCmd.AddCommand(block_storageCmd)
}
