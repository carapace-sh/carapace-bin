package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an MCP server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_removeCmd).Standalone()

	mcp_removeCmd.Flags().BoolP("help", "h", false, "Display help for command")
	mcp_removeCmd.Flags().StringP("scope", "s", "", "Configuration scope (local, user, or project) - if not specified, removes from whichever scope it…")
	mcpCmd.AddCommand(mcp_removeCmd)

	carapace.Gen(mcp_removeCmd).FlagCompletion(carapace.ActionMap{
		"scope": carapace.ActionValues("local", "project", "user"),
	})
}
