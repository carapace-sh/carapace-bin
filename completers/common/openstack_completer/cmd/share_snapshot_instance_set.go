package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_snapshot_instance_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Explicitly update the state of a share snapshot instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_snapshot_instance_setCmd).Standalone()

	share_snapshot_instance_setCmd.Flags().String("status", "", "Indicate state to update the snapshot instance to.")
	share_snapshot_instanceCmd.AddCommand(share_snapshot_instance_setCmd)
}
