package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resize a database cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_resizeCmd).Standalone()
	databases_resizeCmd.Flags().Int("num-nodes", 0, "The number of nodes in the database cluster. Valid values are are 1-3. In addition to the primary node, up to two standby nodes may be added for high availability. (required)")
	databases_resizeCmd.Flags().String("size", "", "The size of the nodes in the database cluster, e.g. `db-s-1vcpu-1gb`` for a 1 CPU, 1GB node (required)")
	databasesCmd.AddCommand(databases_resizeCmd)
}
