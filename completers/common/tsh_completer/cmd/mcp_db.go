package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database access for MCP servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_dbCmd).Standalone()

	mcpCmd.AddCommand(mcp_dbCmd)
}
