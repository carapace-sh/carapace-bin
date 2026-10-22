package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
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

	carapace.Gen(image_metadef_property_deleteCmd).PositionalCompletion(
		action.ActionImageMetadefNamespaces(image_metadef_property_deleteCmd),
		action.ActionImageMetadefProperties(image_metadef_property_deleteCmd),
	)
}
