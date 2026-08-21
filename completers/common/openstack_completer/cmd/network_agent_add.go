package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_agent_addCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_agent_addCmd).Standalone()

	network_agentCmd.AddCommand(network_agent_addCmd)
}
