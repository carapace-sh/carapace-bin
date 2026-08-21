package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_l3_conntrack_helper_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set L3 conntrack helper properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_l3_conntrack_helper_setCmd).Standalone()

	network_l3_conntrack_helper_setCmd.Flags().String("helper", "", "The netfilter conntrack helper module")
	network_l3_conntrack_helper_setCmd.Flags().String("port", "", "The network port for the netfilter conntrack target rule")
	network_l3_conntrack_helper_setCmd.Flags().String("protocol", "", "The network protocol for the netfilter conntrack target rule")
	network_l3_conntrack_helperCmd.AddCommand(network_l3_conntrack_helper_setCmd)
}
