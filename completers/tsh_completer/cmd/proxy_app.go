package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var proxy_appCmd = &cobra.Command{
	Use:   "app",
	Short: "Start local TLS proxy for app connection when using Teleport in single-port mode.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proxy_appCmd).Standalone()

	proxy_appCmd.Flags().StringP("port", "p", "", "Specifies the source port used by by the proxy app listener")
	proxyCmd.AddCommand(proxy_appCmd)

	carapace.Gen(proxy_appCmd).FlagCompletion(carapace.ActionMap{
		"port": net.ActionPorts(),
	})
}
