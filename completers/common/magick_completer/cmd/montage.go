package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/magick_completer/cmd/action"
	"github.com/spf13/cobra"
)

var montageCmd = &cobra.Command{
	Use:   "montage",
	Short: "create a composite image by combining several separate images. The images are tiled on the composite image optionally adorned with a border, frame, image name, and more.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(montageCmd).Standalone()

	montageCmd.Flags().StringSliceS("adaptive-sharpen", "adaptive-sharpen", nil, "adaptively sharpen pixels; increase effect near edges")
	montageCmd.Flags().CountS("adjoin", "adjoin", "join images into a single multi-image file")
	montageCmd.Flags().StringSliceS("affine", "affine", nil, "affine transform matrix")
	montageCmd.Flags().CountS("alpha", "alpha", "on, activate, off, deactivate, set, opaque, copy\", transparent, extract, background, or shape the alpha channel")
	montageCmd.Flags().StringSliceS("annotate", "annotate", nil, "annotate the image with text")
	montageCmd.Flags().StringSliceS("authenticate", "authenticate", nil, "decrypt image with this password")
	montageCmd.Flags().CountS("auto-orient", "auto-orient", "automagically orient image")
	montageCmd.Flags().StringSliceS("background", "background", nil, "background color")
	montageCmd.Flags().StringSliceS("blue-primary", "blue-primary", nil, "chromaticity blue primary point")
	montageCmd.Flags().StringSliceS("blur", "blur", nil, "reduce image noise and reduce detail levels")
	montageCmd.Flags().StringSliceS("border", "border", nil, "surround image with a border of color")
	montageCmd.Flags().StringSliceS("bordercolor", "bordercolor", nil, "border color")
	montageCmd.Flags().StringSliceS("caption", "caption", nil, "assign a caption to an image")
	montageCmd.Flags().StringSliceS("channel", "channel", nil, "apply option to select image channels")
	montageCmd.Flags().StringSliceS("clone", "clone", nil, "clone an image")
	montageCmd.Flags().CountS("coalesce", "coalesce", "merge a sequence of images")
	montageCmd.Flags().StringSliceS("colors", "colors", nil, "preferred number of colors in the image")
	montageCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "set image colorspace")
	montageCmd.Flags().StringSliceS("comment", "comment", nil, "annotate image with comment")
	montageCmd.Flags().StringSliceS("compose", "compose", nil, "set image composite operator")
	montageCmd.Flags().CountS("composite", "composite", "composite image")
	montageCmd.Flags().StringSliceS("compress", "compress", nil, "image compression type")
	montageCmd.Flags().StringSliceS("crop", "crop", nil, "preferred size and location of the cropped image")
	montageCmd.Flags().StringSliceS("debug", "debug", nil, "display copious debugging information")
	montageCmd.Flags().StringSliceS("define", "define", nil, "define one or more image format options")
	montageCmd.Flags().StringSliceS("delay", "delay", nil, "display the next image after pausing")
	montageCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	montageCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	montageCmd.Flags().StringSliceS("display", "display", nil, "get image or font from this X server")
	montageCmd.Flags().StringSliceS("dispose", "dispose", nil, "layer disposal method")
	montageCmd.Flags().StringSliceS("distort", "distort", nil, "distort image")
	montageCmd.Flags().StringSliceS("dither", "dither", nil, "apply error diffusion to image")
	montageCmd.Flags().StringSliceS("draw", "draw", nil, "annotate the image with a graphic primitive")
	montageCmd.Flags().StringSliceS("duplicate", "duplicate", nil, "duplicate an image one or more times")
	montageCmd.Flags().StringSliceS("endian", "endian", nil, "endianness (MSB or LSB) of the image")
	montageCmd.Flags().StringSliceS("extent", "extent", nil, "set the image size")
	montageCmd.Flags().StringSliceS("extract", "extract", nil, "extract area from image")
	montageCmd.Flags().StringSliceS("family", "family", nil, "render text with this font family")
	montageCmd.Flags().StringSliceS("fill", "fill", nil, "color to use when filling a graphic primitive")
	montageCmd.Flags().StringSliceS("filter", "filter", nil, "use this filter when resizing an image")
	montageCmd.Flags().CountS("flatten", "flatten", "flatten a sequence of images")
	montageCmd.Flags().CountS("flip", "flip", "flip image in the vertical direction")
	montageCmd.Flags().CountS("flop", "flop", "flop image in the horizontal direction")
	montageCmd.Flags().StringSliceS("font", "font", nil, "render text with this font")
	montageCmd.Flags().StringSliceS("frame", "frame", nil, "surround image with an ornamental border")
	montageCmd.Flags().StringSliceS("gamma", "gamma", nil, "level of gamma correction")
	montageCmd.Flags().StringSliceS("geometry", "geometry", nil, "preferred size or location of the image")
	montageCmd.Flags().StringSliceS("gravity", "gravity", nil, "horizontal and vertical text placement")
	montageCmd.Flags().StringSliceS("green-primary", "green-primary", nil, "chromaticity green primary point")
	montageCmd.Flags().BoolS("help", "help", false, "print program options")
	montageCmd.Flags().CountS("identify", "identify", "identify the format and characteristics of the image")
	montageCmd.Flags().StringSliceS("interlace", "interlace", nil, "type of image interlacing scheme")
	montageCmd.Flags().StringSliceS("interpolate", "interpolate", nil, "pixel color interpolation method")
	montageCmd.Flags().StringSliceS("kerning", "kerning", nil, "the space between two characters")
	montageCmd.Flags().StringSliceS("label", "label", nil, "assign a label to an image")
	montageCmd.Flags().StringSliceS("limit", "limit", nil, "pixel cache resource limit")
	montageCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	montageCmd.Flags().StringSliceS("mattecolor", "mattecolor", nil, "frame color")
	montageCmd.Flags().StringSliceS("mode", "mode", nil, "framing style")
	montageCmd.Flags().CountS("monitor", "monitor", "monitor progress")
	montageCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	montageCmd.Flags().StringSliceS("origin", "origin", nil, "image origin")
	montageCmd.Flags().StringSliceS("page", "page", nil, "size and location of an image canvas (setting)")
	montageCmd.Flags().StringSliceS("pointsize", "pointsize", nil, "font point size")
	montageCmd.Flags().StringSliceS("polaroid", "polaroid", nil, "simulate a Polaroid picture")
	montageCmd.Flags().StringSliceS("profile", "profile", nil, "add, delete, or apply an image profile")
	montageCmd.Flags().StringSliceS("quality", "quality", nil, "JPEG/MIFF/PNG compression level")
	montageCmd.Flags().StringSliceS("quantize", "quantize", nil, "reduce image colors in this colorspace")
	montageCmd.Flags().CountS("quiet", "quiet", "suppress all warning messages")
	montageCmd.Flags().StringSliceS("red-primary", "red-primary", nil, "chromaticity red primary point")
	montageCmd.Flags().CountS("regard-warnings", "regard-warnings", "pay attention to warning messages.")
	montageCmd.Flags().StringSliceS("repage", "repage", nil, "size and location of an image canvas")
	montageCmd.Flags().StringSliceS("resize", "resize", nil, "resize the image")
	montageCmd.Flags().CountS("respect-parentheses", "respect-parentheses", "settings remain in effect until parenthesis boundary.")
	montageCmd.Flags().StringSliceS("rotate", "rotate", nil, "apply Paeth rotation to the image")
	montageCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factor")
	montageCmd.Flags().StringSliceS("scale", "scale", nil, "scale the image")
	montageCmd.Flags().StringSliceS("scenes", "scenes", nil, "image scene range")
	montageCmd.Flags().StringSliceS("seed", "seed", nil, "seed a new sequence of pseudo-random numbers")
	montageCmd.Flags().StringSliceS("shadow", "shadow", nil, "simulate an image shadow")
	montageCmd.Flags().StringSliceS("size", "size", nil, "width and height of image")
	montageCmd.Flags().CountS("strip", "strip", "strip image of all profiles and comments")
	montageCmd.Flags().StringSliceS("stroke", "stroke", nil, "graphic primitive stroke color")
	montageCmd.Flags().CountS("synchronize", "synchronize", "synchronize image to storage device")
	montageCmd.Flags().CountS("taint", "taint", "mark the image as modified")
	montageCmd.Flags().StringSliceS("texture", "texture", nil, "name of texture to tile onto the image background")
	montageCmd.Flags().StringSliceS("tile", "tile", nil, "number of tiles per row and column (e.g. -tile 8x)")
	montageCmd.Flags().StringSliceS("tile-offset", "tile-offset", nil, "set the image tile offset")
	montageCmd.Flags().CountS("title", "title", "decorate the montage image with a title")
	montageCmd.Flags().CountS("transform", "transform", "affine transform image")
	montageCmd.Flags().StringSliceS("transparent", "transparent", nil, "make this color transparent within the image")
	montageCmd.Flags().StringSliceS("transparent-color", "transparent-color", nil, "transparent color")
	montageCmd.Flags().CountS("transpose", "transpose", "flip image in the vertical direction and rotate 90 degrees")
	montageCmd.Flags().StringSliceS("treedepth", "treedepth", nil, "color tree depth")
	montageCmd.Flags().CountS("trim", "trim", "trim image edges")
	montageCmd.Flags().StringSliceS("type", "type", nil, "image type")
	montageCmd.Flags().StringSliceS("units", "units", nil, "the units of image resolution")
	montageCmd.Flags().StringSliceS("unsharp", "unsharp", nil, "sharpen the image")
	montageCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	montageCmd.Flags().BoolS("version", "version", false, "print version information")
	montageCmd.Flags().CountS("view", "view", "FlashPix viewing transforms")
	montageCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", nil, "access method for pixels outside the boundaries of the image")
	montageCmd.Flags().StringSliceS("white-point", "white-point", nil, "chromaticity white point")
	rootCmd.AddCommand(montageCmd)

	carapace.Gen(montageCmd).FlagCompletion(carapace.ActionMap{
		"background":        action.ActionColors(),
		"bordercolor":       action.ActionColors(),
		"channel":           action.ActionCompleteOption("Channel").UniqueList(","),
		"colorspace":        action.ActionCompleteOption("Colorspace"),
		"compose":           action.ActionCompleteOption("Compose"),
		"compress":          action.ActionCompleteOption("Compress"),
		"debug":             action.ActionCompleteOption("Debug").UniqueList(","),
		"dispose":           action.ActionCompleteOption("Dispose"),
		"distort":           action.ActionCompleteOption("Distort"),
		"dither":            action.ActionCompleteOption("Dither"),
		"endian":            action.ActionCompleteOption("Endian"),
		"family":            action.ActionFontFamilies(),
		"fill":              action.ActionColors(),
		"filter":            action.ActionCompleteOption("Filter"),
		"font":              action.ActionFonts(),
		"gravity":           action.ActionCompleteOption("Gravity"),
		"interlace":         action.ActionCompleteOption("Interlace"),
		"interpolate":       action.ActionCompleteOption("Interpolate"),
		"mattecolor":        action.ActionColors(),
		"mode":              action.ActionCompleteOption("Mode"),
		"profile":           carapace.ActionFiles(),
		"quantize":          action.ActionCompleteOption("Colorspace"),
		"stroke":            action.ActionColors(),
		"texture":           carapace.ActionFiles(),
		"transparent":       action.ActionColors(),
		"transparent-color": action.ActionColors(),
		"type":              action.ActionCompleteOption("Type"),
		"units":             action.ActionCompleteOption("Units"),
		"virtual-pixel":     action.ActionCompleteOption("VirtualPixel"),
	})

	carapace.Gen(montageCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
