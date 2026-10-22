package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var volume_group_snapshot_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a volume group snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_group_snapshot_deleteCmd).Standalone()

	volume_group_snapshotCmd.AddCommand(volume_group_snapshot_deleteCmd)

	carapace.Gen(volume_group_snapshot_deleteCmd).PositionalAnyCompletion(action.ActionVolumeGroupSnapshots(volume_group_snapshot_deleteCmd))
}
