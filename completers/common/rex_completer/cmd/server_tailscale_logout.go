package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverTailscaleLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Sign out of Tailscale and stop tailnet serving",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverTailscaleLogoutCmd).Standalone()

	serverTailscaleLogoutCmd.Flags().Bool("json", false, "output the resulting state as JSON")

	serverTailscaleCmd.AddCommand(serverTailscaleLogoutCmd)
}
