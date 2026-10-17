package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverTailscaleCmd = &cobra.Command{
	Use:   "tailscale",
	Short: "Toggle tailnet serving on the running server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverTailscaleCmd).Standalone()

	serverCmd.AddCommand(serverTailscaleCmd)
}
