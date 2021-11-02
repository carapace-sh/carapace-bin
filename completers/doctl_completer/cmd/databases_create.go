package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/doctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var databases_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a database cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_createCmd).Standalone()
	databases_createCmd.Flags().String("engine", "pg", "The database engine to be used for the cluster. Possible values are: `pg` for PostgreSQL, `mysql`, and `redis`.")
	databases_createCmd.Flags().Int("num-nodes", 1, "The number of nodes in the database cluster. Valid values are are 1-3. In addition to the primary node, up to two standby nodes may be added for high availability.")
	databases_createCmd.Flags().String("private-network-uuid", "", "The UUID of a VPC to create the database cluster in; the default VPC for the region will be used if excluded")
	databases_createCmd.Flags().String("region", "nyc1", "The region where the database cluster will be created, e.g. `nyc1` or `sfo2`")
	databases_createCmd.Flags().String("size", "db-s-1vcpu-1gb", "The size of the nodes in the database cluster, e.g. `db-s-1vcpu-1gb`` for a 1 CPU, 1GB node")
	databases_createCmd.Flags().String("version", "", "The database engine version, e.g. 11 for PostgreSQL version 11")
	databasesCmd.AddCommand(databases_createCmd)

	carapace.Gen(databases_createCmd).FlagCompletion(carapace.ActionMap{
		"engine": carapace.ActionValues("mysql", "redis"),
		"region": action.ActionRegions(),
	})
}
