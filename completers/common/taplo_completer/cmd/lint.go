package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:     "lint",
	Short:   "Lint TOML documents",
	Aliases: []string{"check", "validate"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lintCmd).Standalone()

	lintCmd.Flags().String("cache-path", "", "Set a cache path")
	lintCmd.Flags().StringP("config", "c", "", "Path to the Taplo configuration file")
	lintCmd.Flags().Bool("default-schema-catalogs", false, "Use the default online catalogs for schemas")
	lintCmd.Flags().Bool("no-auto-config", false, "Do not search for a configuration file")
	lintCmd.Flags().Bool("no-schema", false, "Disable all schema validations")
	lintCmd.Flags().String("schema", "", "URL to the schema to be used for validation")
	lintCmd.Flags().String("schema-catalog", "", "URL to a schema catalog (index) that is compatible with Schema Store or Taplo catalogs")
	rootCmd.AddCommand(lintCmd)

	carapace.Gen(lintCmd).FlagCompletion(carapace.ActionMap{
		"cache-path": carapace.ActionDirectories(),
		"config":     carapace.ActionFiles(),
	})

	carapace.Gen(lintCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
