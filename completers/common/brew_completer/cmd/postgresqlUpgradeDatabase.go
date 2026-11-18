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

	postgresqlUpgradeDatabaseCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	postgresqlUpgradeDatabaseCmd.Flags().BoolP("help", "h", false, "Show this message.")
	postgresqlUpgradeDatabaseCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	postgresqlUpgradeDatabaseCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(postgresqlUpgradeDatabaseCmd)
}
