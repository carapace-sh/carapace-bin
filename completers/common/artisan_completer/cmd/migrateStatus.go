package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrateStatusCmd = &cobra.Command{
	Use:   "migrate:status [--database [DATABASE]] [--pending [PENDING]] [--path [PATH]] [--realpath]",
	Short: "Show the status of each migration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateStatusCmd).Standalone()

	migrateStatusCmd.Flags().String("database", "", "The database connection to use")
	migrateStatusCmd.Flags().String("path", "", "The path(s) to the migrations files to use")
	migrateStatusCmd.Flags().String("pending", "", "Only list pending migrations")
	migrateStatusCmd.Flags().Bool("realpath", false, "Indicate any provided migration file paths are pre-resolved absolute paths")
	rootCmd.AddCommand(migrateStatusCmd)
}
