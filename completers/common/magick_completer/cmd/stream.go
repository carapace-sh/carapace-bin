package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/magick_completer/cmd/action"
	"github.com/spf13/cobra"
)

var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "a lightweight tool to stream one or more pixel components of the image or portion of the image to your choice of storage formats. It writes the pixel components as they are read from the input image a row at a time making stream desirable when working with large images or when you require raw pixel components.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(streamCmd).Standalone()

	streamCmd.Flags().StringSliceS("authenticate", "authenticate", nil, "decrypt image with this password")
	streamCmd.Flags().StringSliceS("channel", "channel", nil, "apply option to select image channels")
	streamCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "set image colorspace")
	streamCmd.Flags().StringSliceS("debug", "debug", nil, "display copious debugging information")
	streamCmd.Flags().StringSliceS("define", "define", nil, "define one or more image format options")
	streamCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	streamCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	streamCmd.Flags().StringSliceS("extract", "extract", nil, "extract area from image")
	streamCmd.Flags().BoolS("help", "help", false, "print program options")
	streamCmd.Flags().StringSliceS("interlace", "interlace", nil, "type of image interlacing scheme")
	streamCmd.Flags().StringSliceS("interpolate", "interpolate", nil, "pixel color interpolation method")
	streamCmd.Flags().StringSliceS("limit", "limit", nil, "pixel cache resource limit")
	streamCmd.Flags().StringSliceS("list", "list", nil, "Color, Configure, Delegate, Format, Magic, Module, Resource, or Type")
	streamCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	streamCmd.Flags().StringSliceS("map", "map", nil, "store pixels in this format.")
	streamCmd.Flags().CountS("monitor", "monitor", "monitor progress")
	streamCmd.Flags().StringSliceS("quantize", "quantize", nil, "reduce image colors in this colorspace")
	streamCmd.Flags().CountS("quiet", "quiet", "suppress all warning messages")
	streamCmd.Flags().CountS("regard-warnings", "regard-warnings", "pay attention to warning messages.")
	streamCmd.Flags().CountS("respect-parentheses", "respect-parentheses", "settings remain in effect until parenthesis boundary.")
	streamCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factor")
	streamCmd.Flags().StringSliceS("seed", "seed", nil, "seed a new sequence of pseudo-random numbers")
	streamCmd.Flags().StringSliceS("set", "set", nil, "set an image attribute")
	streamCmd.Flags().StringSliceS("size", "size", nil, "width and height of image")
	streamCmd.Flags().StringSliceS("storage-type", "storage-type", nil, "store pixels with this storage type.")
	streamCmd.Flags().CountS("synchronize", "synchronize", "synchronize image to storage device")
	streamCmd.Flags().CountS("taint", "taint", "mark the image as modified")
	streamCmd.Flags().StringSliceS("transparent-color", "transparent-color", nil, "transparent color")
	streamCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	streamCmd.Flags().BoolS("version", "version", false, "print version information")
	streamCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", nil, "access method for pixels outside the boundaries of the image")
	rootCmd.AddCommand(streamCmd)

	carapace.Gen(streamCmd).FlagCompletion(carapace.ActionMap{
		"channel":           action.ActionCompleteOption("Channel").UniqueList(","),
		"colorspace":        action.ActionCompleteOption("Colorspace"),
		"debug":             action.ActionCompleteOption("Debug").UniqueList(","),
		"interlace":         action.ActionCompleteOption("Interlace"),
		"interpolate":       action.ActionCompleteOption("Interpolate"),
		"list":              action.ActionCompleteOption("List"),
		"quantize":          action.ActionCompleteOption("Colorspace"),
		"storage-type":      action.ActionCompleteOption("Storage"),
		"transparent-color": action.ActionColors(),
		"virtual-pixel":     action.ActionCompleteOption("VirtualPixel"),
	})

	carapace.Gen(streamCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
