package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_network_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete one or more share networks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_network_deleteCmd).Standalone()

	share_network_deleteCmd.Flags().Bool("wait", false, "Wait for the share network(s) to be deleted")
	share_networkCmd.AddCommand(share_network_deleteCmd)
}
