package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display image details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_showCmd).Standalone()

	image_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	image_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	image_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	image_showCmd.Flags().Bool("human-readable", false, "Print image size in a human-friendly format.")
	image_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	image_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	image_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	image_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	image_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	imageCmd.AddCommand(image_showCmd)
}
