package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_mcp_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Connect AI agents to atuin's MCP server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_mcp_installCmd).Standalone()

	help_mcpCmd.AddCommand(help_mcp_installCmd)
}
