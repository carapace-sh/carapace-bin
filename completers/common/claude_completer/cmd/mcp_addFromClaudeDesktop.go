package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_addFromClaudeDesktopCmd = &cobra.Command{
	Use:   "add-from-claude-desktop",
	Short: "Import MCP servers from Claude Desktop (Mac and WSL only)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_addFromClaudeDesktopCmd).Standalone()

	mcp_addFromClaudeDesktopCmd.Flags().BoolP("help", "h", false, "Display help for command")
	mcp_addFromClaudeDesktopCmd.Flags().StringP("scope", "s", "", "Configuration scope (local, user, or project)")
	mcpCmd.AddCommand(mcp_addFromClaudeDesktopCmd)

	carapace.Gen(mcp_addFromClaudeDesktopCmd).FlagCompletion(carapace.ActionMap{
		"scope": carapace.ActionValues("local", "user", "project"),
	})
}
