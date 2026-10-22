package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var image_metadef_namespace_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete metadef namespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_metadef_namespace_deleteCmd).Standalone()

	image_metadef_namespaceCmd.AddCommand(image_metadef_namespace_deleteCmd)

	carapace.Gen(image_metadef_namespace_deleteCmd).PositionalAnyCompletion(action.ActionImageMetadefNamespaces(image_metadef_namespace_deleteCmd))
}
