package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate a package or its individual targets to use the given set of features",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_migrateCmd).Standalone()
	package_migrateCmd.Flags().SetInterspersed(false)

	package_migrateCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_migrateCmd.Flags().String("target", "", "A comma-separated list of targets to migrate")
	package_migrateCmd.Flags().String("to-feature", "", "A comma-separated list of Swift language features to migrate to")
	package_migrateCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_migrateCmd)
}
