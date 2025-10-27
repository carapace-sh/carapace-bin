package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_saveCmd = &cobra.Command{
	Use:   "save IMAGE [ARCHIVE | -]",
	Short: "Save a image from minikube",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_saveCmd).Standalone()

	image_saveCmd.Flags().Bool("daemon", false, "Cache image to docker daemon")
	image_saveCmd.Flags().Bool("remote", false, "Cache image to remote registry")
	imageCmd.AddCommand(image_saveCmd)
}
