package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverTailscaleRenameCmd = &cobra.Command{
	Use:   "rename NAME",
	Short: "Rename the enabled Tailscale node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverTailscaleRenameCmd).Standalone()

	serverTailscaleRenameCmd.Flags().Bool("json", false, "output the resulting state as JSON")

	serverTailscaleCmd.AddCommand(serverTailscaleRenameCmd)
}
