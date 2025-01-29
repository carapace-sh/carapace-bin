package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_historyCmd = &cobra.Command{
	Use:   "history [options] IMAGE",
	Short: "Show history of a specified image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_historyCmd).Standalone()

	image_historyCmd.Flags().String("format", "", "Change the output to JSON or a Go template")
	image_historyCmd.Flags().BoolP("human", "H", false, "Display sizes and dates in human readable format")
	image_historyCmd.Flags().Bool("no-trunc", false, "Do not truncate the output")
	image_historyCmd.Flags().BoolP("quiet", "q", false, "Display the numeric IDs only")
	imageCmd.AddCommand(image_historyCmd)
}
