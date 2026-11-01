package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_transfer_requestCmd = &cobra.Command{
	Use:   "request",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_transfer_requestCmd).Standalone()

	volume_transferCmd.AddCommand(volume_transfer_requestCmd)
}
