package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate [--database [DATABASE]] [--force] [--path [PATH]] [--realpath] [--schema-path [SCHEMA-PATH]] [--pretend] [--seed] [--seeder [SEEDER]] [--step] [--graceful] [--isolated [ISOLATED]]",
	Short: "Run the database migrations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateCmd).Standalone()

	migrateCmd.Flags().String("database", "", "The database connection to use")
	migrateCmd.Flags().Bool("force", false, "Force the operation to run when in production")
	migrateCmd.Flags().Bool("graceful", false, "Return a successful exit code even if an error occurs")
	migrateCmd.Flags().String("isolated", "", "Do not run the command if another instance of the command is already running")
	migrateCmd.Flags().String("path", "", "The path(s) to the migrations files to be executed")
	migrateCmd.Flags().Bool("pretend", false, "Dump the SQL queries that would be run")
	migrateCmd.Flags().Bool("realpath", false, "Indicate any provided migration file paths are pre-resolved absolute paths")
	migrateCmd.Flags().String("schema-path", "", "The path to a schema dump file")
	migrateCmd.Flags().Bool("seed", false, "Indicates if the seed task should be re-run")
	migrateCmd.Flags().String("seeder", "", "The class name of the root seeder")
	migrateCmd.Flags().Bool("step", false, "Force the migrations to be run so they can be rolled back individually")
	rootCmd.AddCommand(migrateCmd)
}
