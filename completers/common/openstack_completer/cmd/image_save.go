package cmd

import (
	"github.com/carapace-sh/carapace"
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
	imageCmd.AddCommand(image_saveCmd)
}
