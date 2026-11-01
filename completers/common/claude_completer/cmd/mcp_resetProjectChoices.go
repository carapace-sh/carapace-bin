package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_resetProjectChoicesCmd = &cobra.Command{
	Use:   "reset-project-choices",
	Short: "Reset all approved and rejected project-scoped (.mcp.json) servers within this project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_resetProjectChoicesCmd).Standalone()

	mcp_resetProjectChoicesCmd.Flags().BoolP("help", "h", false, "Display help for command")
	mcpCmd.AddCommand(mcp_resetProjectChoicesCmd)
}
