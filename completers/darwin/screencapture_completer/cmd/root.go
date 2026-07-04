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

	rootCmd.Flags().BoolS("C", "C", false, "Capture the cursor")
	rootCmd.Flags().BoolS("M", "M", false, "Capture the menu bar")
	rootCmd.Flags().BoolS("P", "P", false, "Capture the screen to the pasteboard")
	rootCmd.Flags().BoolS("S", "S", false, "Capture the screen (not a window)")
	rootCmd.Flags().StringS("T", "T", "", "Capture the whole screen after the specified delay (in seconds)")
	rootCmd.Flags().BoolS("W", "W", false, "Capture a window interactively")
	rootCmd.Flags().BoolP("clipboard", "c", false, "Force screen capture to go to the clipboard")
	rootCmd.Flags().StringP("delay", "d", "", "Set delay before capture (in seconds)")
	rootCmd.Flags().StringP("format", "t", "", "Set image format (png, jpg, tiff, pdf, gif)")
	rootCmd.Flags().BoolP("interactive", "i", false, "Capture screen interactively")
	rootCmd.Flags().BoolS("m", "m", false, "Do not capture the mouse cursor")
	rootCmd.Flags().BoolP("no-sound", "x", false, "Do not play sound")
	rootCmd.Flags().BoolP("open", "o", false, "Open the captured image in Preview")
	rootCmd.Flags().BoolP("screen", "s", false, "Capture the screen only")
	rootCmd.Flags().BoolP("window", "w", false, "Capture a window interactively")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("png", "jpg", "tiff", "pdf", "gif"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
