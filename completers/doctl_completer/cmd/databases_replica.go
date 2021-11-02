package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_replicaCmd = &cobra.Command{
	Use:   "replica",
	Short: "Display commands to manage read-only database replicas",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_replicaCmd).Standalone()
	databasesCmd.AddCommand(databases_replicaCmd)
}
