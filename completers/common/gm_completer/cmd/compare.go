package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "compare two images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compareCmd).Standalone()

	compareCmd.Flags().StringSliceS("authenticate", "authenticate", nil, "decrypt image with this password")
	compareCmd.Flags().CountS("auto-orient", "auto-orient", "orient (rotate) images so they are upright")
	compareCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "alternate image colorspace")
	compareCmd.Flags().StringSliceS("compress", "compress", nil, "image compression type")
	compareCmd.Flags().StringSliceS("debug", "debug", nil, "display copious debugging information")
	compareCmd.Flags().StringSliceS("define", "define", nil, "coder/decoder specific options")
	compareCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	compareCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	compareCmd.Flags().StringSliceS("display", "display", nil, "get image or font from this X server")
	compareCmd.Flags().StringSliceS("endian", "endian", nil, "multibyte word order (LSB, MSB, or Native)")
	compareCmd.Flags().StringSliceS("file", "file", nil, "write difference image to this file")
	compareCmd.Flags().BoolS("help", "help", false, "print program options")
	compareCmd.Flags().StringSliceS("highlight-color", "highlight-color", nil, "color to use when annotating difference pixels")
	compareCmd.Flags().StringSliceS("highlight-style", "highlight-style", nil, "pixel highlight style (assign, threshold, tint, xor)")
	compareCmd.Flags().StringSliceS("interlace", "interlace", nil, "None, Line, Plane, or Partition")
	compareCmd.Flags().StringSliceS("limit", "limit", nil, "Disk, File, Map, Memory, Pixels, Width, Height, Threads, Read, or Write resource limit")
	compareCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	compareCmd.Flags().CountS("matte", "matte", "store matte channel if the image has one")
	compareCmd.Flags().CountS("maximum-error", "maximum-error", "maximum total difference before returning error")
	compareCmd.Flags().CountS("metric", "metric", "comparison metric (MAE, MSE, PAE, PSNR, RMSE)")
	compareCmd.Flags().CountS("monitor", "monitor", "show progress indication")
	compareCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factors")
	compareCmd.Flags().StringSliceS("size", "size", nil, "width and height of image")
	compareCmd.Flags().StringSliceS("type", "type", nil, "image type")
	compareCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	compareCmd.Flags().BoolS("version", "version", false, "print version information")
	rootCmd.AddCommand(compareCmd)

	carapace.Gen(compareCmd).FlagCompletion(carapace.ActionMap{
		"file":            carapace.ActionFiles(),
		"highlight-color": action.ActionColor(),
	})

	carapace.Gen(compareCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
