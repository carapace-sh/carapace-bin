package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_db_startCmd = &cobra.Command{
	Use:    "start",
	Short:  "Start a local MCP server for database access.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_db_startCmd).Standalone()

	mcp_dbCmd.AddCommand(mcp_db_startCmd)
}
