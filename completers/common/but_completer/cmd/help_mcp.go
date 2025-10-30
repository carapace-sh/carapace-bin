package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Starts up the MCP server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_mcpCmd).Standalone()

	helpCmd.AddCommand(help_mcpCmd)
}
