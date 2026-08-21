package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_metadef_object_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a metadef object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_metadef_object_updateCmd).Standalone()

	image_metadef_object_updateCmd.Flags().String("name", "", "New name of the object")
	image_metadef_objectCmd.AddCommand(image_metadef_object_updateCmd)
}
