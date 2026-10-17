package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverTailscaleRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename the enabled Tailscale node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverTailscaleRenameCmd).Standalone()

	serverTailscaleCmd.AddCommand(serverTailscaleRenameCmd)
}
