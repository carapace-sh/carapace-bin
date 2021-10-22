package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_replica_connectionCmd = &cobra.Command{
	Use:   "connection",
	Short: "Retrieve information for connecting to a read-only database replica",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_replica_connectionCmd).Standalone()
	databases_replicaCmd.AddCommand(databases_replica_connectionCmd)
}
