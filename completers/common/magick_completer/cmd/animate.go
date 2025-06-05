package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/magick_completer/cmd/action"
	"github.com/spf13/cobra"
)

var animateCmd = &cobra.Command{
	Use:   "animate",
	Short: "animate an image sequence on any X server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(animateCmd).Standalone()

	animateCmd.Flags().CountS("alpha", "alpha", "on, activate, off, deactivate, set, opaque, copy\", transparent, extract, background, or shape the alpha channel")
	animateCmd.Flags().CountS("antialias", "antialias", "remove pixel-aliasing")
	animateCmd.Flags().StringSliceS("authenticate", "authenticate", nil, "decrypt image with this password")
	animateCmd.Flags().CountS("backdrop", "backdrop", "background color")
	animateCmd.Flags().StringSliceS("background", "background", nil, "background color")
	animateCmd.Flags().StringSliceS("border", "border", nil, "surround image with a border of color")
	animateCmd.Flags().StringSliceS("bordercolor", "bordercolor", nil, "border color")
	animateCmd.Flags().StringSliceS("channel", "channel", nil, "apply option to select image channels")
	animateCmd.Flags().CountS("clip", "clip", "clip along the first path from the 8BIM profile")
	animateCmd.Flags().StringSliceS("clip-path", "clip-path", nil, "clip along a named path from the 8BIM profile")
	animateCmd.Flags().CountS("coalesce", "coalesce", "merge a sequence of images")
	animateCmd.Flags().StringSliceS("colormap", "colormap", nil, "Shared or Private")
	animateCmd.Flags().StringSliceS("colors", "colors", nil, "preferred number of colors in the image")
	animateCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "set image colorspace")
	animateCmd.Flags().StringSliceS("comment", "comment", nil, "annotate image with comment")
	animateCmd.Flags().StringSliceS("compress", "compress", nil, "image compression type")
	animateCmd.Flags().CountS("contrast", "contrast", "enhance or reduce the image contrast")
	animateCmd.Flags().StringSliceS("crop", "crop", nil, "preferred size and location of the cropped image")
	animateCmd.Flags().StringSliceS("debug", "debug", nil, "animate copious debugging information")
	animateCmd.Flags().StringSliceS("decipher", "decipher", nil, "convert cipher pixels to plain")
	animateCmd.Flags().StringSliceS("define", "define", nil, "define one or more image format options")
	animateCmd.Flags().StringSliceS("delay", "delay", nil, "animate the next image after pausing")
	animateCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	animateCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	animateCmd.Flags().CountS("despeckle", "despeckle", "reduce the speckles within an image")
	animateCmd.Flags().StringSliceS("display", "display", nil, "get image or font from this X server")
	animateCmd.Flags().StringSliceS("dispose", "dispose", nil, "layer disposal method")
	animateCmd.Flags().StringSliceS("dither", "dither", nil, "apply error diffusion to image")
	animateCmd.Flags().StringSliceS("edge", "edge", nil, "apply a filter to detect edges in the image")
	animateCmd.Flags().StringSliceS("endian", "endian", nil, "endianness (MSB or LSB) of the image")
	animateCmd.Flags().CountS("enhance", "enhance", "apply a digital filter to enhance a noisy image")
	animateCmd.Flags().StringSliceS("extract", "extract", nil, "extract area from image")
	animateCmd.Flags().StringSliceS("filter", "filter", nil, "use this filter when resizing an image")
	animateCmd.Flags().CountS("flatten", "flatten", "flatten a sequence of images")
	animateCmd.Flags().CountS("flip", "flip", "flip image in the vertical direction")
	animateCmd.Flags().CountS("flop", "flop", "flop image in the horizontal direction")
	animateCmd.Flags().StringSliceS("frame", "frame", nil, "surround image with an ornamental border")
	animateCmd.Flags().StringSliceS("gamma", "gamma", nil, "level of gamma correction")
	animateCmd.Flags().StringSliceS("geometry", "geometry", nil, "preferred size or location of the image")
	animateCmd.Flags().StringSliceS("gravity", "gravity", nil, "horizontal and vertical backdrop placement")
	animateCmd.Flags().BoolS("help", "help", false, "print program options")
	animateCmd.Flags().CountS("identify", "identify", "identify the format and characteristics of the image")
	animateCmd.Flags().StringSliceS("immutable", "immutable", nil, "prohibit image edits")
	animateCmd.Flags().StringSliceS("interlace", "interlace", nil, "type of image interlacing scheme")
	animateCmd.Flags().StringSliceS("interpolate", "interpolate", nil, "pixel color interpolation method")
	animateCmd.Flags().StringSliceS("label", "label", nil, "assign a label to an image")
	animateCmd.Flags().StringSliceS("limit", "limit", nil, "pixel cache resource limit")
	animateCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	animateCmd.Flags().StringSliceS("map", "map", nil, "transform image colors to match this set of colors")
	animateCmd.Flags().StringSliceS("mattecolor", "mattecolor", nil, "frame color")
	animateCmd.Flags().CountS("monitor", "monitor", "monitor progress")
	animateCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	animateCmd.Flags().CountS("negate", "negate", "replace each pixel with its complementary color")
	animateCmd.Flags().StringSliceS("page", "page", nil, "size and location of an image canvas (setting)")
	animateCmd.Flags().StringSliceS("profile", "profile", nil, "add, delete, or apply an image profile")
	animateCmd.Flags().StringSliceS("quantize", "quantize", nil, "reduce image colors in this colorspace")
	animateCmd.Flags().CountS("quiet", "quiet", "suppress all warning messages")
	animateCmd.Flags().StringSliceS("raise", "raise", nil, "lighten/darken image edges to create a 3-D effect")
	animateCmd.Flags().CountS("regard-warnings", "regard-warnings", "pay attention to warning messages.")
	animateCmd.Flags().StringSliceS("remote", "remote", nil, "execute a command in an remote animate process")
	animateCmd.Flags().StringSliceS("resample", "resample", nil, "change the resolution of an image")
	animateCmd.Flags().StringSliceS("resize", "resize", nil, "resize the image")
	animateCmd.Flags().CountS("respect-parentheses", "respect-parentheses", "settings remain in effect until parenthesis boundary.")
	animateCmd.Flags().StringSliceS("roll", "roll", nil, "roll an image vertically or horizontally")
	animateCmd.Flags().StringSliceS("rotate", "rotate", nil, "apply Paeth rotation to the image")
	animateCmd.Flags().StringSliceS("sample", "sample", nil, "scale image with pixel sampling")
	animateCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factor")
	animateCmd.Flags().StringSliceS("scene", "scene", nil, "image scene number")
	animateCmd.Flags().StringSliceS("seed", "seed", nil, "seed a new sequence of pseudo-random numbers")
	animateCmd.Flags().StringSliceS("segment", "segment", nil, "segment an image")
	animateCmd.Flags().StringSliceS("set", "set", nil, "set an image attribute")
	animateCmd.Flags().StringSliceS("sharpen", "sharpen", nil, "sharpen the image")
	animateCmd.Flags().StringSliceS("size", "size", nil, "width and height of image")
	animateCmd.Flags().CountS("strip", "strip", "strip image of all profiles and comments")
	animateCmd.Flags().StringSliceS("thumbnail", "thumbnail", nil, "create a thumbnail of the image")
	animateCmd.Flags().StringSliceS("transparent-color", "transparent-color", nil, "transparent color")
	animateCmd.Flags().CountS("trim", "trim", "trim image edges")
	animateCmd.Flags().StringSliceS("update", "update", nil, "detect when image file is modified and reanimate")
	animateCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	animateCmd.Flags().BoolS("version", "version", false, "print version information")
	animateCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", nil, "access method for pixels outside the boundaries of the image")
	animateCmd.Flags().CountS("visual", "visual", "animate image using this visual type")
	animateCmd.Flags().StringSliceS("window", "window", nil, "animate images to background of this window")
	animateCmd.Flags().StringSliceS("window-group", "window-group", nil, "exit program when this window id is destroyed")
	rootCmd.AddCommand(animateCmd)

	carapace.Gen(animateCmd).FlagCompletion(carapace.ActionMap{
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
		"gravity":           action.ActionCompleteOption("Gravity"),
		"interlace":         action.ActionCompleteOption("Interlace"),
		"interpolate":       action.ActionCompleteOption("Interpolate"),
		"map":               carapace.ActionFiles(),
		"mattecolor":        action.ActionColors(),
		"profile":           carapace.ActionFiles(),
		"quantize":          action.ActionCompleteOption("Colorspace"),
		"transparent-color": action.ActionColors(),
		"virtual-pixel":     action.ActionCompleteOption("VirtualPixel"),
	})

	carapace.Gen(animateCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
