package cmd

import (
	"github.com/carapace-sh/carapace"
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
}
