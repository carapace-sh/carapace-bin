package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_agent_add_routerCmd = &cobra.Command{
	Use:   "router",
	Short: "Add router to an agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_agent_add_routerCmd).Standalone()

	network_agent_add_routerCmd.Flags().String("ha-chassis-priority", "", "HA Chassis priority, ranging from [0, 32767].")
	network_agent_add_routerCmd.Flags().Bool("l3", false, "Add router to an L3 agent")
	network_agent_addCmd.AddCommand(network_agent_add_routerCmd)
}
