package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping a host at the Tailscale layer, see how it routed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pingCmd).Standalone()

	pingCmd.Flags().String("c", "", "max number of pings to send (0 for infinity)")
	pingCmd.Flags().Bool("icmp", false, "do an ICMP-level ping through WireGuard")
	pingCmd.Flags().Bool("peerapi", false, "try hitting the peer's peerapi HTTP server")
	pingCmd.Flags().String("size", "", "size of the ping message (disco pings only)")
	pingCmd.Flags().String("timeout", "", "timeout before giving up on a ping")
	pingCmd.Flags().Bool("tsmp", false, "do a TSMP-level ping through WireGuard")
	pingCmd.Flags().Bool("until-direct", false, "stop once a direct path is established")
	pingCmd.Flags().Bool("verbose", false, "verbose output")
	rootCmd.AddCommand(pingCmd)
}
