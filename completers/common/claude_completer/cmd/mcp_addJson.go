package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_addJsonCmd = &cobra.Command{
	Use:   "add-json",
	Short: "Add an MCP server (stdio, SSE, HTTP, or WebSocket) with a JSON string",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_addJsonCmd).Standalone()

	mcp_addJsonCmd.Flags().Bool("client-secret", false, "Prompt for OAuth client secret (or set MCP_CLIENT_SECRET env var)")
	mcp_addJsonCmd.Flags().BoolP("help", "h", false, "Display help for command")
	mcp_addJsonCmd.Flags().StringP("scope", "s", "", "Configuration scope (local, user, or project)")
	mcpCmd.AddCommand(mcp_addJsonCmd)

	carapace.Gen(mcp_addJsonCmd).FlagCompletion(carapace.ActionMap{
		"scope": carapace.ActionValues("local", "project", "user"),
	})
}
