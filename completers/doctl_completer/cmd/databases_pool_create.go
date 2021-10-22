package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_pool_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a connection pool for a database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_pool_createCmd).Standalone()
	databases_pool_createCmd.Flags().String("db", "", "The name of the specific database within the database cluster (required)")
	databases_pool_createCmd.Flags().String("mode", "transaction", "The pool mode for the connection pool, e.g. `session`, `transaction`, and `statement`")
	databases_pool_createCmd.Flags().Int("size", 0, "pool size (required)")
	databases_pool_createCmd.Flags().String("user", "", "The username for the database user (required)")
	databases_poolCmd.AddCommand(databases_pool_createCmd)

	// TODO flag completion
	carapace.Gen(databases_pool_createCmd).FlagCompletion(carapace.ActionMap{
		"mode": carapace.ActionValues("session", "transaction", "statement"),
	})
}
