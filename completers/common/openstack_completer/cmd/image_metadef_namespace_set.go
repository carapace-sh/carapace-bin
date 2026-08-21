package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_metadef_namespace_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set metadef namespace properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_metadef_namespace_setCmd).Standalone()

	image_metadef_namespace_setCmd.Flags().String("description", "", "Set the description of the namespace")
	image_metadef_namespace_setCmd.Flags().String("display-name", "", "Set a user-friendly name for the namespace.")
	image_metadef_namespace_setCmd.Flags().Bool("private", false, "Metadef namespace is inaccessible to the public (default)")
	image_metadef_namespace_setCmd.Flags().Bool("protected", false, "Prevent metadef namespace from being deleted")
	image_metadef_namespace_setCmd.Flags().Bool("public", false, "Metadef namespace is accessible to the public")
	image_metadef_namespace_setCmd.Flags().String("tag", "", "Set a tag on this metadef namespace (repeat option to set multiple tags)")
	image_metadef_namespace_setCmd.Flags().Bool("unprotected", false, "Allow metadef namespace to be deleted (default)")
	image_metadef_namespaceCmd.AddCommand(image_metadef_namespace_setCmd)
}
