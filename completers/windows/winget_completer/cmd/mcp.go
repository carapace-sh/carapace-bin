package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "MCP information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcpCmd).Standalone()

	mcpCmd.Flags().Bool("disable", false, "Disable extended features. Requires store access.")
	mcpCmd.Flags().Bool("enable", false, "Enable extended features. Requires store access.")
	rootCmd.AddCommand(mcpCmd)
}
