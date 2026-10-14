package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_stageCmd = &cobra.Command{
	Use:   "stage",
	Short: "Upload data for a specific image to staging.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_stageCmd).Standalone()

	image_stageCmd.Flags().String("file", "", "Local file that contains disk image to be uploaded.")
	image_stageCmd.Flags().Bool("progress", false, "Show upload progress bar (ignored if passing data via stdin)")
	imageCmd.AddCommand(image_stageCmd)
}
