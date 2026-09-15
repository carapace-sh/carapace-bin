package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available MCP server applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_lsCmd).Standalone()

	mcp_lsCmd.Flags().StringP("format", "f", "text", "Format output (text, json, yaml).")
	mcp_lsCmd.Flags().Bool("no-verbose", false, "Show extra MCP server fields.")
	mcp_lsCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"').")
	mcp_lsCmd.Flags().String("search", "", "List of comma separated search keywords or phrases enclosed in quotations (e.g. --search=foo,bar,\"some phrase\").")
	mcp_lsCmd.Flags().BoolP("verbose", "v", false, "Show extra MCP server fields.")
	mcp_lsCmd.Flag("no-verbose").Hidden = true
	mcpCmd.AddCommand(mcp_lsCmd)
}
