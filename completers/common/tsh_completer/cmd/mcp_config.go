package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print client configuration details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_configCmd).Standalone()

	mcp_configCmd.Flags().BoolP("all", "R", false, "Select all MCP servers. Mutually exclusive with --labels or --query.")
	mcp_configCmd.Flags().Bool("auto-reconnect", false, "Automatically starts a new remote MCP session when the previous remote session is interrupted by network issues or tsh session expirations. Recommended for stateless MCP sessions. Defaults to true.")
	mcp_configCmd.Flags().String("client-config", "", "If specified, update the specified client config, assuming its format. \"claude\" for default Claude Desktop config, \"cursor\" for global Cursor MCP servers config, or specify a JSON file path. Can also be set with environment variable TELEPORT_MCP_CLIENT_CONFIG.")
	mcp_configCmd.Flags().String("format", "", "Format specifies the configuration format (claude, vscode, cursor). If not provided it will assume format from the configuration file, When no configuration file is provided it defaults to \"claude\".")
	mcp_configCmd.Flags().StringP("header", "H", "", "Extra custom headers used for streamable HTTP MCP servers.")
	mcp_configCmd.Flags().String("json-format", "auto", "Format the JSON file (pretty, compact, auto, none). auto saves in compact if the file is already compact, otherwise pretty. Can also be set with environment variable TELEPORT_MCP_CONFIG_JSON_FORMAT. Default is auto.")
	mcp_configCmd.Flags().String("labels", "", "List of comma separated labels to filter by labels (e.g. key1=value1,key2=value2).")
	mcp_configCmd.Flags().Bool("no-all", false, "Select all MCP servers. Mutually exclusive with --labels or --query.")
	mcp_configCmd.Flags().Bool("no-auto-reconnect", false, "Automatically starts a new remote MCP session when the previous remote session is interrupted by network issues or tsh session expirations. Recommended for stateless MCP sessions. Defaults to true.")
	mcp_configCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"').")
	mcp_configCmd.Flag("no-all").Hidden = true
	mcp_configCmd.Flag("no-auto-reconnect").Hidden = true
	mcpCmd.AddCommand(mcp_configCmd)
}
