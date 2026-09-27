package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_metadef_objectCmd = &cobra.Command{
	Use:   "object",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_metadef_objectCmd).Standalone()

	image_metadefCmd.AddCommand(image_metadef_objectCmd)
}
