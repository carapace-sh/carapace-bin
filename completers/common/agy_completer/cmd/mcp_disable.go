package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/agy"
	"github.com/spf13/cobra"
)

var mcpDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable an MCP server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcpDisableCmd).Standalone()
	mcpCmd.AddCommand(mcpDisableCmd)

	carapace.Gen(mcpDisableCmd).PositionalCompletion(
		agy.ActionMcpServers(),
	)
}
