package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proxy_dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Start local TLS proxy for database connections when using Teleport in single-port mode.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proxy_dbCmd).Standalone()

	proxy_dbCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect.")
	proxy_dbCmd.Flags().StringP("db-name", "n", "", "Database name to log in to.")
	proxy_dbCmd.Flags().StringP("db-roles", "r", "", "List of comma separate database roles to use for auto-provisioned user.")
	proxy_dbCmd.Flags().StringP("db-user", "u", "", "Database user to log in as.")
	proxy_dbCmd.Flags().Bool("disable-access-request", false, "Disable automatic resource access requests.")
	proxy_dbCmd.Flags().Bool("insecure-listen-anywhere", false, "Allows the local proxy to listen on any address without restrictions. WARNING: this will expose unsecured listener to anyone in the network. Only use when network access is otherwise restricted.")
	proxy_dbCmd.Flags().String("labels", "", "List of comma separated labels to filter by labels (e.g. key1=value1,key2=value2).")
	proxy_dbCmd.Flags().String("listen", "", "Specifies the source address used by proxy db listener. Mutually exclusive with --port.")
	proxy_dbCmd.Flags().Bool("no-disable-access-request", false, "Disable automatic resource access requests.")
	proxy_dbCmd.Flags().Bool("no-insecure-listen-anywhere", false, "Allows the local proxy to listen on any address without restrictions. WARNING: this will expose unsecured listener to anyone in the network. Only use when network access is otherwise restricted.")
	proxy_dbCmd.Flags().Bool("no-tunnel", false, "Open authenticated tunnel using database's client certificate so clients don't need to authenticate.")
	proxy_dbCmd.Flags().StringP("port", "p", "", "Specifies the source port used by proxy db listener.")
	proxy_dbCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"').")
	proxy_dbCmd.Flags().String("request-reason", "", "Reason for requesting access.")
	proxy_dbCmd.Flags().Bool("tunnel", false, "Open authenticated tunnel using database's client certificate so clients don't need to authenticate.")
	proxy_dbCmd.Flag("no-disable-access-request").Hidden = true
	proxy_dbCmd.Flag("no-insecure-listen-anywhere").Hidden = true
	proxy_dbCmd.Flag("no-tunnel").Hidden = true
	proxyCmd.AddCommand(proxy_dbCmd)
}
