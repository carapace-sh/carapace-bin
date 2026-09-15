package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proxy_appCmd = &cobra.Command{
	Use:   "app",
	Short: "Start local TLS proxy for app connection when using Teleport in single-port mode.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proxy_appCmd).Standalone()

	proxy_appCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect.")
	proxy_appCmd.Flags().Bool("https-tunnel", false, "Use the teleport-app-https ALPN protocol (HTTPS tunneled over mTLS) for HTTP apps.")
	proxy_appCmd.Flags().Bool("no-https-tunnel", false, "Use the teleport-app-https ALPN protocol (HTTPS tunneled over mTLS) for HTTP apps.")
	proxy_appCmd.Flags().StringP("port", "p", "", "Specifies the listening port used by the proxy app listener. Accepts an optional target port of a multi-port TCP app after a colon, e.g. \"1234:5678\".")
	proxy_appCmd.Flag("https-tunnel").Hidden = true
	proxy_appCmd.Flag("no-https-tunnel").Hidden = true
	proxyCmd.AddCommand(proxy_appCmd)
}
