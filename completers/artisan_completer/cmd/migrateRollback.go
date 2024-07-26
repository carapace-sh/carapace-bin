package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrateRollbackCmd = &cobra.Command{
	Use:   "migrate:rollback [--database [DATABASE]] [--force] [--path [PATH]] [--realpath] [--pretend] [--step [STEP]] [--batch BATCH]",
	Short: "Rollback the last database migration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateRollbackCmd).Standalone()

	migrateRollbackCmd.Flags().String("batch", "", "The batch of migrations (identified by their batch number) to be reverted")
	migrateRollbackCmd.Flags().String("database", "", "The database connection to use")
	migrateRollbackCmd.Flags().Bool("force", false, "Force the operation to run when in production")
	migrateRollbackCmd.Flags().String("path", "", "The path(s) to the migrations files to be executed")
	migrateRollbackCmd.Flags().Bool("pretend", false, "Dump the SQL queries that would be run")
	migrateRollbackCmd.Flags().Bool("realpath", false, "Indicate any provided migration file paths are pre-resolved absolute paths")
	migrateRollbackCmd.Flags().String("step", "", "The number of migrations to be reverted")
	rootCmd.AddCommand(migrateRollbackCmd)
}
