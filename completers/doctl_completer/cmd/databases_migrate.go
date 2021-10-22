package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/doctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var databases_migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate a database cluster to a new region",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_migrateCmd).Standalone()
	databases_migrateCmd.Flags().String("private-network-uuid", "", "The UUID of a VPC to create the database cluster in; the default VPC for the region will be used if excluded")
	databases_migrateCmd.Flags().String("region", "", "The region to which the database cluster should be migrated, e.g. `sfo2` or `nyc3`. (required)")
	databasesCmd.AddCommand(databases_migrateCmd)

	carapace.Gen(databases_migrateCmd).FlagCompletion(carapace.ActionMap{
		"region": action.ActionRegions(),
	})
}
