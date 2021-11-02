package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_replica_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve information about a read-only database replica",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_replica_getCmd).Standalone()
	databases_replica_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Name`, `Region`, `Status`, `URI`, `Created`")
	databases_replica_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	databases_replicaCmd.AddCommand(databases_replica_getCmd)
}
