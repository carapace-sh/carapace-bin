package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_replica_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a read-only database replica",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_replica_createCmd).Standalone()
	databases_replica_createCmd.Flags().String("private-network-uuid", "", "The UUID of a VPC to create the replica in; the default VPC for the region will be used if excluded")
	databases_replica_createCmd.Flags().String("region", "nyc1", "Specifies the region (e.g. nyc3, sfo2) in which to create the replica")
	databases_replica_createCmd.Flags().String("size", "db-s-1vcpu-1gb", "Specifies the machine size for the replica (e.g. db-s-1vcpu-1gb). Must be the same or equal to the original.")
	databases_replicaCmd.AddCommand(databases_replica_createCmd)
}
