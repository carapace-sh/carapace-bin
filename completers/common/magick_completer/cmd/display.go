package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/magick_completer/cmd/action"
	"github.com/spf13/cobra"
)

var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "display an image or image sequence on any X server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(displayCmd).Standalone()

	displayCmd.Flags().CountS("alpha", "alpha", "on, activate, off, deactivate, set, opaque, copy\", transparent, extract, background, or shape the alpha channel")
	displayCmd.Flags().CountS("antialias", "antialias", "remove pixel-aliasing")
	displayCmd.Flags().StringSliceS("authenticate", "authenticate", nil, "decrypt image with this password")
	displayCmd.Flags().CountS("backdrop", "backdrop", "display image centered on a backdrop")
	displayCmd.Flags().StringSliceS("background", "background", nil, "background color")
	displayCmd.Flags().StringSliceS("black-threshold", "black-threshold", nil, "force all pixels below the threshold into black")
	displayCmd.Flags().StringSliceS("border", "border", nil, "surround image with a border of color")
	displayCmd.Flags().StringSliceS("bordercolor", "bordercolor", nil, "border color")
	displayCmd.Flags().StringSliceS("channel", "channel", nil, "apply option to select image channels")
	displayCmd.Flags().CountS("clip", "clip", "clip along the first path from the 8BIM profile")
	displayCmd.Flags().StringSliceS("clip-path", "clip-path", nil, "clip along a named path from the 8BIM profile")
	displayCmd.Flags().CountS("coalesce", "coalesce", "merge a sequence of images")
	displayCmd.Flags().StringSliceS("colormap", "colormap", nil, "Shared or Private")
	displayCmd.Flags().StringSliceS("colors", "colors", nil, "preferred number of colors in the image")
	displayCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "set image colorspace")
	displayCmd.Flags().StringSliceS("comment", "comment", nil, "annotate image with comment")
	displayCmd.Flags().StringSliceS("compress", "compress", nil, "image compression type")
	displayCmd.Flags().CountS("contrast", "contrast", "enhance or reduce the image contrast")
	displayCmd.Flags().StringSliceS("crop", "crop", nil, "preferred size and location of the cropped image")
	displayCmd.Flags().StringSliceS("debug", "debug", nil, "display copious debugging information")
	displayCmd.Flags().StringSliceS("decipher", "decipher", nil, "convert cipher pixels to plain")
	displayCmd.Flags().StringSliceS("define", "define", nil, "define one or more image format options")
	displayCmd.Flags().StringSliceS("delay", "delay", nil, "display the next image after pausing")
	displayCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	displayCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	displayCmd.Flags().CountS("despeckle", "despeckle", "reduce the speckles within an image")
	displayCmd.Flags().StringSliceS("display", "display", nil, "get image or font from this X server")
	displayCmd.Flags().StringSliceS("dispose", "dispose", nil, "layer disposal method")
	displayCmd.Flags().StringSliceS("dither", "dither", nil, "apply error diffusion to image")
	displayCmd.Flags().StringSliceS("edge", "edge", nil, "apply a filter to detect edges in the image")
	displayCmd.Flags().StringSliceS("endian", "endian", nil, "endianness (MSB or LSB) of the image")
	displayCmd.Flags().CountS("enhance", "enhance", "apply a digital filter to enhance a noisy image")
	displayCmd.Flags().CountS("equalize", "equalize", "perform histogram equalization to an image")
	displayCmd.Flags().StringSliceS("extract", "extract", nil, "extract area from image")
	displayCmd.Flags().StringSliceS("filter", "filter", nil, "use this filter when resizing an image")
	displayCmd.Flags().CountS("flatten", "flatten", "flatten a sequence of images")
	displayCmd.Flags().CountS("flip", "flip", "flip image in the vertical direction")
	displayCmd.Flags().CountS("flop", "flop", "flop image in the horizontal direction")
	displayCmd.Flags().StringSliceS("frame", "frame", nil, "surround image with an ornamental border")
	displayCmd.Flags().StringSliceS("fuzz", "fuzz", nil, "colors within this distance are considered equal")
	displayCmd.Flags().StringSliceS("gamma", "gamma", nil, "level of gamma correction")
	displayCmd.Flags().StringSliceS("geometry", "geometry", nil, "preferred size or location of the image")
	displayCmd.Flags().StringSliceS("gravity", "gravity", nil, "horizontal and vertical backdrop placement")
	displayCmd.Flags().BoolS("help", "help", false, "print program options")
	displayCmd.Flags().CountS("identify", "identify", "identify the format and characteristics of the image")
	displayCmd.Flags().StringSliceS("immutable", "immutable", nil, "prohibit image edits")
	displayCmd.Flags().StringSliceS("interlace", "interlace", nil, "type of image interlacing scheme")
	displayCmd.Flags().StringSliceS("interpolate", "interpolate", nil, "pixel color interpolation method")
	displayCmd.Flags().StringSliceS("label", "label", nil, "assign a label to an image")
	displayCmd.Flags().StringSliceS("limit", "limit", nil, "pixel cache resource limit")
	displayCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	displayCmd.Flags().StringSliceS("map", "map", nil, "transform image colors to match this set of colors")
	displayCmd.Flags().StringSliceS("mattecolor", "mattecolor", nil, "frame color")
	displayCmd.Flags().CountS("monitor", "monitor", "monitor progress")
	displayCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	displayCmd.Flags().CountS("negate", "negate", "replace each pixel with its complementary color")
	displayCmd.Flags().CountS("normalize", "normalize", "transform image to span the full range of colors")
	displayCmd.Flags().StringSliceS("page", "page", nil, "size and location of an image canvas (setting)")
	displayCmd.Flags().StringSliceS("profile", "profile", nil, "add, delete, or apply an image profile")
	displayCmd.Flags().StringSliceS("quantize", "quantize", nil, "reduce image colors in this colorspace")
	displayCmd.Flags().CountS("quiet", "quiet", "suppress all warning messages")
	displayCmd.Flags().StringSliceS("raise", "raise", nil, "lighten/darken image edges to create a 3-D effect")
	displayCmd.Flags().CountS("regard-warnings", "regard-warnings", "pay attention to warning messages.")
	displayCmd.Flags().StringSliceS("remote", "remote", nil, "execute a command in an remote display process")
	displayCmd.Flags().StringSliceS("resample", "resample", nil, "change the resolution of an image")
	displayCmd.Flags().StringSliceS("resize", "resize", nil, "resize the image")
	displayCmd.Flags().CountS("respect-parentheses", "respect-parentheses", "settings remain in effect until parenthesis boundary.")
	displayCmd.Flags().StringSliceS("roll", "roll", nil, "roll an image vertically or horizontally")
	displayCmd.Flags().StringSliceS("rotate", "rotate", nil, "apply Paeth rotation to the image")
	displayCmd.Flags().StringSliceS("sample", "sample", nil, "scale image with pixel sampling")
	displayCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factor")
	displayCmd.Flags().StringSliceS("scene", "scene", nil, "image scene number")
	displayCmd.Flags().StringSliceS("seed", "seed", nil, "seed a new sequence of pseudo-random numbers")
	displayCmd.Flags().StringSliceS("segment", "segment", nil, "segment an image")
	displayCmd.Flags().StringSliceS("set", "set", nil, "set an image attribute")
	displayCmd.Flags().StringSliceS("sharpen", "sharpen", nil, "sharpen the image")
	displayCmd.Flags().StringSliceS("size", "size", nil, "width and height of image")
	displayCmd.Flags().CountS("strip", "strip", "strip image of all profiles and comments")
	displayCmd.Flags().StringSliceS("thumbnail", "thumbnail", nil, "create a thumbnail of the image")
	displayCmd.Flags().StringSliceS("transparent-color", "transparent-color", nil, "transparent color")
	displayCmd.Flags().CountS("trim", "trim", "trim image edges")
	displayCmd.Flags().StringSliceS("update", "update", nil, "detect when image file is modified and redisplay")
	displayCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	displayCmd.Flags().BoolS("version", "version", false, "print version information")
	displayCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", nil, "access method for pixels outside the boundaries of the image")
	displayCmd.Flags().CountS("visual", "visual", "display image using this visual type")
	displayCmd.Flags().StringSliceS("window", "window", nil, "display image to background of this window")
	displayCmd.Flags().StringSliceS("window-group", "window-group", nil, "exit program when this window id is destroyed")
	displayCmd.Flags().StringSliceS("write", "write", nil, "write images to this file")
	rootCmd.AddCommand(displayCmd)

	carapace.Gen(displayCmd).FlagCompletion(carapace.ActionMap{
		"background":        action.ActionColors(),
		"bordercolor":       action.ActionColors(),
		"channel":           action.ActionCompleteOption("Channel").UniqueList(","),
		"colormap":          carapace.ActionValues("Shared", "Private"),
		"colorspace":        action.ActionCompleteOption("Colorspace"),
		"compress":          action.ActionCompleteOption("Compress"),
		"debug":             action.ActionCompleteOption("Debug").UniqueList(","),
		"decipher":          carapace.ActionFiles(),
		"dispose":           action.ActionCompleteOption("Dispose"),
		"dither":            action.ActionCompleteOption("Dither"),
		"endian":            action.ActionCompleteOption("Endian"),
		"filter":            action.ActionCompleteOption("Filter"),
		"interlace":         action.ActionCompleteOption("Interlace"),
		"interpolate":       action.ActionCompleteOption("Interpolate"),
		"map":               carapace.ActionFiles(),
		"mattecolor":        action.ActionColors(),
		"profile":           carapace.ActionFiles(),
		"quantize":          action.ActionCompleteOption("Colorspace"),
		"transparent-color": action.ActionColors(),
		"virtual-pixel":     action.ActionCompleteOption("VirtualPixel"),
		"write":             carapace.ActionFiles(),
	})

	carapace.Gen(displayCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
