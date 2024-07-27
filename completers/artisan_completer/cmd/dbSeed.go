package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dbSeedCmd = &cobra.Command{
	Use:   "db:seed [--class [CLASS]] [--database [DATABASE]] [--force] [--] [<class>]",
	Short: "Seed the database with records",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dbSeedCmd).Standalone()

	dbSeedCmd.Flags().String("class", "", "The class name of the root seeder")
	dbSeedCmd.Flags().String("database", "", "The database connection to seed")
	dbSeedCmd.Flags().Bool("force", false, "Force the operation to run when in production")
	rootCmd.AddCommand(dbSeedCmd)
}
