package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_metadef_namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_metadef_namespaceCmd).Standalone()

	image_metadefCmd.AddCommand(image_metadef_namespaceCmd)
}
