package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ttyd",
	Short: "ttyd is a tool for sharing terminal over the web",
	Long:  "https://github.com/tsl0922/ttyd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().StringP("auth-header", "H", "", "HTTP Header name for auth proxy")
	rootCmd.Flags().StringP("base-path", "b", "", "Expected base path for requests coming from a reverse proxy")
	rootCmd.Flags().BoolP("browser", "B", false, "Open terminal with the default system browser")
	rootCmd.Flags().BoolP("check-origin", "O", false, "Do not allow websocket connection from different origin")
	rootCmd.Flags().StringP("client-option", "t", "", "Send option to client")
	rootCmd.Flags().StringP("credential", "c", "", "Credential for basic authentication")
	rootCmd.Flags().StringP("cwd", "w", "", "Working directory to be set for the child program")
	rootCmd.Flags().StringP("debug", "d", "", "Set log level")
	rootCmd.Flags().StringP("gid", "g", "", "Group id to run with")
	rootCmd.Flags().BoolP("help", "h", false, "Print this text and exit")
	rootCmd.Flags().StringP("index", "I", "", "Custom index.html path")
	rootCmd.Flags().StringP("interface", "i", "", "Network interface to bind")
	rootCmd.Flags().BoolP("ipv6", "6", false, "Enable IPv6 support")
	rootCmd.Flags().StringP("max-clients", "m", "", "Maximum clients to support")
	rootCmd.Flags().BoolP("once", "o", false, "Accept only one client and exit on disconnection")
	rootCmd.Flags().StringP("ping-interval", "P", "", "Websocket ping interval(sec)")
	rootCmd.Flags().StringP("port", "p", "", "Port to listen")
	rootCmd.Flags().BoolP("readonly", "R", false, "Do not allow clients to write to the TTY")
	rootCmd.Flags().StringP("signal", "s", "", "Signal to send to the command when exit it")
	rootCmd.Flags().BoolP("ssl", "S", false, "Enable SSL")
	rootCmd.Flags().StringP("ssl-ca", "A", "", "SSL CA file path for client certificate verification")
	rootCmd.Flags().StringP("ssl-cert", "C", "", "SSL certificate file path")
	rootCmd.Flags().StringP("ssl-key", "K", "", "SSL key file path")
	rootCmd.Flags().StringP("terminal-type", "T", "", "Terminal type to report")
	rootCmd.Flags().StringP("uid", "u", "", "User id to run with")
	rootCmd.Flags().StringP("url-arg", "a", "", "Allow client to send command line arguments in URL")
	rootCmd.Flags().BoolP("version", "v", false, "Print the version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cwd": carapace.ActionDirectories(),
		"gid": os.ActionGroups(),
		"interface": carapace.Batch(
			net.ActionDevices(net.IncludedDevices{Ethernet: true, Wifi: true}),
			carapace.ActionFiles(),
		).ToA(),
		"port":     net.ActionPorts(),
		"signal":   ps.ActionKillSignals(),
		"ssl-ca":   carapace.ActionFiles(),
		"ssl-cert": carapace.ActionFiles(),
		"ssl-key":  carapace.ActionFiles(),
		"uid":      os.ActionUsers(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
