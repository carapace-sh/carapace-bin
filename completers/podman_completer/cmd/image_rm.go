package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_rmCmd = &cobra.Command{
	Use:   "rm [options] IMAGE [IMAGE...]",
	Short: "Remove one or more images from local storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_rmCmd).Standalone()

	image_rmCmd.Flags().BoolP("all", "a", false, "Remove all images")
	image_rmCmd.Flags().BoolP("force", "f", false, "Force Removal of the image")
	image_rmCmd.Flags().BoolP("ignore", "i", false, "Ignore errors if a specified image does not exist")
	image_rmCmd.Flags().Bool("no-prune", false, "Do not remove dangling images")
	imageCmd.AddCommand(image_rmCmd)
}
