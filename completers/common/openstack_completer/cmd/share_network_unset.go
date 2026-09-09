package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_network_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset a share network property",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_network_unsetCmd).Standalone()

	share_network_unsetCmd.Flags().Bool("description", false, "Unset share network description.")
	share_network_unsetCmd.Flags().Bool("name", false, "Unset share network name.")
	share_network_unsetCmd.Flags().String("security-service", "", "Disassociate a security service from the share network.")
	share_networkCmd.AddCommand(share_network_unsetCmd)
}
