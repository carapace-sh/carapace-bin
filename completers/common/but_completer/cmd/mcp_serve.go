package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve GitButler tools and MCP App resources over standard input/output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_serveCmd).Standalone()

	mcp_serveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	mcpCmd.AddCommand(mcp_serveCmd)
}