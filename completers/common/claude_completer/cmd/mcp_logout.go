package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Clear stored OAuth credentials for an MCP server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_logoutCmd).Standalone()

	mcp_logoutCmd.Flags().BoolP("help", "h", false, "Display help for command")
	mcpCmd.AddCommand(mcp_logoutCmd)
}
