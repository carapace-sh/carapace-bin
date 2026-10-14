package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proxy_mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Start local proxy for MCP access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proxy_mcpCmd).Standalone()

	proxy_mcpCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect.")
	proxy_mcpCmd.Flags().StringP("port", "p", "", "Specifies the listening port used by the proxy app listener.")
	proxyCmd.AddCommand(proxy_mcpCmd)
}
