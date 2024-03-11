package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var proxy_dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Start local TLS proxy for database connections when using Teleport in single-port mode.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proxy_dbCmd).Standalone()

	proxy_dbCmd.Flags().String("cert-file", "", "Certificate file for proxy client TLS configuration")
	proxy_dbCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect")
	proxy_dbCmd.Flags().String("db-name", "", "Optional database name to log in to.")
	proxy_dbCmd.Flags().String("db-user", "", "Optional database user to log in as.")
	proxy_dbCmd.Flags().String("key-file", "", "Key file for proxy client TLS configuration")
	proxy_dbCmd.Flags().Bool("no-tunnel", false, "Open authenticated tunnel using database's client certificate so clients don't need to authenticate")
	proxy_dbCmd.Flags().StringP("port", "p", "", "Specifies the source port used by proxy db listener")
	proxy_dbCmd.Flags().Bool("tunnel", false, "Open authenticated tunnel using database's client certificate so clients don't need to authenticate")
	proxy_dbCmd.Flag("cert-file").Hidden = true
	proxy_dbCmd.Flag("key-file").Hidden = true
	proxy_dbCmd.Flag("no-tunnel").Hidden = true
	proxyCmd.AddCommand(proxy_dbCmd)

	carapace.Gen(proxy_dbCmd).FlagCompletion(carapace.ActionMap{
		"cert-file": carapace.ActionFiles(),
		"cluster":   tsh.ActionClusters(),
		"key-file":  carapace.ActionFiles(),
		"port":      net.ActionPorts(),
	})
}
