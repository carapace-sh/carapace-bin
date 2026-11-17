package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverTailscaleCmd = &cobra.Command{
	Use:   "tailscale",
	Short: "Manage the server's Tailscale node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverTailscaleCmd).Standalone()

	serverCmd.AddCommand(serverTailscaleCmd)
}
