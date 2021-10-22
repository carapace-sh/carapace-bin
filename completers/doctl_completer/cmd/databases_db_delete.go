package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_db_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the specified database from the cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_db_deleteCmd).Standalone()
	databases_db_deleteCmd.Flags().BoolP("force", "f", false, "Delete the database without a confirmation prompt")
	databases_dbCmd.AddCommand(databases_db_deleteCmd)
}
