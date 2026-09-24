package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_metadef_namespace_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset metadef namespace tags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_metadef_namespace_unsetCmd).Standalone()

	image_metadef_namespace_unsetCmd.Flags().Bool("all-tags", false, "Unset all metadef tags")
	image_metadef_namespace_unsetCmd.Flags().String("tag", "", "Unset a tag on this metadef namespace (repeat option to unset multiple tags)")
	image_metadef_namespaceCmd.AddCommand(image_metadef_namespace_unsetCmd)
}
