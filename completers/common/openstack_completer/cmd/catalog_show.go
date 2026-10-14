package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var catalog_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display service catalog details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catalog_showCmd).Standalone()

	catalog_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	catalog_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	catalog_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	catalog_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	catalog_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	catalog_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	catalog_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	catalog_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	catalogCmd.AddCommand(catalog_showCmd)
}
