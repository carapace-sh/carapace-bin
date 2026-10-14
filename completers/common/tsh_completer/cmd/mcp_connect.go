package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_connectCmd = &cobra.Command{
	Use:    "connect",
	Short:  "Connect to an MCP server.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_connectCmd).Standalone()

	mcp_connectCmd.Flags().Bool("auto-reconnect", true, "Automatically starts a new remote MCP session when the previous remote session is interrupted by network issues or tsh session expirations. Recommended for stateless MCP sessions. Defaults to true.")
	mcp_connectCmd.Flags().StringP("header", "H", "", "Extra custom headers used for streamable HTTP MCP servers.")
	mcp_connectCmd.Flags().Bool("no-auto-reconnect", false, "Automatically starts a new remote MCP session when the previous remote session is interrupted by network issues or tsh session expirations. Recommended for stateless MCP sessions. Defaults to true.")
	mcp_connectCmd.Flag("no-auto-reconnect").Hidden = true
	mcpCmd.AddCommand(mcp_connectCmd)
}
