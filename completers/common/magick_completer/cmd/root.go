package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/magick_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "magick",
	Short: "convert between image formats as well as resize an image, blur, crop, despeckle, dither, draw on, flip, join, re-sample, and much more",
	Long:  "https://imagemagick.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringSliceS("adaptive-blur", "adaptive-blur", nil, "adaptively blur pixels; decrease effect near edges")
	rootCmd.Flags().StringSliceS("adaptive-resize", "adaptive-resize", nil, "adaptively resize image with data dependent triangulation.")
	rootCmd.Flags().StringSliceS("adaptive-sharpen", "adaptive-sharpen", nil, "adaptively sharpen pixels; increase effect near edges")
	rootCmd.Flags().CountS("adjoin", "adjoin", "join images into a single multi-image file")
	rootCmd.Flags().StringSliceS("affine", "affine", nil, "affine transform matrix")
	rootCmd.Flags().CountS("alpha", "alpha", "on, activate, off, deactivate, set, opaque, copy\", transparent, extract, background, or shape the alpha channel")
	rootCmd.Flags().StringSliceS("annotate", "annotate", nil, "annotate the image with text")
	rootCmd.Flags().CountS("antialias", "antialias", "remove pixel-aliasing")
	rootCmd.Flags().CountS("append", "append", "append an image sequence")
	rootCmd.Flags().StringSliceS("authenticate", "authenticate", nil, "decipher image with this password")
	rootCmd.Flags().CountS("auto-gamma", "auto-gamma", "automagically adjust gamma level of image")
	rootCmd.Flags().CountS("auto-level", "auto-level", "automagically adjust color levels of image")
	rootCmd.Flags().CountS("auto-orient", "auto-orient", "automagically orient image")
	rootCmd.Flags().StringSliceS("background", "background", nil, "background color")
	rootCmd.Flags().StringSliceS("bench", "bench", nil, "measure performance")
	rootCmd.Flags().StringSliceS("bias", "bias", nil, "add bias when convolving an image")
	rootCmd.Flags().StringSliceS("bilateral-blur", "bilateral-blur", nil, "non-linear, edge-preserving, and noise-reducing smoothing filter")
	rootCmd.Flags().StringSliceS("black-threshold", "black-threshold", nil, "force all pixels below the threshold into black")
	rootCmd.Flags().StringSliceS("blue-primary", "blue-primary", nil, "chromaticity blue primary point")
	rootCmd.Flags().StringSliceS("blue-shift", "blue-shift", nil, "simulate a scene at nighttime in the moonlight")
	rootCmd.Flags().StringSliceS("blur", "blur", nil, "reduce image noise and reduce detail levels")
	rootCmd.Flags().StringSliceS("border", "border", nil, "surround image with a border of color")
	rootCmd.Flags().StringSliceS("bordercolor", "bordercolor", nil, "border color")
	rootCmd.Flags().StringSliceS("brightness-contrast", "brightness-contrast", nil, "improve brightness / contrast of the image")
	rootCmd.Flags().StringSliceS("canny", "canny", nil, "use a multi-stage algorithm to detect a wide range of edges in the image")
	rootCmd.Flags().StringSliceS("caption", "caption", nil, "assign a caption to an image")
	rootCmd.Flags().StringSliceS("cdl", "cdl", nil, "color correct with a color decision list")
	rootCmd.Flags().StringSliceS("channel", "channel", nil, "apply option to select image channels")
	rootCmd.Flags().StringSliceS("charcoal", "charcoal", nil, "simulate a charcoal drawing")
	rootCmd.Flags().StringSliceS("chop", "chop", nil, "remove pixels from the image interior")
	rootCmd.Flags().StringSliceS("clahe", "clahe", nil, "contrast limited adaptive histogram equalization")
	rootCmd.Flags().CountS("clamp", "clamp", "set each pixel whose value is below zero to zero and any the pixel whose value is above the quantum range to the quantum range (e.g. 65535) otherwise the pixel value remains unchanged.")
	rootCmd.Flags().CountS("clip", "clip", "clip along the first path from the 8BIM profile")
	rootCmd.Flags().StringSliceS("clip-mask", "clip-mask", nil, "associate clip mask with the image")
	rootCmd.Flags().StringSliceS("clip-path", "clip-path", nil, "clip along a named path from the 8BIM profile")
	rootCmd.Flags().StringSliceS("clone", "clone", nil, "clone an image")
	rootCmd.Flags().CountS("clut", "clut", "apply a color lookup table to the image")
	rootCmd.Flags().CountS("coalesce", "coalesce", "merge a sequence of images")
	rootCmd.Flags().StringSliceS("color-matrix", "color-matrix", nil, "apply color correction to the image.")
	rootCmd.Flags().StringSliceS("colorize", "colorize", nil, "colorize the image with the fill color")
	rootCmd.Flags().StringSliceS("colors", "colors", nil, "preferred number of colors in the image")
	rootCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "set image colorspace")
	rootCmd.Flags().CountS("combine", "combine", "combine a sequence of images")
	rootCmd.Flags().StringSliceS("comment", "comment", nil, "annotate image with comment")
	rootCmd.Flags().CountS("compare", "compare", "compare image")
	rootCmd.Flags().StringSliceS("complex", "complex", nil, "perform complex mathematics on an image sequence")
	rootCmd.Flags().StringSliceS("compose", "compose", nil, "set image composite operator")
	rootCmd.Flags().CountS("composite", "composite", "composite image")
	rootCmd.Flags().StringSliceS("compress", "compress", nil, "image compression type")
	rootCmd.Flags().StringSliceS("connected-components", "connected-components", nil, "connected-components uniquely labeled, choose from 4 or 8 way connectivity")
	rootCmd.Flags().CountS("contrast", "contrast", "enhance or reduce the image contrast")
	rootCmd.Flags().StringSliceS("contrast-stretch", "contrast-stretch", nil, "improve the contrast in an image by `stretching' the range of intensity value")
	rootCmd.Flags().StringSliceS("convolve", "convolve", nil, "apply a convolution kernel to the image")
	rootCmd.Flags().StringSliceS("copy", "copy", nil, "copy pixels from one area of an image to another")
	rootCmd.Flags().StringSliceS("crop", "crop", nil, "crop the image")
	rootCmd.Flags().StringSliceS("cycle", "cycle", nil, "cycle the image colormap")
	rootCmd.Flags().StringSliceS("debug", "debug", nil, "display copious debugging information")
	rootCmd.Flags().StringSliceS("decipher", "decipher", nil, "convert cipher pixels to plain")
	rootCmd.Flags().CountS("deconstruct", "deconstruct", "break down an image sequence into constituent parts")
	rootCmd.Flags().StringSliceS("define", "define", nil, "define one or more image format options")
	rootCmd.Flags().StringSliceS("delay", "delay", nil, "display the next image after pausing")
	rootCmd.Flags().StringSliceS("delete", "delete", nil, "delete the image from the image sequence")
	rootCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	rootCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	rootCmd.Flags().CountS("despeckle", "despeckle", "reduce the speckles within an image")
	rootCmd.Flags().StringSliceS("direction", "direction", nil, "render text right-to-left or left-to-right")
	rootCmd.Flags().StringSliceS("display", "display", nil, "get image or font from this X server")
	rootCmd.Flags().StringSliceS("dispose", "dispose", nil, "layer disposal method")
	rootCmd.Flags().StringSliceS("distort", "distort", nil, "distort image")
	rootCmd.Flags().StringSliceS("distribute-cache", "distribute-cache", nil, "launch a distributed pixel cache server")
	rootCmd.Flags().StringSliceS("dither", "dither", nil, "apply error diffusion to image")
	rootCmd.Flags().StringSliceS("draw", "draw", nil, "annotate the image with a graphic primitive")
	rootCmd.Flags().StringSliceS("duplicate", "duplicate", nil, "duplicate an image one or more times")
	rootCmd.Flags().StringSliceS("edge", "edge", nil, "apply a filter to detect edges in the image")
	rootCmd.Flags().StringSliceS("emboss", "emboss", nil, "emboss an image")
	rootCmd.Flags().StringSliceS("encipher", "encipher", nil, "convert plain pixels to cipher pixels")
	rootCmd.Flags().StringSliceS("encoding", "encoding", nil, "text encoding type")
	rootCmd.Flags().StringSliceS("endian", "endian", nil, "endianness (MSB or LSB) of the image")
	rootCmd.Flags().CountS("enhance", "enhance", "apply a digital filter to enhance a noisy image")
	rootCmd.Flags().CountS("equalize", "equalize", "perform histogram equalization to an image")
	rootCmd.Flags().StringSliceS("evaluate", "evaluate", nil, "evaluate an arithmetic, relational, or logical expression")
	rootCmd.Flags().StringSliceS("evaluate-sequence", "evaluate-sequence", nil, "evaluate an arithmetic, relational, or logical expression for an image sequence")
	rootCmd.Flags().StringSliceS("extent", "extent", nil, "set the image size")
	rootCmd.Flags().StringSliceS("extract", "extract", nil, "extract area from image")
	rootCmd.Flags().StringSliceS("family", "family", nil, "render text with this font family")
	rootCmd.Flags().StringSliceS("features", "features", nil, "analyze image features (e.g. contract, correlations, etc.).")
	rootCmd.Flags().CountS("fft", "fft", "implements the discrete Fourier transform (DFT)")
	rootCmd.Flags().StringSliceS("fill", "fill", nil, "color to use when filling a graphic primitive")
	rootCmd.Flags().StringSliceS("filter", "filter", nil, "use this filter when resizing an image")
	rootCmd.Flags().CountS("flatten", "flatten", "flatten a sequence of images")
	rootCmd.Flags().CountS("flip", "flip", "flip image in the vertical direction")
	rootCmd.Flags().StringSliceS("floodfill", "floodfill", nil, "floodfill the image with color")
	rootCmd.Flags().CountS("flop", "flop", "flop image in the horizontal direction")
	rootCmd.Flags().StringSliceS("font", "font", nil, "render text with this font")
	rootCmd.Flags().StringSliceS("format", "format", nil, "output formatted image characteristics")
	rootCmd.Flags().StringSliceS("frame", "frame", nil, "surround image with an ornamental border")
	rootCmd.Flags().StringSliceS("function", "function", nil, "apply a function to the image")
	rootCmd.Flags().StringSliceS("fuzz", "fuzz", nil, "colors within this distance are considered equal")
	rootCmd.Flags().StringSliceS("fx", "fx", nil, "apply mathematical expression to an image channel(s)")
	rootCmd.Flags().StringSliceS("gamma", "gamma", nil, "level of gamma correction")
	rootCmd.Flags().StringSliceS("gaussian-blur", "gaussian-blur", nil, "reduce image noise and reduce detail levels")
	rootCmd.Flags().StringSliceS("geometry", "geometry", nil, "preferred size or location of the image")
	rootCmd.Flags().StringSliceS("gravity", "gravity", nil, "horizontal and vertical text placement")
	rootCmd.Flags().StringSliceS("grayscale", "grayscale", nil, "convert image to grayscale")
	rootCmd.Flags().StringSliceS("green-primary", "green-primary", nil, "chromaticity green primary point")
	rootCmd.Flags().BoolS("help", "help", false, "print program options")
	rootCmd.Flags().StringSliceS("hough-lines", "hough-lines", nil, "identify lines in the image")
	rootCmd.Flags().CountS("identify", "identify", "identify the format and characteristics of the image")
	rootCmd.Flags().CountS("ift", "ift", "implements the inverse discrete Fourier transform (DFT)")
	rootCmd.Flags().StringSliceS("illuminant", "illuminant", nil, "reference illuminant")
	rootCmd.Flags().StringSliceS("implode", "implode", nil, "implode image pixels about the center")
	rootCmd.Flags().StringSliceS("insert", "insert", nil, "insert last image into the image sequence")
	rootCmd.Flags().CountS("integral", "integral", "Calculate the sum of values (pixel values) in the image")
	rootCmd.Flags().StringSliceS("intensity", "intensity", nil, "method to generate an intensity value from a pixel")
	rootCmd.Flags().StringSliceS("intent", "intent", nil, "type of rendering intent when managing the image color")
	rootCmd.Flags().StringSliceS("interlace", "interlace", nil, "type of image interlacing scheme")
	rootCmd.Flags().StringSliceS("interline-spacing", "interline-spacing", nil, "the space between two text lines")
	rootCmd.Flags().StringSliceS("interpolate", "interpolate", nil, "pixel color interpolation method")
	rootCmd.Flags().StringSliceS("interword-spacing", "interword-spacing", nil, "the space between two words")
	rootCmd.Flags().StringSliceS("kerning", "kerning", nil, "the space between two characters")
	rootCmd.Flags().StringSliceS("kuwahara", "kuwahara", nil, "edge preserving noise reduction filter")
	rootCmd.Flags().StringSliceS("label", "label", nil, "assign a label to an image")
	rootCmd.Flags().StringSliceS("lat", "lat", nil, "local adaptive thresholding")
	rootCmd.Flags().StringSliceS("layers", "layers", nil, "optimize or compare image layers")
	rootCmd.Flags().StringSliceS("level", "level", nil, "adjust the level of image contrast")
	rootCmd.Flags().StringSliceS("limit", "limit", nil, "pixel cache resource limit")
	rootCmd.Flags().StringSliceS("linear-stretch", "linear-stretch", nil, "linear with saturation histogram stretch")
	rootCmd.Flags().StringSliceS("liquid-rescale", "liquid-rescale", nil, "rescale image with seam-carving")
	rootCmd.Flags().StringSliceS("list", "list", nil, "Color, Configure, Delegate, Format, Magic, Module, Resource, or Type")
	rootCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	rootCmd.Flags().StringSliceS("loop", "loop", nil, "add Netscape loop extension to your GIF animation")
	rootCmd.Flags().StringSliceS("mask", "mask", nil, "associate a mask with the image")
	rootCmd.Flags().StringSliceS("mattecolor", "mattecolor", nil, "frame color")
	rootCmd.Flags().StringSliceS("mean-shift", "mean-shift", nil, "delineate arbitrarily shaped clusters in the image")
	rootCmd.Flags().StringSliceS("median", "median", nil, "apply a median filter to the image")
	rootCmd.Flags().StringSliceS("metric", "metric", nil, "measure differences between images with this metric")
	rootCmd.Flags().StringSliceS("mode", "mode", nil, "make each pixel the 'predominant color' of the neighborhood")
	rootCmd.Flags().StringSliceS("modulate", "modulate", nil, "vary the brightness, saturation, and hue")
	rootCmd.Flags().CountS("moments", "moments", "display image moments.")
	rootCmd.Flags().CountS("monitor", "monitor", "monitor progress")
	rootCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	rootCmd.Flags().StringSliceS("morph", "morph", nil, "morph an image sequence")
	rootCmd.Flags().StringSliceS("morphology", "morphology", nil, "apply a morphology method to the image")
	rootCmd.Flags().StringSliceS("motion-blur", "motion-blur", nil, "simulate motion blur")
	rootCmd.Flags().CountS("negate", "negate", "replace each pixel with its complementary color")
	rootCmd.Flags().StringSliceS("noise", "noise", nil, "add or reduce noise in an image")
	rootCmd.Flags().CountS("normalize", "normalize", "transform image to span the full range of colors")
	rootCmd.Flags().StringSliceS("opaque", "opaque", nil, "change this color to the fill color")
	rootCmd.Flags().StringSliceS("ordered-dither", "ordered-dither", nil, "ordered dither the image")
	rootCmd.Flags().StringSliceS("orient", "orient", nil, "image orientation")
	rootCmd.Flags().StringSliceS("page", "page", nil, "size and location of an image canvas (setting)")
	rootCmd.Flags().StringSliceS("paint", "paint", nil, "simulate an oil painting")
	rootCmd.Flags().CountS("perceptible", "perceptible", "set each pixel whose value is less than |epsilon| to -epsilon or epsilon (whichever is closer) otherwise the pixel value remains unchanged.")
	rootCmd.Flags().CountS("ping", "ping", "efficiently determine image attributes")
	rootCmd.Flags().StringSliceS("pointsize", "pointsize", nil, "font point size")
	rootCmd.Flags().StringSliceS("polaroid", "polaroid", nil, "simulate a Polaroid picture")
	rootCmd.Flags().StringSliceS("poly", "poly", nil, "build a polynomial from the image sequence and the corresponding terms (coefficients and degree pairs).")
	rootCmd.Flags().StringSliceS("posterize", "posterize", nil, "reduce the image to a limited number of color levels")
	rootCmd.Flags().StringSliceS("precision", "precision", nil, "set the maximum number of significant digits to be printed")
	rootCmd.Flags().StringSliceS("preview", "preview", nil, "image preview type")
	rootCmd.Flags().StringSliceS("print", "print", nil, "interpret string and print to console")
	rootCmd.Flags().StringSliceS("process", "process", nil, "process the image with a custom image filter")
	rootCmd.Flags().StringSliceS("profile", "profile", nil, "add, delete, or apply an image profile")
	rootCmd.Flags().StringSliceS("quality", "quality", nil, "JPEG/MIFF/PNG compression level")
	rootCmd.Flags().StringSliceS("quantize", "quantize", nil, "reduce image colors in this colorspace")
	rootCmd.Flags().CountS("quiet", "quiet", "suppress all warning messages")
	rootCmd.Flags().StringSliceS("raise", "raise", nil, "lighten/darken image edges to create a 3-D effect")
	rootCmd.Flags().StringSliceS("random-threshold", "random-threshold", nil, "random threshold the image")
	rootCmd.Flags().StringSliceS("red-primary", "red-primary", nil, "chromaticity red primary point")
	rootCmd.Flags().CountS("regard-warnings", "regard-warnings", "pay attention to warning messages.")
	rootCmd.Flags().StringSliceS("region", "region", nil, "apply options to a portion of the image")
	rootCmd.Flags().StringSliceS("remap", "remap", nil, "transform image colors to match this set of colors")
	rootCmd.Flags().CountS("render", "render", "render vector graphics")
	rootCmd.Flags().StringSliceS("repage", "repage", nil, "size and location of an image canvas")
	rootCmd.Flags().StringSliceS("resample", "resample", nil, "change the resolution of an image")
	rootCmd.Flags().StringSliceS("reshape", "reshape", nil, "reshape the image")
	rootCmd.Flags().StringSliceS("resize", "resize", nil, "resize the image")
	rootCmd.Flags().CountS("respect-parentheses", "respect-parentheses", "settings remain in effect until parenthesis boundary.")
	rootCmd.Flags().StringSliceS("roll", "roll", nil, "roll an image vertically or horizontally")
	rootCmd.Flags().StringSliceS("rotate", "rotate", nil, "apply Paeth rotation to the image")
	rootCmd.Flags().StringSliceS("rotational-blur", "rotational-blur", nil, "radial blur the image")
	rootCmd.Flags().StringSliceS("sample", "sample", nil, "scale image with pixel sampling")
	rootCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factor")
	rootCmd.Flags().StringSliceS("scale", "scale", nil, "scale the image")
	rootCmd.Flags().StringSliceS("scene", "scene", nil, "image scene number")
	rootCmd.Flags().StringSliceS("seed", "seed", nil, "seed a new sequence of pseudo-random numbers")
	rootCmd.Flags().StringSliceS("segment", "segment", nil, "segment an image")
	rootCmd.Flags().StringSliceS("selective-blur", "selective-blur", nil, "selectively blur pixels within a contrast threshold")
	rootCmd.Flags().CountS("separate", "separate", "separate an image channel into a grayscale image")
	rootCmd.Flags().StringSliceS("sepia-tone", "sepia-tone", nil, "simulate a sepia-toned photo")
	rootCmd.Flags().StringSliceS("set", "set", nil, "set an image attribute")
	rootCmd.Flags().StringSliceS("shade", "shade", nil, "shade the image using a distant light source")
	rootCmd.Flags().StringSliceS("shadow", "shadow", nil, "simulate an image shadow")
	rootCmd.Flags().StringSliceS("sharpen", "sharpen", nil, "sharpen the image")
	rootCmd.Flags().StringSliceS("shave", "shave", nil, "shave pixels from the image edges")
	rootCmd.Flags().StringSliceS("shear", "shear", nil, "slide one edge of the image along the X or Y axis")
	rootCmd.Flags().StringSliceS("sigmoidal-contrast", "sigmoidal-contrast", nil, "increase the contrast without saturating highlights or shadows")
	rootCmd.Flags().StringSliceS("size", "size", nil, "width and height of image")
	rootCmd.Flags().StringSliceS("sketch", "sketch", nil, "simulate a pencil sketch")
	rootCmd.Flags().StringSliceS("smush", "smush", nil, "smush an image sequence together")
	rootCmd.Flags().StringSliceS("solarize", "solarize", nil, "negate all pixels above the threshold level")
	rootCmd.Flags().CountS("sort-pixels", "sort-pixels", "sorts pixels within each scanline in ascending order of intensity")
	rootCmd.Flags().StringSliceS("splice", "splice", nil, "splice the background color into the image")
	rootCmd.Flags().StringSliceS("spread", "spread", nil, "displace image pixels by a random amount")
	rootCmd.Flags().StringSliceS("statistic", "statistic", nil, "replace each pixel with corresponding statistic from the neighborhood")
	rootCmd.Flags().StringSliceS("stretch", "stretch", nil, "render text with this font stretch")
	rootCmd.Flags().CountS("strip", "strip", "strip image of all profiles and comments")
	rootCmd.Flags().StringSliceS("stroke", "stroke", nil, "graphic primitive stroke color")
	rootCmd.Flags().StringSliceS("strokewidth", "strokewidth", nil, "graphic primitive stroke width")
	rootCmd.Flags().StringSliceS("style", "style", nil, "render text with this font style")
	rootCmd.Flags().StringSliceS("swap", "swap", nil, "swap two images in the image sequence")
	rootCmd.Flags().StringSliceS("swirl", "swirl", nil, "swirl image pixels about the center")
	rootCmd.Flags().CountS("synchronize", "synchronize", "synchronize image to storage device")
	rootCmd.Flags().CountS("taint", "taint", "mark the image as modified")
	rootCmd.Flags().StringSliceS("texture", "texture", nil, "name of texture to tile onto the image background")
	rootCmd.Flags().StringSliceS("threshold", "threshold", nil, "threshold the image")
	rootCmd.Flags().StringSliceS("thumbnail", "thumbnail", nil, "create a thumbnail of the image")
	rootCmd.Flags().StringSliceS("tile", "tile", nil, "tile image when filling a graphic primitive")
	rootCmd.Flags().StringSliceS("tile-offset", "tile-offset", nil, "set the image tile offset")
	rootCmd.Flags().StringSliceS("tint", "tint", nil, "tint the image with the fill color")
	rootCmd.Flags().CountS("transform", "transform", "affine transform image")
	rootCmd.Flags().StringSliceS("transparent", "transparent", nil, "make this color transparent within the image")
	rootCmd.Flags().StringSliceS("transparent-color", "transparent-color", nil, "transparent color")
	rootCmd.Flags().CountS("transpose", "transpose", "flip image in the vertical direction and rotate 90 degrees")
	rootCmd.Flags().CountS("transverse", "transverse", "flop image in the horizontal direction and rotate 270 degrees")
	rootCmd.Flags().StringSliceS("treedepth", "treedepth", nil, "color tree depth")
	rootCmd.Flags().CountS("trim", "trim", "trim image edges")
	rootCmd.Flags().StringSliceS("type", "type", nil, "image type")
	rootCmd.Flags().StringSliceS("undercolor", "undercolor", nil, "annotation bounding box color")
	rootCmd.Flags().CountS("unique-colors", "unique-colors", "discard all but one of any pixel color.")
	rootCmd.Flags().StringSliceS("units", "units", nil, "the units of image resolution")
	rootCmd.Flags().StringSliceS("unsharp", "unsharp", nil, "sharpen the image")
	rootCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	rootCmd.Flags().BoolS("version", "version", false, "print version information")
	rootCmd.Flags().CountS("view", "view", "FlashPix viewing transforms")
	rootCmd.Flags().StringSliceS("vignette", "vignette", nil, "soften the edges of the image in vignette style")
	rootCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", nil, "access method for pixels outside the boundaries of the image")
	rootCmd.Flags().StringSliceS("wave", "wave", nil, "alter an image along a sine wave")
	rootCmd.Flags().StringSliceS("wavelet-denoise", "wavelet-denoise", nil, "removes noise from the image using a wavelet transform")
	rootCmd.Flags().StringSliceS("weight", "weight", nil, "render text with this font weight")
	rootCmd.Flags().StringSliceS("white-point", "white-point", nil, "chromaticity white point")
	rootCmd.Flags().StringSliceS("white-threshold", "white-threshold", nil, "force all pixels above the threshold into white")
	rootCmd.Flags().StringSliceS("word-break", "word-break", nil, "sets whether line breaks appear wherever the text would otherwise overflow its content box. Choose from normal, the default, or break-word.")
	rootCmd.Flags().StringSliceS("write", "write", nil, "write images to this file")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"background":        action.ActionColors(),
		"bordercolor":       action.ActionColors(),
		"cdl":               carapace.ActionFiles(),
		"channel":           action.ActionCompleteOption("Channel").UniqueList(","),
		"clip-mask":         carapace.ActionFiles(),
		"colorspace":        action.ActionCompleteOption("Colorspace"),
		"complex":           action.ActionCompleteOption("Complex"),
		"compose":           action.ActionCompleteOption("Compose"),
		"compress":          action.ActionCompleteOption("Compress"),
		"debug":             action.ActionCompleteOption("Debug").UniqueList(","),
		"decipher":          carapace.ActionFiles(),
		"direction":         action.ActionCompleteOption("Direction"),
		"dispose":           action.ActionCompleteOption("Dispose"),
		"distort":           action.ActionCompleteOption("Distort"),
		"dither":            action.ActionCompleteOption("Dither"),
		"encipher":          carapace.ActionFiles(),
		"encoding":          carapace.ActionValues("AdobeCustom", "AdobeExpert", "AdobeStandard", "AppleRoman", "BIG5", "GB2312", "Latin 2", "None", "SJIScode", "Symbol", "Unicode", "Wansung"),
		"endian":            action.ActionCompleteOption("Endian"),
		"evaluate":          action.ActionCompleteOption("Evaluate"),
		"evaluate-sequence": action.ActionCompleteOption("Evaluate"),
		"family":            action.ActionFontFamilies(),
		"fill":              action.ActionColors(),
		"filter":            action.ActionCompleteOption("Filter"),
		"font":              action.ActionFonts(),
		"function":          action.ActionCompleteOption("Function"),
		"gravity":           action.ActionCompleteOption("Gravity"),
		"grayscale":         action.ActionCompleteOption("Intensity"),
		"illuminant":        action.ActionCompleteOption("Illuminant"),
		"intensity":         action.ActionCompleteOption("Intensity"),
		"intent":            action.ActionCompleteOption("Intent"),
		"interlace":         action.ActionCompleteOption("Interlace"),
		"interpolate":       action.ActionCompleteOption("Interpolate"),
		"layers":            action.ActionCompleteOption("Layers"),
		"list":              action.ActionCompleteOption("List"),
		"mask":              carapace.ActionFiles(),
		"mattecolor":        action.ActionColors(),
		"metric":            action.ActionCompleteOption("Metric"),
		"mode":              action.ActionCompleteOption("Mode"),
		"morphology":        action.ActionCompleteOption("Morphology"),
		"opaque":            action.ActionColors(),
		"orient":            action.ActionCompleteOption("Orientation"),
		"preview":           action.ActionCompleteOption("Preview"),
		"process":           action.ActionCompleteOption("Filter"),
		"profile":           carapace.ActionFiles(),
		"quantize":          action.ActionCompleteOption("Colorspace"),
		"remap":             carapace.ActionFiles(),
		"statistic":         action.ActionCompleteOption("Statistic"),
		"stretch":           action.ActionCompleteOption("Stretch"),
		"stroke":            action.ActionColors(),
		"style":             action.ActionCompleteOption("Style"),
		"texture":           carapace.ActionFiles(),
		"tile":              carapace.ActionFiles(),
		"transparent":       action.ActionColors(),
		"transparent-color": action.ActionColors(),
		"type":              action.ActionCompleteOption("Type"),
		"undercolor":        action.ActionColors(),
		"units":             action.ActionCompleteOption("Units"),
		"virtual-pixel":     action.ActionCompleteOption("VirtualPixel"),
		"weight":            action.ActionCompleteOption("Weight"),
		"word-break":        action.ActionCompleteOption("WordBreak"),
		"write":             carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
