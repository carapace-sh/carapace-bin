package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_metadef_property_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete metadef propert(ies)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_metadef_property_deleteCmd).Standalone()

	image_metadef_propertyCmd.AddCommand(image_metadef_property_deleteCmd)
}
