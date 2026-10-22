package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var image_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save an image locally",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_saveCmd).Standalone()

	image_saveCmd.Flags().String("chunk-size", "", "Size in bytes to read from the wire and buffer at one time (default: 1024)")
	image_saveCmd.Flags().String("file", "", "Downloaded image save filename (default: stdout)")
	image_saveCmd.Flags().String("store-preference", "", "Preferred store to download the image from (repeat option to specify multiple store preferences in order)")
	imageCmd.AddCommand(image_saveCmd)

	carapace.Gen(image_saveCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	carapace.Gen(image_saveCmd).PositionalCompletion(action.ActionImages(image_saveCmd))
}
