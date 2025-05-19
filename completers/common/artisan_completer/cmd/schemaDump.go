package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemaDumpCmd = &cobra.Command{
	Use:   "schema:dump [--database [DATABASE]] [--path [PATH]] [--prune]",
	Short: "Dump the given database schema",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemaDumpCmd).Standalone()

	schemaDumpCmd.Flags().String("database", "", "The database connection to use")
	schemaDumpCmd.Flags().String("path", "", "The path where the schema dump file should be stored")
	schemaDumpCmd.Flags().Bool("prune", false, "Delete all existing migration files")
	rootCmd.AddCommand(schemaDumpCmd)
}
