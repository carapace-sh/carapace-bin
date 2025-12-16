package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/sway"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wl-mirror",
	Short: "a simple Wayland output mirror client",
	Long:  "https://github.com/Ferdi265/wl-mirror",
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

	rootCmd.Flags().StringP("backend", "b", "", "use a specific backend for capturing the screen")
	rootCmd.Flags().BoolP("freeze", "f", false, "freeze the current image on the screen")
	rootCmd.Flags().BoolP("fullscreen", "F", false, "open wl-mirror as fullscreen")
	rootCmd.Flags().String("fullscreen-output", "", "open wl-mirror as fullscreen on output O")
	rootCmd.Flags().BoolP("help", "h", false, "show this help")
	rootCmd.Flags().BoolP("invert-colors", "i", false, "invert colors in the mirrored screen")
	rootCmd.Flags().Bool("no-fullscreen", false, "open wl-mirror as a window (default)")
	rootCmd.Flags().Bool("no-fullscreen-output", false, "open wl-mirror as fullscreen on the current output (default)")
	rootCmd.Flags().Bool("no-invert-colors", false, "don't invert colors in the mirrored screen (default)")
	rootCmd.Flags().Bool("no-region", false, "capture the entire output (default)")
	rootCmd.Flags().Bool("no-show-cursor", false, "don't show the cursor on the mirrored screen")
	rootCmd.Flags().Bool("no-verbose", false, "disable debug logging (default)")
	rootCmd.Flags().StringP("region", "r", "", "capture custom region R")
	rootCmd.Flags().StringSliceP("scaling", "s", nil, "scaling mode")
	rootCmd.Flags().BoolP("show-cursor", "c", false, "show the cursor on the mirrored screen (default)")
	rootCmd.Flags().BoolP("stream", "S", false, "accept a stream of additional options on stdin")
	rootCmd.Flags().Bool("toggle-freeze", false, "toggle freeze state of screen capture")
	rootCmd.Flags().StringP("transform", "t", "", "apply custom transform T")
	rootCmd.Flags().Bool("unfreeze", false, "resume the screen capture after a freeze")
	rootCmd.Flags().BoolP("verbose", "v", false, "enable debug logging")
	rootCmd.Flags().BoolP("version", "V", false, "print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"backend": carapace.ActionValuesDescribed(
			"auto", "automatically try the backends in order and use the first that works",
			"dmabuf", "use the wlr-export-dmabuf-unstable-v1 protocol to capture outputs",
			"screencopy", "use the wlr-screencopy-unstable-v1 protocol to capture outputs",
		),
		"fullscreen-output": carapace.ActionValues(), // TODO
		"scaling": carapace.ActionValuesDescribed(
			"fit", "scale to fit (default)",
			"cover", "scale to cover, cropping if needed",
			"exact", "only scale to exact multiples of the output size",
			"linear", "use linear scaling (default)",
			"nearest", "use nearest neighbor scaling",
		),
		"transform": carapace.ActionValuesDescribed(
			"normal", "no transformation",
			"flipX", "flip the X coordinate",
			"flipY", "flip the Y coordinate",
			"0cw", "apply a clockwise rotation",
			"90cw", "apply a clockwise rotation",
			"180cw", "apply a clockwise rotation",
			"270cw", "apply a clockwise rotation",
			"0ccw", "apply a counter-clockwise rotation",
			"90ccw", "apply a counter-clockwise rotation",
			"180ccw", "apply a counter-clockwise rotation",
			"270ccw", "apply a counter-clockwise rotation",

			"flipped", "flip the X coordinate",
			"0", "apply a clockwise rotation",
			"90", "apply a clockwise rotation",
			"180", "apply a clockwise rotation",
			"270", "apply a clockwise rotation",
		).UniqueList("-"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		sway.ActionOutputs(), // TODO other sources
	)
}
