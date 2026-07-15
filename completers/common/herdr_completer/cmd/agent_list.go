package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agent_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List agents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_listCmd).Standalone()

	agentCmd.AddCommand(agent_listCmd)
}
