package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var icatCmd = &cobra.Command{
	Use:   "icat",
	Short: "Display images in the terminal",
}

func init() {
	rootCmd.AddCommand(icatCmd)
	carapace.Gen(icatCmd).Standalone()

	icatCmd.Flags().String("align", "", "Horizontal alignment for the displayed image.")
	icatCmd.Flags().String("place", "", "Choose where on the screen to display the image. The image will be scaled to fit into the specified rectangle. The syntax for specifying rectangles is <width>x<height>@<left>x<top>. All measurements are in cells (i.e. cursor positions) with the origin (0, 0) at the top-left corner of the screen. Note that the --align option will horizontally align the image within this rectangle. By default, the image is horizontally centered within the rectangle. Using place will cause the cursor to be positioned at the top left corner of the image, instead of on the line after the image.")
	icatCmd.Flags().Bool("scale-up", false, "Cause images that are smaller than the specified area to be scaled up to use as much of the specified area as possible. The specified area depends on either the --place or the --fit options.")
	icatCmd.Flags().String("fit", "", "When not using --place, control how the image is scaled relative to the screen. You can have it fit in the screen width or height or both or neither.")
	icatCmd.Flags().String("background", "", "Specify a background color, this will cause transparent images to be")
	icatCmd.Flags().String("mirror", "", "Mirror the image about a horizontal or vertical axis or both.")
	icatCmd.Flags().Bool("clear", false, "Remove all images currently displayed on the screen. Note that this cannot work with terminal multiplexers such as tmux since only the multiplexer can know the position of the screen.")
	icatCmd.Flags().Bool("clear-all", false, "Remove all images from screen and scrollback. Note that with terminal multiplexers like tmux, this will move images from all panes.")
	icatCmd.Flags().String("transfer-mode", "", "Which mechanism to use to transfer images to the terminal. The default is to auto-detect. file means to use a temporary file, memory means to use shared memory, stream means to send the data via terminal escape codes. Note that if you use the file or memory transfer modes and you are connecting over a remote session then image display will not work.")
	icatCmd.Flags().Bool("detect-support", false, "Detect support for image display in the terminal. If not supported, will exit with exit code 1, otherwise will exit with code 0 and print the supported transfer mode to stderr, which can be used with the --transfer-mode option.")
	icatCmd.Flags().Float64("detection-timeout", 0, "The amount of time (in seconds) to wait for a response from the terminal, when detecting image display support.")
	icatCmd.Flags().String("use-window-size", "", "Instead of querying the terminal for the window size, use the specified size, which must be of the format: width_in_cells,height_in_cells,width_in_pixels,height_in_pixels")
	icatCmd.Flags().Bool("print-window-size", false, "Print out the window size as <width>x<height> (in pixels) and quit. This is a convenience method to query the window size if using kitten icat from a scripting language that cannot make termios calls.")
	icatCmd.Flags().String("stdin", "", "Read image data from STDIN. The default is to do it automatically, when STDIN is not a terminal, but you can turn it off or on explicitly, if needed.")
	icatCmd.Flags().Bool("silent", false, "Not used, present for legacy compatibility.")
	icatCmd.Flags().String("engine", "", "The engine used for decoding and processing of images. The default is to use the most appropriate engine. The builtin engine uses Go's native imaging libraries. The magick engine uses ImageMagick which requires it to be installed on the system.")
	icatCmd.Flags().Float64P("z-index", "z", 0, "Z-index of the image. When negative, text will be displayed on top of the image. Use a double minus for values under the threshold for drawing images under cell background colors. For example, --1 evaluates as -1,073,741,825.")
	icatCmd.Flags().Float64P("loop", "l", 0, "Number of times to loop animations. Negative values loop forever. Zero means only the first frame of the animation is displayed. Otherwise, the animation is looped the specified number of times.")
	icatCmd.Flags().Bool("hold", false, "Wait for a key press before exiting after displaying the images.")
	icatCmd.Flags().Bool("unicode-placeholder", false, "Use the Unicode placeholder method to display the images. Useful to display images from within full screen terminal programs that do not understand the kitty graphics protocol such as multiplexers or editors. See graphics_unicode_placeholders for details. Note that when using this method, images placed (with --place) that do not fit on the screen, will get wrapped at the screen edge instead of getting truncated. This wrapping is per line and therefore the image will look like it is interleaved with blank lines.")
	icatCmd.Flags().String("passthrough", "", "Whether to surround graphics commands with escape sequences that allow them to passthrough programs like tmux. The default is to detect when running inside tmux and automatically use the tmux passthrough escape codes. Note that when this option is enabled it implies --unicode-placeholder as well.")
	icatCmd.Flags().Float64("image-id", 0, "The graphics protocol id to use for the created image. Normally, a random id is created if needed. This option allows control of the id. When multiple images are sent, sequential ids starting from the specified id are used. Valid ids are from 1 to 4294967295. Numbers outside this range are automatically wrapped.")
	icatCmd.Flags().BoolP("no-trailing-newline", "n", false, "By default, the cursor is moved to the next line after displaying an image. This option, prevents that. Should not be used when catting multiple images. Also has no effect when the --place option is used.")

	carapace.Gen(icatCmd).FlagCompletion(carapace.ActionMap{
		"align":         carapace.ActionValues("center", "left", "right"),
		"fit":           carapace.ActionValues("width", "height", "both", "none"),
		"mirror":        carapace.ActionValues("none", "both", "horizontal", "vertical"),
		"transfer-mode": carapace.ActionValues("detect", "file", "memory", "stream"),
		"stdin":         carapace.ActionValues("detect", "no", "yes"),
		"engine":        carapace.ActionValues("auto", "builtin", "magick"),
		"passthrough":   carapace.ActionValues("detect", "none", "tmux"),
	})
	carapace.Gen(icatCmd).PositionalAnyCompletion(
		carapace.Batch(carapace.ActionFiles(), carapace.ActionDirectories()).ToA(),
	)
}
