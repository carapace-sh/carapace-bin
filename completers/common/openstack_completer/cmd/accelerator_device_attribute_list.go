package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accelerator_device_attribute_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all attributes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accelerator_device_attribute_listCmd).Standalone()

	accelerator_device_attribute_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	accelerator_device_attribute_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	accelerator_device_attribute_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	accelerator_device_attribute_listCmd.Flags().Bool("long", false, "List additional fields in output")
	accelerator_device_attribute_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	accelerator_device_attribute_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	accelerator_device_attribute_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	accelerator_device_attribute_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	accelerator_device_attribute_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	accelerator_device_attribute_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	accelerator_device_attribute_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	accelerator_device_attributeCmd.AddCommand(accelerator_device_attribute_listCmd)
}
