package cmd

import (
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/mcp"
	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "run MCP server",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return mcp.NewMCPServer(cmd.Root().Version, os.Stdin, os.Stdout).Run()

	},
}

func init() {
	carapace.Gen(diffCmd).Standalone()
}
