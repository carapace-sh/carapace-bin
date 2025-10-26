package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a MCP server with stdio transport. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_serveCmd).Standalone()

	mcpCmd.AddCommand(mcp_serveCmd)
}
