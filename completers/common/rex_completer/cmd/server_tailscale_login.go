package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverTailscaleLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Request a Tailscale sign-in link",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverTailscaleLoginCmd).Standalone()

	serverTailscaleCmd.AddCommand(serverTailscaleLoginCmd)
}
