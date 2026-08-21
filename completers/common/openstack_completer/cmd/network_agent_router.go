package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_agent_routerCmd = &cobra.Command{
	Use:   "router",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_agent_routerCmd).Standalone()

	network_agentCmd.AddCommand(network_agent_routerCmd)
}
