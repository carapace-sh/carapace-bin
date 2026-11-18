package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_helpCmd).Standalone()

	mcpCmd.AddCommand(mcp_helpCmd)
}
