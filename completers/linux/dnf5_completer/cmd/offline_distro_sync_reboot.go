package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineDistroSyncRebootCmd = &cobra.Command{
	Use:   "reboot [options]",
	Short: "prepare the system to perform the offline transaction and reboot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineDistroSyncRebootCmd).Standalone()

	offlineDistroSyncRebootCmd.Flags().Bool("poweroff", false, "Power off the system after the operation is complete")

	offlineDistroSyncCmd.AddCommand(offlineDistroSyncRebootCmd)
}
