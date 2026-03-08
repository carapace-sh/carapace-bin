package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_upgradeRebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "prepare system to perform offline transaction and reboot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_upgradeRebootCmd).Standalone()

	system_upgradeRebootCmd.Flags().Int("number", 0, "boot number")
	system_upgradeRebootCmd.Flags().Bool("poweroff", false, "power off after transaction instead of rebooting")

	system_upgradeCmd.AddCommand(system_upgradeRebootCmd)
}
