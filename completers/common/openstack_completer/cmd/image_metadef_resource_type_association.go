package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_metadef_resource_type_associationCmd = &cobra.Command{
	Use:   "association",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_metadef_resource_type_associationCmd).Standalone()

	image_metadef_resource_typeCmd.AddCommand(image_metadef_resource_type_associationCmd)
}
