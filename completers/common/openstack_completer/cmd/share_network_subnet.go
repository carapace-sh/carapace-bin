package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_network_subnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_network_subnetCmd).Standalone()

	share_networkCmd.AddCommand(share_network_subnetCmd)
}
