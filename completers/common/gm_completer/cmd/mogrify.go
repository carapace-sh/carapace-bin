package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mogrifyCmd = &cobra.Command{
	Use:   "mogrify",
	Short: "transform an image or sequence of images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mogrifyCmd).Standalone()

	mogrifyCmd.Flags().StringSliceS("affine", "affine", nil, "affine transform matrix")
	mogrifyCmd.Flags().CountS("antialias", "antialias", "remove pixel-aliasing")
	mogrifyCmd.Flags().StringSliceS("asc-cdl", "asc-cdl", nil, "apply ASC CDL transform")
	mogrifyCmd.Flags().StringSliceS("authenticate", "authenticate", nil, "decrypt image with this password")
	mogrifyCmd.Flags().CountS("auto-orient", "auto-orient", "orient (rotate) image so it is upright")
	mogrifyCmd.Flags().StringSliceS("background", "background", nil, "background color")
	mogrifyCmd.Flags().StringSliceS("black-threshold", "black-threshold", nil, "pixels below the threshold become black")
	mogrifyCmd.Flags().StringSliceS("blue-primary", "blue-primary", nil, "chomaticity blue primary point")
	mogrifyCmd.Flags().StringSliceS("blur", "blur", nil, "blur the image")
	mogrifyCmd.Flags().StringSliceS("border", "border", nil, "surround image with a border of color")
	mogrifyCmd.Flags().StringSliceS("bordercolor", "bordercolor", nil, "border color")
	mogrifyCmd.Flags().StringSliceS("box", "box", nil, "set the color of the annotation bounding box")
	mogrifyCmd.Flags().StringSliceS("channel", "channel", nil, "extract a particular color channel from image")
	mogrifyCmd.Flags().StringSliceS("charcoal", "charcoal", nil, "simulate a charcoal drawing")
	mogrifyCmd.Flags().StringSliceS("chop", "chop", nil, "remove pixels from the image interior")
	mogrifyCmd.Flags().StringSliceS("colorize", "colorize", nil, "colorize the image with the fill color")
	mogrifyCmd.Flags().CountS("colors", "colors", "preferred number of colors in the image")
	mogrifyCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "alternate image colorspace")
	mogrifyCmd.Flags().StringSliceS("comment", "comment", nil, "annotate image with comment")
	mogrifyCmd.Flags().StringSliceS("compose", "compose", nil, "composite operator")
	mogrifyCmd.Flags().StringSliceS("compress", "compress", nil, "image compression type")
	mogrifyCmd.Flags().CountS("contrast", "contrast", "enhance or reduce the image contrast")
	mogrifyCmd.Flags().StringSliceS("convolve", "convolve", nil, "convolve image with the specified convolution kernel")
	mogrifyCmd.Flags().CountS("create-directories", "create-directories", "create output directories if required")
	mogrifyCmd.Flags().StringSliceS("crop", "crop", nil, "preferred size and location of the cropped image")
	mogrifyCmd.Flags().StringSliceS("cycle", "cycle", nil, "cycle the image colormap")
	mogrifyCmd.Flags().StringSliceS("debug", "debug", nil, "display copious debugging information")
	mogrifyCmd.Flags().StringSliceS("define", "define", nil, "Coder/decoder specific options")
	mogrifyCmd.Flags().StringSliceS("delay", "delay", nil, "display the next image after pausing")
	mogrifyCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	mogrifyCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	mogrifyCmd.Flags().CountS("despeckle", "despeckle", "reduce the speckles within an image")
	mogrifyCmd.Flags().StringSliceS("display", "display", nil, "get image or font from this X server")
	mogrifyCmd.Flags().StringSliceS("dispose", "dispose", nil, "Undefined, None, Background, Previous")
	mogrifyCmd.Flags().CountS("dither", "dither", "apply Floyd/Steinberg error diffusion to image")
	mogrifyCmd.Flags().StringSliceS("draw", "draw", nil, "annotate the image with a graphic primitive")
	mogrifyCmd.Flags().StringSliceS("edge", "edge", nil, "apply a filter to detect edges in the image")
	mogrifyCmd.Flags().StringSliceS("emboss", "emboss", nil, "emboss an image")
	mogrifyCmd.Flags().StringSliceS("encoding", "encoding", nil, "text encoding type")
	mogrifyCmd.Flags().StringSliceS("endian", "endian", nil, "multibyte word order (LSB, MSB, or Native)")
	mogrifyCmd.Flags().CountS("enhance", "enhance", "apply a digital filter to enhance a noisy image")
	mogrifyCmd.Flags().StringSliceS("equalize", "equalize", nil, "perform histogram equalization to an image")
	mogrifyCmd.Flags().CountS("extent", "extent", "composite image on background color canvas image")
	mogrifyCmd.Flags().StringSliceS("fill", "fill", nil, "color to use when filling a graphic primitive")
	mogrifyCmd.Flags().StringSliceS("filter", "filter", nil, "use this filter when resizing an image")
	mogrifyCmd.Flags().CountS("flip", "flip", "flip image in the vertical direction")
	mogrifyCmd.Flags().CountS("flop", "flop", "flop image in the horizontal direction")
	mogrifyCmd.Flags().StringSliceS("font", "font", nil, "render text with this font")
	mogrifyCmd.Flags().StringSliceS("format", "format", nil, "image format type")
	mogrifyCmd.Flags().StringSliceS("frame", "frame", nil, "surround image with an ornamental border")
	mogrifyCmd.Flags().StringSliceS("fuzz", "fuzz", nil, "colors within this distance are considered equal")
	mogrifyCmd.Flags().StringSliceS("gamma", "gamma", nil, "level of gamma correction")
	mogrifyCmd.Flags().StringSliceS("gaussian", "gaussian", nil, "gaussian blur an image")
	mogrifyCmd.Flags().StringSliceS("geometry", "geometry", nil, "preferred size or location of the image")
	mogrifyCmd.Flags().StringSliceS("gravity", "gravity", nil, "horizontal and vertical text/object placement")
	mogrifyCmd.Flags().StringSliceS("green-primary", "green-primary", nil, "chomaticity green primary point")
	mogrifyCmd.Flags().StringSliceS("hald-clut", "hald-clut", nil, "apply a Hald CLUT to the image")
	mogrifyCmd.Flags().BoolS("help", "help", false, "print program options")
	mogrifyCmd.Flags().StringSliceS("implode", "implode", nil, "implode image pixels about the center")
	mogrifyCmd.Flags().StringSliceS("interlace", "interlace", nil, "None, Line, Plane, or Partition")
	mogrifyCmd.Flags().StringSliceS("label", "label", nil, "assign a label to an image")
	mogrifyCmd.Flags().StringSliceS("lat", "lat", nil, "local adaptive thresholding")
	mogrifyCmd.Flags().StringSliceS("level", "level", nil, "adjust the level of image contrast")
	mogrifyCmd.Flags().StringSliceS("limit", "limit", nil, "Disk, File, Map, Memory, Pixels, Width, Height or Threads resource limit")
	mogrifyCmd.Flags().StringSliceS("linewidth", "linewidth", nil, "the line width for subsequent draw operations")
	mogrifyCmd.Flags().StringSliceS("list", "list", nil, "Color, Delegate, Format, Magic, Module, Resource, or Type")
	mogrifyCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	mogrifyCmd.Flags().StringSliceS("loop", "loop", nil, "add Netscape loop extension to your GIF animation")
	mogrifyCmd.Flags().CountS("magnify", "magnify", "interpolate image to double size")
	mogrifyCmd.Flags().StringSliceS("map", "map", nil, "transform image colors to match this set of colors")
	mogrifyCmd.Flags().StringSliceS("mask", "mask", nil, "set the image clip mask")
	mogrifyCmd.Flags().CountS("matte", "matte", "store matte channel if the image has one")
	mogrifyCmd.Flags().StringSliceS("mattecolor", "mattecolor", nil, "specify the color to be used with the -frame option")
	mogrifyCmd.Flags().StringSliceS("median", "median", nil, "apply a median filter to the image")
	mogrifyCmd.Flags().CountS("minify", "minify", "interpolate the image to half size")
	mogrifyCmd.Flags().StringSliceS("modulate", "modulate", nil, "vary the brightness, saturation, and hue")
	mogrifyCmd.Flags().CountS("monitor", "monitor", "show progress indication")
	mogrifyCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	mogrifyCmd.Flags().StringSliceS("motion-blur", "motion-blur", nil, "simulate motion blur")
	mogrifyCmd.Flags().CountS("negate", "negate", "replace every pixel with its complementary color")
	mogrifyCmd.Flags().StringSliceS("noise", "noise", nil, "add or reduce noise in an image")
	mogrifyCmd.Flags().CountS("noop", "noop", "do not apply options to image")
	mogrifyCmd.Flags().CountS("normalize", "normalize", "transform image to span the full range of colors")
	mogrifyCmd.Flags().StringSliceS("opaque", "opaque", nil, "change this color to the fill color")
	mogrifyCmd.Flags().StringSliceS("operator", "operator", nil, "apply a mathematical or bitwise operator to channel")
	mogrifyCmd.Flags().StringSliceS("ordered-dither", "ordered-dither", nil, "ordered dither the image")
	mogrifyCmd.Flags().StringSliceS("orient", "orient", nil, "set image orientation attribute")
	mogrifyCmd.Flags().StringSliceS("output-directory", "output-directory", nil, "write output files to directory")
	mogrifyCmd.Flags().StringSliceS("page", "page", nil, "size and location of an image canvas")
	mogrifyCmd.Flags().StringSliceS("paint", "paint", nil, "simulate an oil painting")
	mogrifyCmd.Flags().StringSliceS("pointsize", "pointsize", nil, "font point size")
	mogrifyCmd.Flags().CountS("preserve-timestamp", "preserve-timestamp", "preserve original timestamps of the file")
	mogrifyCmd.Flags().StringSliceS("profile", "profile", nil, "add ICM or IPTC information profile to image")
	mogrifyCmd.Flags().StringSliceS("quality", "quality", nil, "JPEG/MIFF/PNG compression level")
	mogrifyCmd.Flags().StringSliceS("raise", "raise", nil, "lighten/darken image edges to create a 3-D effect")
	mogrifyCmd.Flags().StringSliceS("random-threshold", "random-threshold", nil, "random threshold the image")
	mogrifyCmd.Flags().StringSliceS("recolor", "recolor", nil, "apply a color translation matrix to image channels")
	mogrifyCmd.Flags().StringSliceS("red-primary", "red-primary", nil, "chomaticity red primary point")
	mogrifyCmd.Flags().StringSliceS("region", "region", nil, "apply options to a portion of the image")
	mogrifyCmd.Flags().CountS("render", "render", "render vector graphics")
	mogrifyCmd.Flags().StringSliceS("repage", "repage", nil, "adjust current page offsets by geometry")
	mogrifyCmd.Flags().StringSliceS("resample", "resample", nil, "resample to horizontal and vertical resolution")
	mogrifyCmd.Flags().StringSliceS("resize", "resize", nil, "preferred size or location of the image")
	mogrifyCmd.Flags().StringSliceS("roll", "roll", nil, "roll an image vertically or horizontally")
	mogrifyCmd.Flags().StringSliceS("rotate", "rotate", nil, "apply Paeth rotation to the image")
	mogrifyCmd.Flags().StringSliceS("sample", "sample", nil, "scale image with pixel sampling")
	mogrifyCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factors")
	mogrifyCmd.Flags().StringSliceS("scale", "scale", nil, "scale the image")
	mogrifyCmd.Flags().StringSliceS("scene", "scene", nil, "image scene number")
	mogrifyCmd.Flags().StringSliceS("seed", "seed", nil, "pseudo-random number generator seed value")
	mogrifyCmd.Flags().StringSliceS("segment", "segment", nil, "segment an image")
	mogrifyCmd.Flags().StringSliceS("set", "set", nil, "set image attribute")
	mogrifyCmd.Flags().StringSliceS("shade", "shade", nil, "shade the image using a distant light source")
	mogrifyCmd.Flags().StringSliceS("sharpen", "sharpen", nil, "sharpen the image")
	mogrifyCmd.Flags().StringSliceS("shave", "shave", nil, "shave pixels from the image edges")
	mogrifyCmd.Flags().StringSliceS("shear", "shear", nil, "slide one edge of the image along the X or Y axis")
	mogrifyCmd.Flags().StringSliceS("size", "size", nil, "width and height of image")
	mogrifyCmd.Flags().StringSliceS("solarize", "solarize", nil, "negate all pixels above the threshold level")
	mogrifyCmd.Flags().StringSliceS("spread", "spread", nil, "displace image pixels by a random amount")
	mogrifyCmd.Flags().CountS("strip", "strip", "strip all profiles and text attributes from image")
	mogrifyCmd.Flags().StringSliceS("stroke", "stroke", nil, "graphic primitive stroke color")
	mogrifyCmd.Flags().StringSliceS("strokewidth", "strokewidth", nil, "graphic primitive stroke width")
	mogrifyCmd.Flags().StringSliceS("swirl", "swirl", nil, "swirl image pixels about the center")
	mogrifyCmd.Flags().StringSliceS("texture", "texture", nil, "name of texture to tile onto the image background")
	mogrifyCmd.Flags().StringSliceS("threshold", "threshold", nil, "threshold the image")
	mogrifyCmd.Flags().StringSliceS("thumbnail", "thumbnail", nil, "resize the image (optimized for thumbnails)")
	mogrifyCmd.Flags().StringSliceS("tile", "tile", nil, "tile image when filling a graphic primitive")
	mogrifyCmd.Flags().CountS("transform", "transform", "affine transform image")
	mogrifyCmd.Flags().StringSliceS("transparent", "transparent", nil, "make this color transparent within the image")
	mogrifyCmd.Flags().StringSliceS("treedepth", "treedepth", nil, "color tree depth")
	mogrifyCmd.Flags().CountS("trim", "trim", "trim image edges")
	mogrifyCmd.Flags().StringSliceS("type", "type", nil, "image type")
	mogrifyCmd.Flags().StringSliceS("undercolor", "undercolor", nil, "annotation bounding box color")
	mogrifyCmd.Flags().StringSliceS("units", "units", nil, "PixelsPerInch, PixelsPerCentimeter, or Undefined")
	mogrifyCmd.Flags().StringSliceS("unsharp", "unsharp", nil, "sharpen the image")
	mogrifyCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	mogrifyCmd.Flags().BoolS("version", "version", false, "print version information")
	mogrifyCmd.Flags().CountS("view", "view", "FlashPix viewing transforms")
	mogrifyCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", nil, "Constant, Edge, Mirror, or Tile")
	mogrifyCmd.Flags().StringSliceS("wave", "wave", nil, "alter an image along a sine wave")
	mogrifyCmd.Flags().StringSliceS("white-point", "white-point", nil, "chomaticity white point")
	mogrifyCmd.Flags().StringSliceS("white-threshold", "white-threshold", nil, "pixels above the threshold become white")
	rootCmd.AddCommand(mogrifyCmd)

	carapace.Gen(mogrifyCmd).FlagCompletion(carapace.ActionMap{
		"fill": action.ActionColor(),
	})

	carapace.Gen(mogrifyCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
