package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "viu",
	Short: "View images right from the terminal",
	Long:  "https://github.com/atanunq/viu",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("blocks", "b", false, "Force block output")
	rootCmd.Flags().StringP("frame-rate", "f", "", "Play gif at the given frame rate")
	rootCmd.Flags().StringP("height", "h", "", "Resize the image to a provided height")
	rootCmd.Flags().Bool("help", false, "Prints help information")
	rootCmd.Flags().BoolP("mirror", "m", false, "Display a mirror of the original image")
	rootCmd.Flags().BoolP("name", "n", false, "Output the name of the file before displaying")
	rootCmd.Flags().BoolP("once", "1", false, "Only loop once through the animation")
	rootCmd.Flags().BoolP("recursive", "r", false, "Recurse down directories if passed one")
	rootCmd.Flags().BoolP("static", "s", false, "Show only first frame of gif")
	rootCmd.Flags().BoolP("transparent", "t", false, "Display transparent image with transparent background")
	rootCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.Flags().StringP("width", "w", "", "Resize the image to a provided width")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
