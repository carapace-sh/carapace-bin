package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_metadef_property_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update metadef namespace property",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_metadef_property_setCmd).Standalone()

	image_metadef_property_setCmd.Flags().String("name", "", "Internal name of the property")
	image_metadef_property_setCmd.Flags().String("schema", "", "Valid JSON schema of the property")
	image_metadef_property_setCmd.Flags().String("title", "", "Property name displayed to the user")
	image_metadef_property_setCmd.Flags().String("type", "", "Property type")
	image_metadef_propertyCmd.AddCommand(image_metadef_property_setCmd)
}
