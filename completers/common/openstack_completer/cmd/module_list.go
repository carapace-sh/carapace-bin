package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var module_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List module versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(module_listCmd).Standalone()

	module_listCmd.Flags().Bool("all", false, "Show all modules that have version information")
	module_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	module_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	module_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	module_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	module_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	module_listCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	module_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	module_listCmd.Flags().String("variable", "", "==SUPPRESS==")
	moduleCmd.AddCommand(module_listCmd)
}
