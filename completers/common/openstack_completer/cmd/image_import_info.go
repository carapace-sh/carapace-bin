package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_import_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show available import methods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_import_infoCmd).Standalone()

	image_import_infoCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	image_import_infoCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	image_import_infoCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	image_import_infoCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	image_import_infoCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	image_import_infoCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	image_import_infoCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	image_import_infoCmd.Flags().String("variable", "", "==SUPPRESS==")
	image_importCmd.AddCommand(image_import_infoCmd)
}
