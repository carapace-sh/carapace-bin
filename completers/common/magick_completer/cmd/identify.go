package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/magick_completer/cmd/action"
	"github.com/spf13/cobra"
)

var identifyCmd = &cobra.Command{
	Use:   "identify",
	Short: "describe the format and characteristics of one or more image files.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identifyCmd).Standalone()

	identifyCmd.Flags().CountS("alpha", "alpha", "on, activate, off, deactivate, set, opaque, copy\", transparent, extract, background, or shape the alpha channel")
	identifyCmd.Flags().CountS("antialias", "antialias", "remove pixel-aliasing")
	identifyCmd.Flags().StringSliceS("authenticate", "authenticate", nil, "decrypt image with this password")
	identifyCmd.Flags().CountS("auto-orient", "auto-orient", "automagically orient image")
	identifyCmd.Flags().StringSliceS("channel", "channel", nil, "apply option to select image channels")
	identifyCmd.Flags().CountS("clip", "clip", "clip along the first path from the 8BIM profile")
	identifyCmd.Flags().CountS("clip-mask", "clip-mask", "associate clip mask with the image")
	identifyCmd.Flags().StringSliceS("clip-path", "clip-path", nil, "clip along a named path from the 8BIM profile")
	identifyCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "set image colorspace")
	identifyCmd.Flags().StringSliceS("crop", "crop", nil, "crop the image")
	identifyCmd.Flags().StringSliceS("debug", "debug", nil, "display copious debugging information")
	identifyCmd.Flags().StringSliceS("define", "define", nil, "define one or more image format options")
	identifyCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	identifyCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	identifyCmd.Flags().StringSliceS("endian", "endian", nil, "endianness (MSB or LSB) of the image")
	identifyCmd.Flags().StringSliceS("extract", "extract", nil, "extract area from image")
	identifyCmd.Flags().StringSliceS("features", "features", nil, "analyze image features (e.g. contract, correlations, etc.).")
	identifyCmd.Flags().StringSliceS("format", "format", nil, "output formatted image characteristics")
	identifyCmd.Flags().StringSliceS("gamma", "gamma", nil, "level of gamma correction")
	identifyCmd.Flags().StringSliceS("grayscale", "grayscale", nil, "convert image to grayscale")
	identifyCmd.Flags().BoolS("help", "help", false, "print program options")
	identifyCmd.Flags().StringSliceS("interlace", "interlace", nil, "type of image interlacing scheme")
	identifyCmd.Flags().StringSliceS("interpolate", "interpolate", nil, "pixel color interpolation method")
	identifyCmd.Flags().StringSliceS("limit", "limit", nil, "pixel cache resource limit")
	identifyCmd.Flags().StringSliceS("list", "list", nil, "Color, Configure, Delegate, Format, Magic, Module, Resource, or Type")
	identifyCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	identifyCmd.Flags().StringSliceS("mask", "mask", nil, "associate a mask with the image")
	identifyCmd.Flags().CountS("moments", "moments", "display image moments and perceptual hash.")
	identifyCmd.Flags().CountS("monitor", "monitor", "monitor progress")
	identifyCmd.Flags().CountS("negate", "negate", "replace each pixel with its complementary color")
	identifyCmd.Flags().CountS("ping", "ping", "by default, efficiently determine certain image characteristics by only reading the requisite image metadata.  To accurately identify all the image metadata and pixel characteristics, use")
	identifyCmd.Flags().StringSliceS("precision", "precision", nil, "set the maximum number of significant digits to be printed")
	identifyCmd.Flags().CountS("quiet", "quiet", "suppress all warning messages")
	identifyCmd.Flags().CountS("regard-warnings", "regard-warnings", "pay attention to warning messages.")
	identifyCmd.Flags().CountS("respect-parentheses", "respect-parentheses", "settings remain in effect until parenthesis boundary.")
	identifyCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factor")
	identifyCmd.Flags().StringSliceS("set", "set", nil, "set an image attribute")
	identifyCmd.Flags().StringSliceS("size", "size", nil, "width and height of image")
	identifyCmd.Flags().CountS("strip", "strip", "strip image of all profiles and comments")
	identifyCmd.Flags().CountS("unique", "unique", "display image the number of unique colors in the image.")
	identifyCmd.Flags().StringSliceS("units", "units", nil, "the units of image resolution")
	identifyCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	identifyCmd.Flags().BoolS("version", "version", false, "print version information")
	identifyCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", nil, "access method for pixels outside the boundaries of the image")
	rootCmd.AddCommand(identifyCmd)

	carapace.Gen(identifyCmd).FlagCompletion(carapace.ActionMap{
		"channel":       action.ActionCompleteOption("Channel").UniqueList(","),
		"clip-mask":     carapace.ActionFiles(),
		"colorspace":    action.ActionCompleteOption("Colorspace"),
		"debug":         action.ActionCompleteOption("Debug").UniqueList(","),
		"endian":        action.ActionCompleteOption("Endian"),
		"grayscale":     action.ActionCompleteOption("Intensity"),
		"interlace":     action.ActionCompleteOption("Interlace"),
		"interpolate":   action.ActionCompleteOption("Interpolate"),
		"list":          action.ActionCompleteOption("List"),
		"mask":          carapace.ActionFiles(),
		"units":         action.ActionCompleteOption("Units"),
		"virtual-pixel": action.ActionCompleteOption("VirtualPixel"),
	})

	carapace.Gen(identifyCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
