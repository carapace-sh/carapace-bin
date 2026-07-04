package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sips",
	Short: "scriptable image processing system",
	Long:  "https://keith.github.io/xcode-manpages/sips.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("cropToHeightWidth", "c", "", "Crop to specified height and width")
	rootCmd.Flags().BoolP("deleteProperty", "d", false, "Delete the specified property")
	rootCmd.Flags().StringP("format", "s", "", "Set output format")
	rootCmd.Flags().BoolP("getProperty", "g", false, "Get the specified property")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().StringP("out", "o", "", "Set output file")
	rootCmd.Flags().StringP("resampleHeight", "Z", "", "Resample to specified height")
	rootCmd.Flags().StringP("resampleHeightWidth", "z", "", "Resample to specified height and width")
	rootCmd.Flags().StringP("rotate", "r", "", "Rotate by specified degrees")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("jpeg", "tiff", "png", "gif", "jp2", "pict", "bmp", "qtif", "psd", "sgi", "tga"),
		"out":    carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
