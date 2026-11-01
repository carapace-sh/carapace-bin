package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Claude Code MCP server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_serveCmd).Standalone()

	mcp_serveCmd.Flags().BoolP("debug", "d", false, "Enable debug mode")
	mcp_serveCmd.Flags().BoolP("help", "h", false, "Display help for command")
	mcp_serveCmd.Flags().Bool("verbose", false, "Override verbose mode setting from config")
	mcpCmd.AddCommand(mcp_serveCmd)
}
