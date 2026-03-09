package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offline_rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "prepare system to perform offline transaction and reboot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offline_rebootCmd).Standalone()

	offline_rebootCmd.Flags().Bool("poweroff", false, "power off after transaction instead of rebooting")
	offlineCmd.AddCommand(offline_rebootCmd)
}
