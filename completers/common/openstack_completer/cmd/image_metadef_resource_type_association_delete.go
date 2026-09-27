package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_metadef_resource_type_association_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete metadef resource type association",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_metadef_resource_type_association_deleteCmd).Standalone()

	image_metadef_resource_type_association_deleteCmd.Flags().Bool("force", false, "Force delete the resource type association if thenamespace is protected")
	image_metadef_resource_type_associationCmd.AddCommand(image_metadef_resource_type_association_deleteCmd)
}
