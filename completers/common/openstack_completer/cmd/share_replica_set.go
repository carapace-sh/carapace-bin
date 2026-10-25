package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var share_replica_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Explicitly set share replica status and/or replica-state and/or property",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_replica_setCmd).Standalone()

	share_replica_setCmd.Flags().String("property", "", "Set a property to this replica (repeat option to set multiple properties).")
	share_replica_setCmd.Flags().String("replica-state", "", "Indicate which replica_state to assign the replica.")
	share_replica_setCmd.Flags().String("status", "", "Indicate which status to assign the replica.")
	share_replicaCmd.AddCommand(share_replica_setCmd)

	carapace.Gen(share_replica_setCmd).FlagCompletion(carapace.ActionMap{
		"replica-state": carapace.ActionValues("active", "error", "in_sync", "out_of_sync"),
		"status":        carapace.ActionValues("available", "creating", "deleting", "error", "error_deleting"),
	})

	carapace.Gen(share_replica_setCmd).PositionalCompletion(action.ActionShareReplicas(share_replica_setCmd))
}
