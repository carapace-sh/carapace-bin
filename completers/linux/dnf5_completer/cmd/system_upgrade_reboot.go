package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var systemUpgradeRebootCmd = &cobra.Command{
	Use:   "reboot [options]",
	Short: "prepare the system to perform the offline transaction and reboot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(systemUpgradeRebootCmd).Standalone()

	systemUpgradeRebootCmd.Flags().Bool("poweroff", false, "Power off the system after the operation is complete")

	systemUpgradeCmd.AddCommand(systemUpgradeRebootCmd)
}
