package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details about an MCP server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_getCmd).Standalone()

	mcp_getCmd.Flags().BoolP("help", "h", false, "Display help for command")
	mcpCmd.AddCommand(mcp_getCmd)
}
