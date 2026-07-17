package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_flashApplianceCmd = &cobra.Command{
	Use:   "flash-appliance",
	Short: "Download a signed Tailscale appliance image and write it to a local disk [experimental]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_flashApplianceCmd).Standalone()

	configureCmd.AddCommand(configure_flashApplianceCmd)
}
