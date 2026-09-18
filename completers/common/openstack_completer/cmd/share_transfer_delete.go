package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_transfer_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove one or more transfers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_transfer_deleteCmd).Standalone()

	share_transferCmd.AddCommand(share_transfer_deleteCmd)
}
