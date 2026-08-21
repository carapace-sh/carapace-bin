package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_group_type_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update a volume group type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_group_type_setCmd).Standalone()

	volume_group_type_setCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_group_type_setCmd.Flags().String("description", "", "New description for volume group type.")
	volume_group_type_setCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_group_type_setCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_group_type_setCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_group_type_setCmd.Flags().String("name", "", "New name for volume group type.")
	volume_group_type_setCmd.Flags().Bool("no-property", false, "Remove all properties from this volume group type (specify both --no-property and --property to remove the current properties before setting new properties)")
	volume_group_type_setCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_group_type_setCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_group_type_setCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_group_type_setCmd.Flags().Bool("private", false, "Make volume group type unavailable to other projects.")
	volume_group_type_setCmd.Flags().String("property", "", "Property to add or modify for this volume group type (repeat option to set multiple properties)")
	volume_group_type_setCmd.Flags().Bool("public", false, "Make volume group type available to other projects.")
	volume_group_type_setCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_group_typeCmd.AddCommand(volume_group_type_setCmd)
}
