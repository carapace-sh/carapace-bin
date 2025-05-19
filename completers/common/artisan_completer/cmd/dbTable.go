package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dbTableCmd = &cobra.Command{
	Use:   "db:table [--database [DATABASE]] [--json] [--] [<table>]",
	Short: "Display information about the given database table",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dbTableCmd).Standalone()

	dbTableCmd.Flags().String("database", "", "The database connection")
	dbTableCmd.Flags().Bool("json", false, "Output the table information as JSON")
	rootCmd.AddCommand(dbTableCmd)
}
