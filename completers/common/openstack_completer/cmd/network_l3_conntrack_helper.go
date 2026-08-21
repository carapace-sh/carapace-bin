package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_l3_conntrack_helperCmd = &cobra.Command{
	Use:   "helper",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_l3_conntrack_helperCmd).Standalone()

	network_l3_conntrackCmd.AddCommand(network_l3_conntrack_helperCmd)
}
