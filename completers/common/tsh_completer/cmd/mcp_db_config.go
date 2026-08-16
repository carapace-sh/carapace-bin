package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mcp_db_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print client configuration details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcp_db_configCmd).Standalone()

	mcp_db_configCmd.Flags().String("client-config", "", "If specified, update the specified client config, assuming its format. \"claude\" for default Claude Desktop config, \"cursor\" for global Cursor MCP servers config, or specify a JSON file path. Can also be set with environment variable TELEPORT_MCP_CLIENT_CONFIG.")
	mcp_db_configCmd.Flags().StringP("db-name", "n", "", "Database name to log in to.")
	mcp_db_configCmd.Flags().StringP("db-user", "u", "", "Database user to log in as.")
	mcp_db_configCmd.Flags().String("format", "", "Format specifies the configuration format (claude, vscode, cursor). If not provided it will assume format from the configuration file, When no configuration file is provided it defaults to \"claude\".")
	mcp_db_configCmd.Flags().String("json-format", "auto", "Format the JSON file (pretty, compact, auto, none). auto saves in compact if the file is already compact, otherwise pretty. Can also be set with environment variable TELEPORT_MCP_CONFIG_JSON_FORMAT. Default is auto.")
	mcp_db_configCmd.Flags().Bool("no-overwrite", false, "Overwrites command and environment variable from the config file.")
	mcp_db_configCmd.Flags().Bool("overwrite", false, "Overwrites command and environment variable from the config file.")
	mcp_db_configCmd.Flag("no-overwrite").Hidden = true
	mcp_dbCmd.AddCommand(mcp_db_configCmd)
}
