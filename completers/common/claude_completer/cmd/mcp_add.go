package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an MCP server to Claude Code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_addCmd).Standalone()

	mcp_addCmd.Flags().String("callback-port", "", "Fixed port for OAuth callback (for servers requiring pre-registered redirect URIs)")
	mcp_addCmd.Flags().String("client-id", "", "OAuth client ID for HTTP/SSE servers")
	mcp_addCmd.Flags().Bool("client-secret", false, "Prompt for OAuth client secret (or set MCP_CLIENT_SECRET env var)")
	mcp_addCmd.Flags().StringSliceP("env", "e", nil, "Set environment variables (e.g. -e KEY=value)")
	mcp_addCmd.Flags().StringSliceP("header", "H", nil, "Set headers for HTTP/SSE servers (e.g. -H \"X-Api-Key: abc123\" -H \"X-Custom: value\")")
	mcp_addCmd.Flags().BoolP("help", "h", false, "Display help for command")
	mcp_addCmd.Flags().StringP("scope", "s", "", "Configuration scope (local, user, or project)")
	mcp_addCmd.Flags().StringP("transport", "t", "", "Transport type (stdio, sse, http)")
	mcpCmd.AddCommand(mcp_addCmd)

	carapace.Gen(mcp_addCmd).FlagCompletion(carapace.ActionMap{
		"scope":     carapace.ActionValues("local", "project", "user"),
		"transport": carapace.ActionValues("http", "sse", "stdio"),
	})
}
