package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var image_metadef_object_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete metadata definitions object(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_metadef_object_deleteCmd).Standalone()

	image_metadef_objectCmd.AddCommand(image_metadef_object_deleteCmd)

	carapace.Gen(image_metadef_object_deleteCmd).PositionalCompletion(
		action.ActionImageMetadefNamespaces(image_metadef_object_deleteCmd),
		action.ActionImageMetadefObjects(image_metadef_object_deleteCmd),
	)
}
