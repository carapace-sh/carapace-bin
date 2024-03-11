package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imgcatCmd = &cobra.Command{
	Use:   "imgcat [OPTIONS] [FILE",
	Short: "Output an image to the terminal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imgcatCmd).Standalone()

	imgcatCmd.Flags().String("height", "", "Specify the display height")
	imgcatCmd.Flags().BoolP("help", "h", false, "Print help")
	imgcatCmd.Flags().Bool("no-preserve-aspect-ratio", false, "Do not respect the aspect ratio")
	imgcatCmd.Flags().String("width", "", "Specify the display width")
	rootCmd.AddCommand(imgcatCmd)

	carapace.Gen(imgcatCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
