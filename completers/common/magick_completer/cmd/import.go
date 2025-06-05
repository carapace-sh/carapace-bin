package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/magick_completer/cmd/action"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "save any visible window on an X server and outputs it as an image file. You can capture a single window, the entire screen, or any rectangular portion of the screen.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().CountS("adjoin", "adjoin", "join images into a single multi-image file")
	importCmd.Flags().StringSliceS("annotate", "annotate", nil, "annotate the image with text")
	importCmd.Flags().CountS("border", "border", "include window border in the output image")
	importCmd.Flags().StringSliceS("channel", "channel", nil, "apply option to select image channels")
	importCmd.Flags().StringSliceS("colors", "colors", nil, "preferred number of colors in the image")
	importCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "set image colorspace")
	importCmd.Flags().StringSliceS("comment", "comment", nil, "annotate image with comment")
	importCmd.Flags().StringSliceS("compress", "compress", nil, "image compression type")
	importCmd.Flags().CountS("contrast", "contrast", "enhance or reduce the image contrast")
	importCmd.Flags().StringSliceS("crop", "crop", nil, "preferred size and location of the cropped image")
	importCmd.Flags().StringSliceS("debug", "debug", nil, "import copious debugging information")
	importCmd.Flags().StringSliceS("define", "define", nil, "define one or more image format options")
	importCmd.Flags().StringSliceS("delay", "delay", nil, "import the next image after pausing")
	importCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	importCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	importCmd.Flags().CountS("descend", "descend", "obtain image by descending window hierarchy")
	importCmd.Flags().StringSliceS("display", "display", nil, "get image or font from this X server")
	importCmd.Flags().StringSliceS("dispose", "dispose", nil, "layer disposal method")
	importCmd.Flags().StringSliceS("dither", "dither", nil, "apply error diffusion to image")
	importCmd.Flags().StringSliceS("encipher", "encipher", nil, "convert plain pixels to cipher pixels")
	importCmd.Flags().StringSliceS("encoding", "encoding", nil, "text encoding type")
	importCmd.Flags().StringSliceS("endian", "endian", nil, "endianness (MSB or LSB) of the image")
	importCmd.Flags().StringSliceS("filter", "filter", nil, "use this filter when resizing an image")
	importCmd.Flags().CountS("frame", "frame", "include window manager frame")
	importCmd.Flags().StringSliceS("geometry", "geometry", nil, "preferred size or location of the image")
	importCmd.Flags().StringSliceS("gravity", "gravity", nil, "horizontal and vertical text placement")
	importCmd.Flags().BoolS("help", "help", false, "print program options")
	importCmd.Flags().CountS("identify", "identify", "identify the format and characteristics of the image")
	importCmd.Flags().StringSliceS("interlace", "interlace", nil, "type of image interlacing scheme")
	importCmd.Flags().StringSliceS("interpolate", "interpolate", nil, "pixel color interpolation method")
	importCmd.Flags().StringSliceS("label", "label", nil, "assign a label to an image")
	importCmd.Flags().StringSliceS("limit", "limit", nil, "pixel cache resource limit")
	importCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	importCmd.Flags().CountS("monitor", "monitor", "monitor progress")
	importCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	importCmd.Flags().CountS("negate", "negate", "replace each pixel with its complementary color")
	importCmd.Flags().StringSliceS("page", "page", nil, "size and location of an image canvas (setting)")
	importCmd.Flags().StringSliceS("pause", "pause", nil, "seconds delay between snapshots")
	importCmd.Flags().StringSliceS("quality", "quality", nil, "JPEG/MIFF/PNG compression level")
	importCmd.Flags().StringSliceS("quantize", "quantize", nil, "reduce image colors in this colorspace")
	importCmd.Flags().CountS("quiet", "quiet", "suppress all warning messages")
	importCmd.Flags().CountS("regard-warnings", "regard-warnings", "pay attention to warning messages.")
	importCmd.Flags().StringSliceS("repage", "repage", nil, "size and location of an image canvas")
	importCmd.Flags().StringSliceS("resize", "resize", nil, "resize the image")
	importCmd.Flags().CountS("respect-parentheses", "respect-parentheses", "settings remain in effect until parenthesis boundary.")
	importCmd.Flags().StringSliceS("rotate", "rotate", nil, "apply Paeth rotation to the image")
	importCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factor")
	importCmd.Flags().StringSliceS("scene", "scene", nil, "image scene number")
	importCmd.Flags().CountS("screen", "screen", "select image from root window")
	importCmd.Flags().StringSliceS("seed", "seed", nil, "seed a new sequence of pseudo-random numbers")
	importCmd.Flags().StringSliceS("set", "set", nil, "set an image attribute")
	importCmd.Flags().CountS("silent", "silent", "operate silently, i.e. don't ring any bells")
	importCmd.Flags().CountS("strip", "strip", "strip image of all profiles and comments")
	importCmd.Flags().CountS("synchronize", "synchronize", "synchronize image to storage device")
	importCmd.Flags().CountS("taint", "taint", "mark the image as modified")
	importCmd.Flags().StringSliceS("transparent-color", "transparent-color", nil, "transparent color")
	importCmd.Flags().CountS("trim", "trim", "trim image edges")
	importCmd.Flags().StringSliceS("type", "type", nil, "image type")
	importCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	importCmd.Flags().BoolS("version", "version", false, "print version information")
	importCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", nil, "access method for pixels outside the boundaries of the image")
	importCmd.Flags().StringSliceS("window", "window", nil, "select window with this id or name")
	rootCmd.AddCommand(importCmd)

	carapace.Gen(importCmd).FlagCompletion(carapace.ActionMap{
		"channel":           action.ActionCompleteOption("Channel").UniqueList(","),
		"colorspace":        action.ActionCompleteOption("Colorspace"),
		"compress":          action.ActionCompleteOption("Compress"),
		"debug":             action.ActionCompleteOption("Debug").UniqueList(","),
		"dispose":           action.ActionCompleteOption("Dispose"),
		"dither":            action.ActionCompleteOption("Dither"),
		"encipher":          carapace.ActionFiles(),
		"encoding":          carapace.ActionValues("AdobeCustom", "AdobeExpert", "AdobeStandard", "AppleRoman", "BIG5", "GB2312", "Latin 2", "None", "SJIScode", "Symbol", "Unicode", "Wansung"),
		"endian":            action.ActionCompleteOption("Endian"),
		"filter":            action.ActionCompleteOption("Filter"),
		"gravity":           action.ActionCompleteOption("Gravity"),
		"interlace":         action.ActionCompleteOption("Interlace"),
		"interpolate":       action.ActionCompleteOption("Interpolate"),
		"quantize":          action.ActionCompleteOption("Colorspace"),
		"transparent-color": action.ActionColors(),
		"type":              action.ActionCompleteOption("Type"),
		"virtual-pixel":     action.ActionCompleteOption("VirtualPixel"),
	})

	carapace.Gen(importCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
