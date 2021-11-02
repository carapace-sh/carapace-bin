package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a database cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_deleteCmd).Standalone()
	databases_deleteCmd.Flags().BoolP("force", "f", false, "Delete the database cluster without a confirmation prompt")
	databasesCmd.AddCommand(databases_deleteCmd)
}
