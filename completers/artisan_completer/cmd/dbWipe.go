package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dbWipeCmd = &cobra.Command{
	Use:   "db:wipe [--database [DATABASE]] [--drop-views] [--drop-types] [--force]",
	Short: "Drop all tables, views, and types",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dbWipeCmd).Standalone()

	dbWipeCmd.Flags().String("database", "", "The database connection to use")
	dbWipeCmd.Flags().Bool("drop-types", false, "Drop all tables and types (Postgres only)")
	dbWipeCmd.Flags().Bool("drop-views", false, "Drop all tables and views")
	dbWipeCmd.Flags().Bool("force", false, "Force the operation to run when in production")
	rootCmd.AddCommand(dbWipeCmd)
}
