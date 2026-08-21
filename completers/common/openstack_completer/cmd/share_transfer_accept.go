package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_transfer_acceptCmd = &cobra.Command{
	Use:   "accept",
	Short: "Accepts a share transfer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_transfer_acceptCmd).Standalone()

	share_transfer_acceptCmd.Flags().Bool("clear_rules", false, "Whether manila should clean up the access rules after the transfer is complete.")
	share_transferCmd.AddCommand(share_transfer_acceptCmd)
}
