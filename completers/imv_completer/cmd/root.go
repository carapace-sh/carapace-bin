package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "imv",
	Short: "Image viewer for X11 and Wayland",
	Long:  "https://git.sr.ht/~exec64/imv/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("b", "b", "", "Set the background colour")
	rootCmd.Flags().StringS("c", "c", "", "Specify a command to be run on launch")
	rootCmd.Flags().BoolS("d", "d", false, "Start with overlay visible")
	rootCmd.Flags().BoolS("f", "f", false, "Start fullscreen")
	rootCmd.Flags().BoolS("h", "h", false, "Show help message and quit")
	rootCmd.Flags().BoolS("l", "l", false, "List open files to stdout at exit")
	rootCmd.Flags().StringS("n", "n", "", "Start with the given path, or index selected")
	rootCmd.Flags().BoolS("r", "r", false, "Load directories recursively")
	rootCmd.Flags().StringS("s", "s", "", "Set scaling mode to use")
	rootCmd.Flags().StringS("t", "t", "", "Start in slideshow mode, with each image shown for the given number of seconds")
	rootCmd.Flags().StringS("u", "u", "", "Set upscaling method used by imv")
	rootCmd.Flags().BoolS("v", "v", false, "Show version and quit")
	rootCmd.Flags().BoolS("x", "x", false, "Disable looping of input paths")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"b": carapace.Batch(
			color.ActionHexColors(),
			carapace.ActionValuesDescribed("checks", "chequered background"),
		).ToA(),
		"n": carapace.ActionFiles(),
		"s": carapace.ActionValues("none", "shrink", "full", "crop"),
		"u": carapace.ActionValues("linear", "nearest_neighbour"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
