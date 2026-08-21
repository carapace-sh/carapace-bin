package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_metadefCmd = &cobra.Command{
	Use:   "metadef",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_metadefCmd).Standalone()

	imageCmd.AddCommand(image_metadefCmd)
}
