package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Launch Model Context Protocol server for AI assistants",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcpCmd).Standalone()

	mcpCmd.Flags().String("http", "", "Run as HTTP server instead of stdio")
	mcpCmd.Flag("http").NoOptDefVal = " "

	rootCmd.AddCommand(mcpCmd)

	carapace.Gen(mcpCmd).FlagCompletion(carapace.ActionMap{
		"http": net.ActionPorts(),
	})
}
