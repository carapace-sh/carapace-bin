package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_migration_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show an in-progress live migration for a given server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_migration_showCmd).Standalone()

	server_migration_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_migration_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_migration_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_migration_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_migration_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_migration_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	server_migration_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_migration_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	server_migrationCmd.AddCommand(server_migration_showCmd)
}
