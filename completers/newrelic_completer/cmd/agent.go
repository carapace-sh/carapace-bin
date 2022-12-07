package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Utilities for New Relic Agents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentCmd).Standalone()
	rootCmd.AddCommand(agentCmd)
}
