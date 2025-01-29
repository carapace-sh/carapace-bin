package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_mountCmd = &cobra.Command{
	Use:   "mount [options] [IMAGE...]",
	Short: "Mount an image's root filesystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_mountCmd).Standalone()

	image_mountCmd.Flags().BoolP("all", "a", false, "Mount all images")
	image_mountCmd.Flags().String("format", "", "Print the mounted images in specified format (json)")
	imageCmd.AddCommand(image_mountCmd)
}
