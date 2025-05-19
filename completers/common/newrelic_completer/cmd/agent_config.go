package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agent_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration utilities/helpers for New Relic agents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_configCmd).Standalone()
	agentCmd.AddCommand(agent_configCmd)
}
