package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "screencapture",
	Short: "capture images from the screen",
	Long:  "https://keith.github.io/xcode-manpages/screencapture.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("clipboard", "c", false, "Force screen capture to go to the clipboard")
	rootCmd.Flags().BoolP("cursor", "C", false, "Capture the cursor")
	rootCmd.Flags().StringP("delay", "d", "", "Set delay before capture")
	rootCmd.Flags().BoolP("display", "D", false, "Capture the display")
	rootCmd.Flags().StringP("format", "t", "", "Set image format")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("interactive", "i", false, "Capture screen interactively")
	rootCmd.Flags().BoolP("menu", "M", false, "Capture menu bar")
	rootCmd.Flags().BoolP("mouse", "m", false, "Capture mouse cursor")
	rootCmd.Flags().BoolP("no-sound", "x", false, "Do not play sound")
	rootCmd.Flags().StringP("output", "o", "", "Set output file")
	rootCmd.Flags().StringP("quality", "q", "", "Set image quality")
	rootCmd.Flags().BoolP("screen", "S", false, "Capture screen")
	rootCmd.Flags().StringP("type", "T", "", "Set capture type")
	rootCmd.Flags().BoolP("window", "w", false, "Capture window interactively")
	rootCmd.Flags().StringP("window-id", "l", "", "Capture the window with the specified ID")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("png", "jpg", "tiff", "pdf", "gif"),
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
