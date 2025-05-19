package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrateRefreshCmd = &cobra.Command{
	Use:   "migrate:refresh [--database [DATABASE]] [--force] [--path [PATH]] [--realpath] [--seed] [--seeder [SEEDER]] [--step [STEP]]",
	Short: "Reset and re-run all migrations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateRefreshCmd).Standalone()

	migrateRefreshCmd.Flags().String("database", "", "The database connection to use")
	migrateRefreshCmd.Flags().Bool("force", false, "Force the operation to run when in production")
	migrateRefreshCmd.Flags().String("path", "", "The path(s) to the migrations files to be executed")
	migrateRefreshCmd.Flags().Bool("realpath", false, "Indicate any provided migration file paths are pre-resolved absolute paths")
	migrateRefreshCmd.Flags().Bool("seed", false, "Indicates if the seed task should be re-run")
	migrateRefreshCmd.Flags().String("seeder", "", "The class name of the root seeder")
	migrateRefreshCmd.Flags().String("step", "", "The number of migrations to be reverted & re-run")
	rootCmd.AddCommand(migrateRefreshCmd)
}
