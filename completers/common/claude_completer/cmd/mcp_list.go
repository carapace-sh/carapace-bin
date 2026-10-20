package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured MCP servers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_listCmd).Standalone()

	mcp_listCmd.Flags().BoolP("help", "h", false, "Display help for command")
	mcpCmd.AddCommand(mcp_listCmd)
}
