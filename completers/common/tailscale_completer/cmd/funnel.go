package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var funnelCmd = &cobra.Command{
	Use:   "funnel",
	Short: "Serve content and local servers on the internet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(funnelCmd).Standalone()

	funnelCmd.Flags().Bool("bg", false, "run the command as a background process")
	funnelCmd.Flags().Uint("https", 0, "expose an HTTPS server at the specified port")
	funnelCmd.Flags().Uint("proxy-protocol", 0, "PROXY protocol version (1 or 2) for TCP forwarding")
	funnelCmd.Flags().String("set-path", "", "append the specified path to the base URL for accessing the underlying service")
	funnelCmd.Flags().Uint("tcp", 0, "expose a TCP forwarder to forward raw TCP packets at the specified port")
	funnelCmd.Flags().Uint("tls-terminated-tcp", 0, "expose a TCP forwarder to forward TLS-terminated TCP packets at the specified port")
	funnelCmd.Flags().Bool("yes", false, "update without interactive prompts")
	rootCmd.AddCommand(funnelCmd)

	carapace.Gen(funnelCmd).FlagCompletion(carapace.ActionMap{
		"proxy-protocol": carapace.ActionValues("1", "2"),
	})
}
