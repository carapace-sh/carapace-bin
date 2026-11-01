package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/agy"
	"github.com/spf13/cobra"
)

var mcpRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an MCP server configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcpRemoveCmd).Standalone()
	mcpCmd.AddCommand(mcpRemoveCmd)

	carapace.Gen(mcpRemoveCmd).PositionalCompletion(
		agy.ActionMcpServers(),
	)
}
