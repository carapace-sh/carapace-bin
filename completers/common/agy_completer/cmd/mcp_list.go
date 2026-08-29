package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcpListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured MCP servers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcpListCmd).Standalone()
	mcpCmd.AddCommand(mcpListCmd)
}
