package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_synologyCmd = &cobra.Command{
	Use:   "synology",
	Short: "Configure Synology to enable outbound connections",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_synologyCmd).Standalone()

	configureCmd.AddCommand(configure_synologyCmd)
}
