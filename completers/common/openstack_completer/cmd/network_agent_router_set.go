package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_agent_router_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set properties of a router associated to an agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_agent_router_setCmd).Standalone()

	network_agent_router_setCmd.Flags().String("ha-chassis-priority", "", "HA Chassis priority, ranging from [0, 32767].")
	network_agent_router_setCmd.MarkFlagRequired("ha-chassis-priority")
	network_agent_routerCmd.AddCommand(network_agent_router_setCmd)
}
