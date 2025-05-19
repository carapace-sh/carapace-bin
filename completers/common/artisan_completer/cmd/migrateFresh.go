package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrateFreshCmd = &cobra.Command{
	Use:   "migrate:fresh [--database [DATABASE]] [--drop-views] [--drop-types] [--force] [--path [PATH]] [--realpath] [--schema-path [SCHEMA-PATH]] [--seed] [--seeder [SEEDER]] [--step]",
	Short: "Drop all tables and re-run all migrations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateFreshCmd).Standalone()

	migrateFreshCmd.Flags().String("database", "", "The database connection to use")
	migrateFreshCmd.Flags().Bool("drop-types", false, "Drop all tables and types (Postgres only)")
	migrateFreshCmd.Flags().Bool("drop-views", false, "Drop all tables and views")
	migrateFreshCmd.Flags().Bool("force", false, "Force the operation to run when in production")
	migrateFreshCmd.Flags().String("path", "", "The path(s) to the migrations files to be executed")
	migrateFreshCmd.Flags().Bool("realpath", false, "Indicate any provided migration file paths are pre-resolved absolute paths")
	migrateFreshCmd.Flags().String("schema-path", "", "The path to a schema dump file")
	migrateFreshCmd.Flags().Bool("seed", false, "Indicates if the seed task should be re-run")
	migrateFreshCmd.Flags().String("seeder", "", "The class name of the root seeder")
	migrateFreshCmd.Flags().Bool("step", false, "Force the migrations to be run so they can be rolled back individually")
	rootCmd.AddCommand(migrateFreshCmd)
}
