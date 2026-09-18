package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_l3_conntrack_helper_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete L3 conntrack helper",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_l3_conntrack_helper_deleteCmd).Standalone()

	network_l3_conntrack_helperCmd.AddCommand(network_l3_conntrack_helper_deleteCmd)
}
