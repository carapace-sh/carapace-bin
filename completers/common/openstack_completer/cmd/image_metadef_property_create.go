package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_metadef_property_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a metadef property",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_metadef_property_createCmd).Standalone()

	image_metadef_property_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	image_metadef_property_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	image_metadef_property_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	image_metadef_property_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	image_metadef_property_createCmd.Flags().String("name", "", "Internal name of the property")
	image_metadef_property_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	image_metadef_property_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	image_metadef_property_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	image_metadef_property_createCmd.Flags().String("schema", "", "Valid JSON schema of the property")
	image_metadef_property_createCmd.Flags().String("title", "", "Property name displayed to the user")
	image_metadef_property_createCmd.Flags().String("type", "", "Property type")
	image_metadef_property_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	image_metadef_property_createCmd.MarkFlagRequired("name")
	image_metadef_property_createCmd.MarkFlagRequired("schema")
	image_metadef_property_createCmd.MarkFlagRequired("title")
	image_metadef_property_createCmd.MarkFlagRequired("type")
	image_metadef_propertyCmd.AddCommand(image_metadef_property_createCmd)
}
