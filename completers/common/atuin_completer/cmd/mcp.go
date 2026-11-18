package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Start an MCP server exposing history search to AI tools (stdio)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcpCmd).Standalone()

	mcpCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(mcpCmd)
}
