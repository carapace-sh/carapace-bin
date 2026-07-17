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
	setCmd.Flags().String("netfilter-mode", "", "netfilter mode (one of on, nodivert, off)")
	setCmd.Flags().String("nickname", "", "nickname for the current account")
	setCmd.Flags().String("operator", "", "Unix username to allow to operate on tailscaled without sudo")
	setCmd.Flags().String("relay-server-port", "", "UDP port for the relay server to bind to")
	setCmd.Flags().String("relay-server-static-endpoints", "", "static IP:port endpoints to advertise for relay connections (comma-separated)")
	setCmd.Flags().Bool("report-posture", false, "allow management plane to gather device posture information")
	setCmd.Flags().Bool("shields-up", false, "don't allow incoming connections")
	setCmd.Flags().Bool("snat-subnet-routes", false, "source NAT traffic to local routes advertised with --advertise-routes")
	setCmd.Flags().Bool("ssh", false, "run an SSH server, permitting access per tailnet admin's declared policy")
	setCmd.Flags().Bool("stateful-filtering", false, "apply stateful filtering to forwarded packets (subnet routers, exit nodes, and so on)")
	setCmd.Flags().Bool("unattended", false, "run in \"Unattended Mode\" where Tailscale keeps running even after the current GUI user logs out")
	setCmd.Flags().Bool("update-check", false, "notify about available Tailscale updates")
	setCmd.Flags().Bool("webclient", false, "expose the web interface for managing this node over Tailscale at port 5252")
	rootCmd.AddCommand(setCmd)

	carapace.Gen(setCmd).FlagCompletion(carapace.ActionMap{
		"accept-risk":    carapace.ActionValues("lose-ssh", "mac-app-connector", "all"),
		"netfilter-mode": carapace.ActionValues("on", "nodivert", "off"),
	})
}
