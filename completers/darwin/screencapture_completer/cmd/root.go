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

	rootCmd.Flags().BoolS("C", "C", false, "Capture the cursor as well as the screen")
	rootCmd.Flags().BoolS("M", "M", false, "Screen capture output will go to a new Mail message")
	rootCmd.Flags().BoolS("P", "P", false, "Screen capture output will open in Preview or QuickTime Player")
	rootCmd.Flags().BoolS("S", "S", false, "In window capture mode, capture the screen not the window")
	rootCmd.Flags().StringS("T", "T", "", "Take the picture after a delay of <seconds>, default is 5")
	rootCmd.Flags().BoolS("U", "U", false, "Capture the screen in Touch Bar mode")
	rootCmd.Flags().BoolS("W", "W", false, "Start interaction in window selection mode")
	rootCmd.Flags().StringP("bundleid", "B", "", "Open screen capture output in app with bundle id")
	rootCmd.Flags().BoolP("clipboard", "c", false, "Force screen capture to go to the clipboard")
	rootCmd.Flags().BoolP("defaults", "p", false, "Use default settings for capture, ignore files argument")
	rootCmd.Flags().StringP("display", "D", "", "Capture from the specified display (1=main, 2=secondary)")
	rootCmd.Flags().BoolP("display-errors", "d", false, "Display errors to the user graphically")
	rootCmd.Flags().StringP("format", "t", "", "Image format: png (default), pdf, jpg, tiff, gif")
	rootCmd.Flags().BoolP("interactive", "i", false, "Capture screen interactively, by selection or window")
	rootCmd.Flags().BoolS("m", "m", false, "Only capture the main monitor")
	rootCmd.Flags().BoolP("no-attached", "a", false, "Do not capture windows attached to selected window")
	rootCmd.Flags().BoolP("no-sound", "x", false, "Do not play sounds")
	rootCmd.Flags().BoolP("open", "o", false, "In window capture mode, do not capture the shadow of the window")
	rootCmd.Flags().StringP("region", "R", "", "Capture a region of the screen (x,y,width,height)")
	rootCmd.Flags().BoolP("screen", "s", false, "Only allow mouse selection mode")
	rootCmd.Flags().StringP("style", "J", "", "Set the starting style of interactive capture (selection, window, video)")
	rootCmd.Flags().BoolP("touchbar", "b", false, "Capture Touch Bar (non-interactive modes only)")
	rootCmd.Flags().StringP("video", "V", "", "Capture a video recording for <seconds>")
	rootCmd.Flags().BoolP("window", "w", false, "Only allow window selection mode")
	rootCmd.Flags().StringP("windowid", "l", "", "Capture a specific window by window id")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"bundleid": carapace.ActionValues(),
		"display":  carapace.ActionValues("1", "2", "3"),
		"format":   carapace.ActionValues("png", "pdf", "jpg", "tiff", "gif"),
		"style":    carapace.ActionValues("selection", "window", "video"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
