package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrateInstallCmd = &cobra.Command{
	Use:   "migrate:install [--database [DATABASE]]",
	Short: "Create the migration repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateInstallCmd).Standalone()

	migrateInstallCmd.Flags().String("database", "", "The database connection to use")
	rootCmd.AddCommand(migrateInstallCmd)
}
