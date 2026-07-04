package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scrot",
	Short: "command line screen capture utility",
	Long:  "https://en.wikipedia.org/wiki/Scrot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("autoselect", "a", "", "non-interactively choose a rectangle of x,y,w,h")
	rootCmd.Flags().BoolP("border", "b", false, "When selecting a window, grab wm border too")
	rootCmd.Flags().StringP("class", "C", "", "Window class name. Associative with options: -k")
	rootCmd.Flags().StringP("compression", "Z", "", "Compression level (0-9)")
	rootCmd.Flags().BoolP("count", "c", false, "show a countdown before taking the shot")
	rootCmd.Flags().StringP("delay", "d", "", "wait NUM seconds before taking a shot")
	rootCmd.Flags().StringP("display", "D", "", "Set DISPLAY target other than current")
	rootCmd.Flags().StringP("exec", "e", "", "run APP on the resulting screenshot")
	rootCmd.Flags().StringP("file", "F", "", "Specify the output file")
	rootCmd.Flags().BoolP("focused", "u", false, "use the currently focused window")
	rootCmd.Flags().StringP("format", "n", "", "Specify the output file format (e.g. png, jpg)")
	rootCmd.Flags().BoolP("freeze", "f", false, "Freeze the screen when the selection is used: --select")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("ignorekeyboard", "i", false, "Don't exit for keyboard input")
	rootCmd.Flags().StringP("line", "l", "", "Indicates the style of the line when the selection is used: --select")
	rootCmd.Flags().StringP("monitor", "M", "", "Capture Xrandr monitor number NUM")
	rootCmd.Flags().BoolP("multidisp", "m", false, "For multiple heads, grab shot from each and join them together.")
	rootCmd.Flags().BoolP("overwrite", "o", false, "By default scrot does not overwrite the files, use this option to allow it.")
	rootCmd.Flags().BoolP("pointer", "p", false, "Capture the mouse pointer.")
	rootCmd.Flags().StringP("quality", "q", "", "Image quality (1-100) high value means high size, low compression")
	rootCmd.Flags().BoolP("select", "s", false, "interactively choose a window or rectangle with the mouse")
	rootCmd.Flags().BoolP("silent", "z", false, "Prevent beeping")
	rootCmd.Flags().BoolP("stack", "k", false, "Capture stack/overlapped windows and join them together.")
	rootCmd.Flags().StringP("thumb", "t", "", "generate thumbnail too")
	rootCmd.Flags().BoolP("version", "v", false, "output version information and exit")
	rootCmd.Flags().StringP("window", "w", "", "Window identifier (decimal or hex) to capture")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"exec":   carapace.ActionFiles(),
		"file":   carapace.ActionFiles(),
		"format": carapace.ActionValues("png", "jpg", "jpeg", "gif", "bmp"),
	})
}
