package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_agent_remove_routerCmd = &cobra.Command{
	Use:   "router",
	Short: "Remove router from an agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_agent_remove_routerCmd).Standalone()

	network_agent_remove_routerCmd.Flags().Bool("l3", false, "Remove router from an L3 agent")
	network_agent_removeCmd.AddCommand(network_agent_remove_routerCmd)
}
