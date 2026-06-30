package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Connect to Tailscale, logging in if needed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upCmd).Standalone()

	upCmd.Flags().Bool("accept-dns", false, "accept DNS configuration from the admin panel")
	upCmd.Flags().String("accept-risk", "", "accept risk and skip confirmation for risk types")
	upCmd.Flags().Bool("accept-routes", false, "accept routes advertised by other Tailscale nodes")
	upCmd.Flags().Bool("advertise-connector", false, "advertise this node as an app connector")
	upCmd.Flags().Bool("advertise-exit-node", false, "offer to be an exit node for internet traffic")
	upCmd.Flags().String("advertise-routes", "", "routes to advertise to other nodes (comma-separated, e.g. \"10.0.0.0/8,192.168.0.0/24\")")
	upCmd.Flags().String("advertise-tags", "", "comma-separated ACL tags to request (e.g. \"tag:eng,tag:montreal\")")
	upCmd.Flags().String("audience", "", "audience used when requesting an ID token from an identity provider")
	upCmd.Flags().String("auth-key", "", "node authorization key")
	upCmd.Flags().String("client-id", "", "client ID used to generate authkeys via workload identity federation")
	upCmd.Flags().String("client-secret", "", "client secret used to generate authkeys via OAuth")
	upCmd.Flags().String("exit-node", "", "Tailscale exit node (IP, base name, or auto:any) for internet traffic")
	upCmd.Flags().Bool("exit-node-allow-lan-access", false, "allow direct access to the local network when routing via an exit node")
	upCmd.Flags().Bool("force-reauth", false, "force reauthentication")
	upCmd.Flags().String("hostname", "", "hostname to use instead of the one provided by the OS")
	upCmd.Flags().String("id-token", "", "ID token from the identity provider for workload identity federation")
	upCmd.Flags().Bool("json", false, "output in JSON format")
	upCmd.Flags().String("login-server", "", "base URL of control server")
	upCmd.Flags().String("operator", "", "Unix username to allow to operate on tailscaled without sudo")
	upCmd.Flags().Bool("qr", false, "show QR code for login URLs")
	upCmd.Flags().String("qr-format", "", "QR code formatting")
	upCmd.Flags().Bool("report-posture", false, "allow management plane to gather device posture information")
	upCmd.Flags().Bool("reset", false, "reset unspecified settings to their default values")
	upCmd.Flags().Bool("shields-up", false, "don't allow incoming connections")
	upCmd.Flags().Bool("ssh", false, "run an SSH server, permitting access per tailnet admin's declared policy")
	upCmd.Flags().String("timeout", "", "maximum amount of time to wait for tailscaled to enter a Running state")
	rootCmd.AddCommand(upCmd)

	carapace.Gen(upCmd).FlagCompletion(carapace.ActionMap{
		"accept-risk": carapace.ActionValues("lose-ssh", "mac-app-connector", "all"),
		"qr-format":   carapace.ActionValues("ascii", "auto", "large", "small"),
	})
}
