package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_add_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Associate project with image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_add_projectCmd).Standalone()

	image_add_projectCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	image_add_projectCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	image_add_projectCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	image_add_projectCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	image_add_projectCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	image_add_projectCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	image_add_projectCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	image_add_projectCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	image_add_projectCmd.Flags().String("variable", "", "==SUPPRESS==")
	image_addCmd.AddCommand(image_add_projectCmd)
}
