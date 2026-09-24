package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var object_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Upload object to container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(object_createCmd).Standalone()

	object_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	object_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	object_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	object_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	object_createCmd.Flags().String("name", "", "Upload a file and rename it.")
	object_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	object_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	object_createCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	object_createCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	object_createCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	object_createCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	objectCmd.AddCommand(object_createCmd)
}
