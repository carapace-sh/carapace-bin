package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/agy"
	"github.com/spf13/cobra"
)

var mcpEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable an MCP server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcpEnableCmd).Standalone()
	mcpCmd.AddCommand(mcpEnableCmd)

	carapace.Gen(mcpEnableCmd).PositionalCompletion(
		agy.ActionMcpServers(),
	)
}
