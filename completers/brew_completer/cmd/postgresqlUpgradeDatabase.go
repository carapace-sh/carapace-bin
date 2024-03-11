package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var postgresqlUpgradeDatabaseCmd = &cobra.Command{
	Use:     "postgresql-upgrade-database",
	Short:   "Upgrades the database for the `postgresql` formula",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(postgresqlUpgradeDatabaseCmd).Standalone()

	postgresqlUpgradeDatabaseCmd.Flags().Bool("debug", false, "Display any debugging information.")
	postgresqlUpgradeDatabaseCmd.Flags().Bool("help", false, "Show this message.")
	postgresqlUpgradeDatabaseCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	postgresqlUpgradeDatabaseCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(postgresqlUpgradeDatabaseCmd)
}
