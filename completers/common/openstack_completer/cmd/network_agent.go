package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_agentCmd).Standalone()

	networkCmd.AddCommand(network_agentCmd)
}
