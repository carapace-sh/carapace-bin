package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var consistency_group_snapshot_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete consistency group snapshot(s).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consistency_group_snapshot_deleteCmd).Standalone()

	consistency_group_snapshotCmd.AddCommand(consistency_group_snapshot_deleteCmd)

	carapace.Gen(consistency_group_snapshot_deleteCmd).PositionalAnyCompletion(action.ActionConsistencyGroupSnapshots(consistency_group_snapshot_deleteCmd))
}
