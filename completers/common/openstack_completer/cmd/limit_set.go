package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var limit_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update information about a limit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(limit_setCmd).Standalone()

	limit_setCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	limit_setCmd.Flags().String("description", "", "Description of the limit")
	limit_setCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	limit_setCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	limit_setCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	limit_setCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	limit_setCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	limit_setCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	limit_setCmd.Flags().String("resource-limit", "", "The resource limit for the project to assume")
	limit_setCmd.Flags().String("variable", "", "==SUPPRESS==")
	limitCmd.AddCommand(limit_setCmd)
}
