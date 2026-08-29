package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:     "mcp",
	GroupID: "integration",
	Short:   "Manage MCP servers (add, remove, list, enable, disable)",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcpCmd).Standalone()
	rootCmd.AddCommand(mcpCmd)
}
