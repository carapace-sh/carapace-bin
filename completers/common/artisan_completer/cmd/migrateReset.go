package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrateResetCmd = &cobra.Command{
	Use:   "migrate:reset [--database [DATABASE]] [--force] [--path [PATH]] [--realpath] [--pretend]",
	Short: "Rollback all database migrations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateResetCmd).Standalone()

	migrateResetCmd.Flags().String("database", "", "The database connection to use")
	migrateResetCmd.Flags().Bool("force", false, "Force the operation to run when in production")
	migrateResetCmd.Flags().String("path", "", "The path(s) to the migrations files to be executed")
	migrateResetCmd.Flags().Bool("pretend", false, "Dump the SQL queries that would be run")
	migrateResetCmd.Flags().Bool("realpath", false, "Indicate any provided migration file paths are pre-resolved absolute paths")
	rootCmd.AddCommand(migrateResetCmd)
}
