package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Change specified preferences",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	setCmd.Flags().Bool("accept-dns", false, "accept DNS configuration from the admin panel")
	setCmd.Flags().String("accept-risk", "", "accept risk and skip confirmation for risk types")
	setCmd.Flags().Bool("accept-routes", false, "accept routes advertised by other Tailscale nodes")
	setCmd.Flags().Bool("advertise-connector", false, "offer to be an app connector for domain specific internet traffic")
	setCmd.Flags().Bool("advertise-exit-node", false, "offer to be an exit node for internet traffic")
	setCmd.Flags().String("advertise-routes", "", "routes to advertise to other nodes (comma-separated)")
	setCmd.Flags().Bool("auto-update", false, "automatically update to the latest available version")
	setCmd.Flags().String("exit-node", "", "Tailscale exit node (IP, base name, or auto:any) for internet traffic")
	setCmd.Flags().Bool("exit-node-allow-lan-access", false, "allow direct access to the local network when routing via an exit node")
	setCmd.Flags().String("hostname", "", "hostname to use instead of the one provided by the OS")
	setCmd.Flags().String("nickname", "", "nickname for the current account")
	setCmd.Flags().String("operator", "", "Unix username to allow to operate on tailscaled without sudo")
	setCmd.Flags().String("relay-server-port", "", "UDP port for the relay server to bind to")
	setCmd.Flags().String("relay-server-static-endpoints", "", "static IP:port endpoints to advertise for relay connections (comma-separated)")
	rootCmd.AddCommand(setCmd)

	carapace.Gen(setCmd).FlagCompletion(carapace.ActionMap{
		"accept-risk": carapace.ActionValues("lose-ssh", "mac-app-connector", "all"),
	})
}
