package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_transfer_request_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete volume transfer request(s).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_transfer_request_deleteCmd).Standalone()

	volume_transfer_requestCmd.AddCommand(volume_transfer_request_deleteCmd)
}
