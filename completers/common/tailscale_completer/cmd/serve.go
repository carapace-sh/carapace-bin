package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve content and local servers on your tailnet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serveCmd).Standalone()

	serveCmd.Flags().String("accept-app-caps", "", "app capabilities to forward to the server (comma-separated)")
	serveCmd.Flags().Bool("bg", false, "run the command as a background process")
	serveCmd.Flags().Uint("http", 0, "expose an HTTP server at the specified port")
	serveCmd.Flags().Uint("https", 0, "expose an HTTPS server at the specified port")
	serveCmd.Flags().Uint("proxy-protocol", 0, "PROXY protocol version (1 or 2) for TCP forwarding")
	serveCmd.Flags().String("service", "", "serve for a service with distinct virtual IP instead of the node itself")
	serveCmd.Flags().String("set-path", "", "append the specified path to the base URL for accessing the underlying service")
	serveCmd.Flags().Uint("tcp", 0, "expose a TCP forwarder to forward raw TCP packets at the specified port")
	serveCmd.Flags().Uint("tls-terminated-tcp", 0, "expose a TCP forwarder to forward TLS-terminated TCP packets at the specified port")
	serveCmd.Flags().Bool("tun", false, "forward all traffic to the local machine (only supported for services)")
	serveCmd.Flags().Bool("yes", false, "update without interactive prompts")
	rootCmd.AddCommand(serveCmd)

	carapace.Gen(serveCmd).FlagCompletion(carapace.ActionMap{
		"proxy-protocol": carapace.ActionValues("1", "2"),
	})
}
