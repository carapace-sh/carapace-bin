package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Connect AI agents to atuin's MCP server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_installCmd).Standalone()

	mcp_installCmd.Flags().BoolP("help", "h", false, "Print help")
	mcpCmd.AddCommand(mcp_installCmd)
}
