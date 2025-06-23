package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/magick_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compositeCmd = &cobra.Command{
	Use:   "composite",
	Short: "overlap one image over another.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compositeCmd).Standalone()

	compositeCmd.Flags().StringSliceS("affine", "affine", nil, "affine transform matrix")
	compositeCmd.Flags().CountS("alpha", "alpha", "on, activate, off, deactivate, set, opaque, copy\", transparent, extract, background, or shape the alpha channel")
	compositeCmd.Flags().StringSliceS("authenticate", "authenticate", nil, "decrypt image with this password")
	compositeCmd.Flags().StringSliceS("blend", "blend", nil, "blend images")
	compositeCmd.Flags().StringSliceS("blue-primary", "blue-primary", nil, "chromaticity blue primary point")
	compositeCmd.Flags().StringSliceS("border", "border", nil, "surround image with a border of color")
	compositeCmd.Flags().StringSliceS("bordercolor", "bordercolor", nil, "border color")
	compositeCmd.Flags().StringSliceS("channel", "channel", nil, "apply option to select image channels")
	compositeCmd.Flags().StringSliceS("colors", "colors", nil, "preferred number of colors in the image")
	compositeCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "set image colorspace")
	compositeCmd.Flags().StringSliceS("comment", "comment", nil, "annotate image with comment")
	compositeCmd.Flags().StringSliceS("compose", "compose", nil, "set image composite operator")
	compositeCmd.Flags().StringSliceS("compress", "compress", nil, "image compression type")
	compositeCmd.Flags().StringSliceS("debug", "debug", nil, "display copious debugging information")
	compositeCmd.Flags().StringSliceS("decipher", "decipher", nil, "convert cipher pixels to plain")
	compositeCmd.Flags().StringSliceS("define", "define", nil, "define one or more image format options")
	compositeCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	compositeCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	compositeCmd.Flags().StringSliceS("displace", "displace", nil, "shift image pixels defined by a displacement map")
	compositeCmd.Flags().StringSliceS("dissolve", "dissolve", nil, "dissolve the two images a given percent")
	compositeCmd.Flags().StringSliceS("dither", "dither", nil, "apply error diffusion to image")
	compositeCmd.Flags().StringSliceS("encipher", "encipher", nil, "convert plain pixels to cipher pixels")
	compositeCmd.Flags().StringSliceS("encoding", "encoding", nil, "text encoding type")
	compositeCmd.Flags().StringSliceS("endian", "endian", nil, "endianness (MSB or LSB) of the image")
	compositeCmd.Flags().StringSliceS("extract", "extract", nil, "extract area from image")
	compositeCmd.Flags().StringSliceS("filter", "filter", nil, "use this filter when resizing an image")
	compositeCmd.Flags().StringSliceS("font", "font", nil, "render text with this font")
	compositeCmd.Flags().StringSliceS("geometry", "geometry", nil, "preferred size or location of the image")
	compositeCmd.Flags().StringSliceS("gravity", "gravity", nil, "horizontal and vertical text placement")
	compositeCmd.Flags().StringSliceS("green-primary", "green-primary", nil, "chromaticity green primary point")
	compositeCmd.Flags().BoolS("help", "help", false, "print program options")
	compositeCmd.Flags().CountS("identify", "identify", "identify the format and characteristics of the image")
	compositeCmd.Flags().StringSliceS("interlace", "interlace", nil, "type of image interlacing scheme")
	compositeCmd.Flags().StringSliceS("interpolate", "interpolate", nil, "pixel color interpolation method")
	compositeCmd.Flags().StringSliceS("label", "label", nil, "assign a label to an image")
	compositeCmd.Flags().StringSliceS("level", "level", nil, "adjust the level of image contrast")
	compositeCmd.Flags().StringSliceS("limit", "limit", nil, "pixel cache resource limit")
	compositeCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	compositeCmd.Flags().CountS("monitor", "monitor", "monitor progress")
	compositeCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	compositeCmd.Flags().CountS("negate", "negate", "replace each pixel with its complementary color")
	compositeCmd.Flags().StringSliceS("page", "page", nil, "size and location of an image canvas (setting)")
	compositeCmd.Flags().StringSliceS("pointsize", "pointsize", nil, "font point size")
	compositeCmd.Flags().StringSliceS("profile", "profile", nil, "add, delete, or apply an image profile")
	compositeCmd.Flags().StringSliceS("quality", "quality", nil, "JPEG/MIFF/PNG compression level")
	compositeCmd.Flags().StringSliceS("quantize", "quantize", nil, "reduce image colors in this colorspace")
	compositeCmd.Flags().CountS("quiet", "quiet", "suppress all warning messages")
	compositeCmd.Flags().StringSliceS("red-primary", "red-primary", nil, "chromaticity red primary point")
	compositeCmd.Flags().CountS("regard-warnings", "regard-warnings", "pay attention to warning messages.")
	compositeCmd.Flags().CountS("respect-parentheses", "respect-parentheses", "settings remain in effect until parenthesis boundary.")
	compositeCmd.Flags().StringSliceS("rotate", "rotate", nil, "apply Paeth rotation to the image")
	compositeCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factor")
	compositeCmd.Flags().StringSliceS("scene", "scene", nil, "image scene number")
	compositeCmd.Flags().StringSliceS("seed", "seed", nil, "seed a new sequence of pseudo-random numbers")
	compositeCmd.Flags().StringSliceS("set", "set", nil, "set an image attribute")
	compositeCmd.Flags().StringSliceS("sharpen", "sharpen", nil, "sharpen the image")
	compositeCmd.Flags().StringSliceS("shave", "shave", nil, "shave pixels from the image edges")
	compositeCmd.Flags().StringSliceS("size", "size", nil, "width and height of image")
	compositeCmd.Flags().StringSliceS("stegano", "stegano", nil, "hide watermark within an image")
	compositeCmd.Flags().StringSliceS("stereo", "stereo", nil, "combine two image to create a stereo anaglyph")
	compositeCmd.Flags().CountS("strip", "strip", "strip image of all profiles and comments")
	compositeCmd.Flags().StringSliceS("swap", "swap", nil, "swap two images in the image sequence")
	compositeCmd.Flags().CountS("synchronize", "synchronize", "synchronize image to storage device")
	compositeCmd.Flags().CountS("taint", "taint", "mark the image as modified")
	compositeCmd.Flags().StringSliceS("thumbnail", "thumbnail", nil, "create a thumbnail of the image")
	compositeCmd.Flags().CountS("tile", "tile", "repeat composite operation across and down image")
	compositeCmd.Flags().CountS("transform", "transform", "affine transform image")
	compositeCmd.Flags().StringSliceS("transparent-color", "transparent-color", nil, "transparent color")
	compositeCmd.Flags().StringSliceS("treedepth", "treedepth", nil, "color tree depth")
	compositeCmd.Flags().StringSliceS("type", "type", nil, "image type")
	compositeCmd.Flags().StringSliceS("units", "units", nil, "the units of image resolution")
	compositeCmd.Flags().StringSliceS("unsharp", "unsharp", nil, "sharpen the image")
	compositeCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	compositeCmd.Flags().BoolS("version", "version", false, "print version information")
	compositeCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", nil, "access method for pixels outside the boundaries of the image")
	compositeCmd.Flags().StringSliceS("watermark", "watermark", nil, "percent brightness and saturation of a watermark")
	compositeCmd.Flags().StringSliceS("white-point", "white-point", nil, "chromaticity white point")
	compositeCmd.Flags().StringSliceS("white-threshold", "white-threshold", nil, "force all pixels above the threshold into white")
	compositeCmd.Flags().StringSliceS("write", "write", nil, "write images to this file")
	rootCmd.AddCommand(compositeCmd)

	carapace.Gen(compositeCmd).FlagCompletion(carapace.ActionMap{
		"bordercolor":       action.ActionColors(),
		"channel":           action.ActionCompleteOption("Channel").UniqueList(","),
		"colorspace":        action.ActionCompleteOption("Colorspace"),
		"compose":           action.ActionCompleteOption("Compose"),
		"compress":          action.ActionCompleteOption("Compress"),
		"debug":             action.ActionCompleteOption("Debug").UniqueList(","),
		"decipher":          carapace.ActionFiles(),
		"dither":            action.ActionCompleteOption("Dither"),
		"encipher":          carapace.ActionFiles(),
		"encoding":          carapace.ActionValues("AdobeCustom", "AdobeExpert", "AdobeStandard", "AppleRoman", "BIG5", "GB2312", "Latin 2", "None", "SJIScode", "Symbol", "Unicode", "Wansung"),
		"endian":            action.ActionCompleteOption("Endian"),
		"filter":            action.ActionCompleteOption("Filter"),
		"font":              action.ActionFonts(),
		"gravity":           action.ActionCompleteOption("Gravity"),
		"interlace":         action.ActionCompleteOption("Interlace"),
		"interpolate":       action.ActionCompleteOption("Interpolate"),
		"profile":           carapace.ActionFiles(),
		"quantize":          action.ActionCompleteOption("Colorspace"),
		"transparent-color": action.ActionColors(),
		"type":              action.ActionCompleteOption("Type"),
		"units":             action.ActionCompleteOption("Units"),
		"virtual-pixel":     action.ActionCompleteOption("VirtualPixel"),
		"write":             carapace.ActionFiles(),
	})

	carapace.Gen(compositeCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
