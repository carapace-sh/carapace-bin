package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Set up MCP agents and configuration for Vercel integration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcpCmd).Standalone()

	mcpCmd.Flags().String("clients", "", "Comma-separated list of MCP clients to set up")
	mcpCmd.Flags().String("project", "", "Set up project-specific MCP access")

	rootCmd.AddCommand(mcpCmd)
}
