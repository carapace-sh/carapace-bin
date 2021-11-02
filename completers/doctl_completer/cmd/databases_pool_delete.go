package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_pool_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a connection pool for a database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_pool_deleteCmd).Standalone()
	databases_pool_deleteCmd.Flags().BoolP("force", "f", false, "Delete connection pool without confirmation prompt")
	databases_poolCmd.AddCommand(databases_pool_deleteCmd)
}
