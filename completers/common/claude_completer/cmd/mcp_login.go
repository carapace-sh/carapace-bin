package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with an MCP server (HTTP, SSE, or claude.ai connector)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_loginCmd).Standalone()

	mcp_loginCmd.Flags().BoolP("help", "h", false, "Display help for command")
	mcp_loginCmd.Flags().Bool("no-browser", false, "Print the authorization URL instead of opening a browser…")
	mcpCmd.AddCommand(mcp_loginCmd)
}
