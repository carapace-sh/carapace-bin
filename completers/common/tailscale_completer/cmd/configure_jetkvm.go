package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_jetkvmCmd = &cobra.Command{
	Use:   "jetkvm",
	Short: "Configure JetKVM to run tailscaled at boot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_jetkvmCmd).Standalone()

	configureCmd.AddCommand(configure_jetkvmCmd)
}
