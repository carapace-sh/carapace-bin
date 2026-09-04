package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_agent_remove_networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Remove network from an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_agent_remove_networkCmd).Standalone()

	network_agent_remove_networkCmd.Flags().Bool("dhcp", false, "Remove network from DHCP agent")
	network_agent_removeCmd.AddCommand(network_agent_remove_networkCmd)
}
