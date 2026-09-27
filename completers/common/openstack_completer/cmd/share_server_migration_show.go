package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_server_migration_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Gets migration progress of a given share server when copying",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_server_migration_showCmd).Standalone()

	share_server_migration_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_server_migration_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_server_migration_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_server_migration_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_server_migration_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_server_migration_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_server_migration_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_server_migration_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_server_migrationCmd.AddCommand(share_server_migration_showCmd)
}
