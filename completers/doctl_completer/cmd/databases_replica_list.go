package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_replica_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve list of read-only database replicas",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_replica_listCmd).Standalone()
	databases_replica_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Name`, `Region`, `Status`, `URI`, `Created`")
	databases_replica_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	databases_replicaCmd.AddCommand(databases_replica_listCmd)
}
