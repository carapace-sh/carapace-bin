package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loginServerCmd = &cobra.Command{
	Use:   "login-server",
	Short: "Run the multi-account login server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginServerCmd).Standalone()

	loginServerCmd.Flags().String("accounts", "", "map callers to accounts with the HuJSON FILE")
	loginServerCmd.Flags().Bool("allow-anonymous", false, "admit anonymous callers on plain TCP listeners")
	loginServerCmd.Flags().StringArray("extra-listen", nil, "also listen on ENDPOINT (repeatable)")
	loginServerCmd.Flags().String("listen", "", "listen on the root-only Unix ENDPOINT")
	loginServerCmd.Flags().String("log-format", "text", "write server logs as FORMAT: text or json")
	loginServerCmd.Flags().Bool("pam", true, "open PAM sessions and run PAM account checks")
	loginServerCmd.Flags().String("pam-service", "rex-login-server", "use the PAM policy for service NAME")
	loginServerCmd.Flags().String("runtime-root", "", "keep the socket and accounts/ below DIR")
	loginServerCmd.Flags().String("tailscaled-socket", "", "use the tailscaled LocalAPI socket at PATH")

	rootCmd.AddCommand(loginServerCmd)

	carapace.Gen(loginServerCmd).FlagCompletion(carapace.ActionMap{
		"accounts":          carapace.ActionFiles(),
		"log-format":        carapace.ActionValues("json", "text"),
		"runtime-root":      carapace.ActionDirectories(),
		"tailscaled-socket": carapace.ActionFiles(),
	})
}
