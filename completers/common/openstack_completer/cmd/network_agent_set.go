package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_agent_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set network agent properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_agent_setCmd).Standalone()

	network_agent_setCmd.Flags().String("description", "", "Set network agent description")
	network_agent_setCmd.Flags().Bool("disable", false, "Disable network agent")
	network_agent_setCmd.Flags().Bool("enable", false, "Enable network agent")
	network_agentCmd.AddCommand(network_agent_setCmd)
}
