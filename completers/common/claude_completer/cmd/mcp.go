package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Configure and manage MCP servers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcpCmd).Standalone()

	mcpCmd.Flags().BoolP("help", "h", false, "Display help for command")
	rootCmd.AddCommand(mcpCmd)
}
