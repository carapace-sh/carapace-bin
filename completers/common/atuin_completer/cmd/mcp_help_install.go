package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_help_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Connect AI agents to atuin's MCP server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_help_installCmd).Standalone()

	mcp_helpCmd.AddCommand(mcp_help_installCmd)
}
