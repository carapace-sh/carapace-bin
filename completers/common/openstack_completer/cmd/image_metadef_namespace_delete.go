package cmd

import (
	"github.com/carapace-sh/carapace"
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
}
