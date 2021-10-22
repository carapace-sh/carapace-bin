package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_replica_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a read-only database replica",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_replica_deleteCmd).Standalone()
	databases_replica_deleteCmd.Flags().BoolP("force", "f", false, "Deletes the replica without a confirmation prompt")
	databases_replicaCmd.AddCommand(databases_replica_deleteCmd)
}
