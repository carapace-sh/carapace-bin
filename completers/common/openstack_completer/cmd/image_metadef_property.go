package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_metadef_propertyCmd = &cobra.Command{
	Use:   "property",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_metadef_propertyCmd).Standalone()

	image_metadefCmd.AddCommand(image_metadef_propertyCmd)
}
