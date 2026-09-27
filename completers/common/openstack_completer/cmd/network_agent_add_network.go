package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_agent_add_networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Add network to an agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_agent_add_networkCmd).Standalone()

	network_agent_add_networkCmd.Flags().Bool("dhcp", false, "Add network to a DHCP agent")
	network_agent_addCmd.AddCommand(network_agent_add_networkCmd)
}
