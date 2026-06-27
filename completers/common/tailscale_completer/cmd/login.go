package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to a Tailscale account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()

	loginCmd.Flags().Bool("accept-dns", false, "accept DNS configuration from the admin panel")
	loginCmd.Flags().Bool("accept-routes", false, "accept routes advertised by other Tailscale nodes")
	loginCmd.Flags().Bool("advertise-connector", false, "advertise this node as an app connector")
	loginCmd.Flags().Bool("advertise-exit-node", false, "offer to be an exit node for internet traffic")
	loginCmd.Flags().String("advertise-routes", "", "routes to advertise to other nodes (comma-separated)")
	loginCmd.Flags().String("advertise-tags", "", "comma-separated ACL tags to request")
	loginCmd.Flags().String("audience", "", "audience used when requesting an ID token from an identity provider")
	loginCmd.Flags().String("auth-key", "", "node authorization key")
	loginCmd.Flags().String("client-id", "", "client ID for workload identity federation")
	loginCmd.Flags().String("client-secret", "", "client secret for OAuth workload identity federation")
	loginCmd.Flags().String("exit-node", "", "Tailscale exit node for internet traffic")
	loginCmd.Flags().Bool("exit-node-allow-lan-access", false, "allow direct access to the local network when routing via an exit node")
	loginCmd.Flags().String("hostname", "", "hostname to use instead of the one provided by the OS")
	loginCmd.Flags().String("id-token", "", "ID token from the identity provider")
	loginCmd.Flags().String("login-server", "", "base URL of control server")
	loginCmd.Flags().String("nickname", "", "short name for the account")
	loginCmd.Flags().String("operator", "", "Unix username to allow to operate on tailscaled without sudo")
	loginCmd.Flags().Bool("qr", false, "show QR code for login URLs")
	loginCmd.Flags().String("qr-format", "", "QR code formatting")
	loginCmd.Flags().Bool("report-posture", false, "allow management plane to gather device posture information")
	loginCmd.Flags().Bool("shields-up", false, "don't allow incoming connections")
	loginCmd.Flags().Bool("ssh", false, "run an SSH server")
	loginCmd.Flags().String("timeout", "", "maximum amount of time to wait for tailscaled to enter a Running state")
	rootCmd.AddCommand(loginCmd)

	carapace.Gen(loginCmd).FlagCompletion(carapace.ActionMap{
		"qr-format": carapace.ActionValues("ascii", "auto", "large", "small"),
	})
}
