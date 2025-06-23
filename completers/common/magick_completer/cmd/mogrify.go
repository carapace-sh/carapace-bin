package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/magick_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mogrifyCmd = &cobra.Command{
	Use:   "mogrify",
	Short: "resize an image, blur, crop, despeckle, dither, draw on, flip, join, re-sample, and much more. Mogrify overwrites the original image file, whereas, magick writes to a different image file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mogrifyCmd).Standalone()

	mogrifyCmd.Flags().StringSliceS("adaptive-blur", "adaptive-blur", nil, "adaptively blur pixels; decrease effect near edges")
	mogrifyCmd.Flags().StringSliceS("adaptive-resize", "adaptive-resize", nil, "adaptively resize image with data dependent triangulation.")
	mogrifyCmd.Flags().StringSliceS("adaptive-sharpen", "adaptive-sharpen", nil, "adaptively sharpen pixels; increase effect near edges")
	mogrifyCmd.Flags().CountS("adjoin", "adjoin", "join images into a single multi-image file")
	mogrifyCmd.Flags().StringSliceS("affine", "affine", nil, "affine transform matrix")
	mogrifyCmd.Flags().CountS("alpha", "alpha", "on, activate, off, deactivate, set, opaque, copy\", transparent, extract, background, or shape the alpha channel")
	mogrifyCmd.Flags().StringSliceS("annotate", "annotate", nil, "annotate the image with text")
	mogrifyCmd.Flags().CountS("antialias", "antialias", "remove pixel-aliasing")
	mogrifyCmd.Flags().CountS("append", "append", "append an image sequence")
	mogrifyCmd.Flags().StringSliceS("authenticate", "authenticate", nil, "decipher image with this password")
	mogrifyCmd.Flags().CountS("auto-gamma", "auto-gamma", "automagically adjust gamma level of image")
	mogrifyCmd.Flags().CountS("auto-level", "auto-level", "automagically adjust color levels of image")
	mogrifyCmd.Flags().CountS("auto-orient", "auto-orient", "automagically orient image")
	mogrifyCmd.Flags().StringSliceS("auto-threshold", "auto-threshold", nil, "automatically perform image thresholding")
	mogrifyCmd.Flags().StringSliceS("background", "background", nil, "background color")
	mogrifyCmd.Flags().StringSliceS("bench", "bench", nil, "measure performance")
	mogrifyCmd.Flags().StringSliceS("bias", "bias", nil, "add bias when convolving an image")
	mogrifyCmd.Flags().StringSliceS("bilateral-blur", "bilateral-blur", nil, "non-linear, edge-preserving, and noise-reducing smoothing filter")
	mogrifyCmd.Flags().StringSliceS("black-threshold", "black-threshold", nil, "force all pixels below the threshold into black")
	mogrifyCmd.Flags().StringSliceS("blue-primary", "blue-primary", nil, "chromaticity blue primary point")
	mogrifyCmd.Flags().StringSliceS("blue-shift", "blue-shift", nil, "simulate a scene at nighttime in the moonlight")
	mogrifyCmd.Flags().StringSliceS("blur", "blur", nil, "reduce image noise and reduce detail levels")
	mogrifyCmd.Flags().StringSliceS("border", "border", nil, "surround image with a border of color")
	mogrifyCmd.Flags().StringSliceS("bordercolor", "bordercolor", nil, "border color")
	mogrifyCmd.Flags().StringSliceS("brightness-contrast", "brightness-contrast", nil, "improve brightness / contrast of the image")
	mogrifyCmd.Flags().StringSliceS("canny", "canny", nil, "use a multi-stage algorithm to detect a wide range of edges in the image")
	mogrifyCmd.Flags().StringSliceS("caption", "caption", nil, "assign a caption to an image")
	mogrifyCmd.Flags().StringSliceS("cdl", "cdl", nil, "color correct with a color decision list")
	mogrifyCmd.Flags().StringSliceS("channel", "channel", nil, "apply option to select image channels")
	mogrifyCmd.Flags().StringSliceS("charcoal", "charcoal", nil, "simulate a charcoal drawing")
	mogrifyCmd.Flags().StringSliceS("chop", "chop", nil, "remove pixels from the image interior")
	mogrifyCmd.Flags().StringSliceS("clahe", "clahe", nil, "contrast limited adaptive histogram equalization")
	mogrifyCmd.Flags().CountS("clamp", "clamp", "set each pixel whose value is below zero to zero and any the pixel whose value is above the quantum range to the quantum range (e.g. 65535) otherwise the pixel value remains unchanged.")
	mogrifyCmd.Flags().CountS("clip", "clip", "clip along the first path from the 8BIM profile")
	mogrifyCmd.Flags().CountS("clip-mask", "clip-mask", "associate clip mask with the image")
	mogrifyCmd.Flags().StringSliceS("clip-path", "clip-path", nil, "clip along a named path from the 8BIM profile")
	mogrifyCmd.Flags().CountS("clut", "clut", "apply a color lookup table to the image")
	mogrifyCmd.Flags().CountS("coalesce", "coalesce", "merge a sequence of images")
	mogrifyCmd.Flags().StringSliceS("color-matrix", "color-matrix", nil, "apply color correction to the image.")
	mogrifyCmd.Flags().StringSliceS("color-threshold", "color-threshold", nil, "force all pixels in the color range to white otherwise black")
	mogrifyCmd.Flags().StringSliceS("colorize", "colorize", nil, "colorize the image with the fill color")
	mogrifyCmd.Flags().StringSliceS("colors", "colors", nil, "preferred number of colors in the image")
	mogrifyCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "set image colorspace")
	mogrifyCmd.Flags().CountS("combine", "combine", "combine a sequence of images")
	mogrifyCmd.Flags().StringSliceS("comment", "comment", nil, "annotate image with comment")
	mogrifyCmd.Flags().CountS("complex", "complex", "perform complex mathematics on an image sequence")
	mogrifyCmd.Flags().StringSliceS("compose", "compose", nil, "set image composite operator")
	mogrifyCmd.Flags().CountS("composite", "composite", "composite image")
	mogrifyCmd.Flags().StringSliceS("compress", "compress", nil, "image compression type")
	mogrifyCmd.Flags().StringSliceS("connected-components", "connected-components", nil, "connected-components uniquely labeled, choose from 4 or 8 way connectivity")
	mogrifyCmd.Flags().CountS("contrast", "contrast", "enhance or reduce the image contrast")
	mogrifyCmd.Flags().StringSliceS("contrast-stretch", "contrast-stretch", nil, "improve the contrast in an image by `stretching' the range of intensity value")
	mogrifyCmd.Flags().StringSliceS("convolve", "convolve", nil, "apply a convolution kernel to the image")
	mogrifyCmd.Flags().StringSliceS("copy", "copy", nil, "copy pixels from one area of an image to another")
	mogrifyCmd.Flags().StringSliceS("crop", "crop", nil, "crop the image")
	mogrifyCmd.Flags().StringSliceS("cycle", "cycle", nil, "cycle the image colormap")
	mogrifyCmd.Flags().StringSliceS("debug", "debug", nil, "display copious debugging information")
	mogrifyCmd.Flags().StringSliceS("decipher", "decipher", nil, "convert cipher pixels to plain")
	mogrifyCmd.Flags().CountS("deconstruct", "deconstruct", "break down an image sequence into constituent parts")
	mogrifyCmd.Flags().StringSliceS("define", "define", nil, "define one or more image format options")
	mogrifyCmd.Flags().StringSliceS("delay", "delay", nil, "display the next image after pausing")
	mogrifyCmd.Flags().StringSliceS("delete", "delete", nil, "delete the image from the image sequence")
	mogrifyCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	mogrifyCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	mogrifyCmd.Flags().CountS("despeckle", "despeckle", "reduce the speckles within an image")
	mogrifyCmd.Flags().StringSliceS("direction", "direction", nil, "render text right-to-left or left-to-right")
	mogrifyCmd.Flags().StringSliceS("display", "display", nil, "get image or font from this X server")
	mogrifyCmd.Flags().StringSliceS("dispose", "dispose", nil, "layer disposal method")
	mogrifyCmd.Flags().StringSliceS("distort", "distort", nil, "distort image")
	mogrifyCmd.Flags().StringSliceS("distribute-cache", "distribute-cache", nil, "launch a pixel cache server")
	mogrifyCmd.Flags().StringSliceS("dither", "dither", nil, "apply error diffusion to image")
	mogrifyCmd.Flags().StringSliceS("draw", "draw", nil, "annotate the image with a graphic primitive")
	mogrifyCmd.Flags().StringSliceS("duplicate", "duplicate", nil, "duplicate an image one or more times")
	mogrifyCmd.Flags().StringSliceS("edge", "edge", nil, "apply a filter to detect edges in the image")
	mogrifyCmd.Flags().StringSliceS("emboss", "emboss", nil, "emboss an image")
	mogrifyCmd.Flags().StringSliceS("encipher", "encipher", nil, "convert plain pixels to cipher pixels")
	mogrifyCmd.Flags().StringSliceS("encoding", "encoding", nil, "text encoding type")
	mogrifyCmd.Flags().StringSliceS("endian", "endian", nil, "endianness (MSB or LSB) of the image")
	mogrifyCmd.Flags().CountS("enhance", "enhance", "apply a digital filter to enhance a noisy image")
	mogrifyCmd.Flags().CountS("equalize", "equalize", "perform histogram equalization to an image")
	mogrifyCmd.Flags().StringSliceS("evaluate", "evaluate", nil, "evaluate an arithmetic, relational, or logical expression")
	mogrifyCmd.Flags().StringSliceS("evaluate-sequence", "evaluate-sequence", nil, "evaluate an arithmetic, relational, or logical expression for an image sequence")
	mogrifyCmd.Flags().StringSliceS("extent", "extent", nil, "set the image size")
	mogrifyCmd.Flags().StringSliceS("extract", "extract", nil, "extract area from image")
	mogrifyCmd.Flags().StringSliceS("family", "family", nil, "render text with this font family")
	mogrifyCmd.Flags().StringSliceS("features", "features", nil, "analyze image features (e.g. contract, correlations, etc.).")
	mogrifyCmd.Flags().CountS("fft", "fft", "implements the discrete Fourier transform (DFT)")
	mogrifyCmd.Flags().StringSliceS("fill", "fill", nil, "color to use when filling a graphic primitive")
	mogrifyCmd.Flags().StringSliceS("filter", "filter", nil, "use this filter when resizing an image")
	mogrifyCmd.Flags().CountS("flatten", "flatten", "flatten a sequence of images")
	mogrifyCmd.Flags().CountS("flip", "flip", "flip image in the vertical direction")
	mogrifyCmd.Flags().StringSliceS("floodfill", "floodfill", nil, "floodfill the image with color")
	mogrifyCmd.Flags().CountS("flop", "flop", "flop image in the horizontal direction")
	mogrifyCmd.Flags().StringSliceS("font", "font", nil, "render text with this font")
	mogrifyCmd.Flags().StringSliceS("format", "format", nil, "output formatted image characteristics")
	mogrifyCmd.Flags().StringSliceS("frame", "frame", nil, "surround image with an ornamental border")
	mogrifyCmd.Flags().StringSliceS("function", "function", nil, "apply a function to the image")
	mogrifyCmd.Flags().StringSliceS("fuzz", "fuzz", nil, "colors within this distance are considered equal")
	mogrifyCmd.Flags().StringSliceS("fx", "fx", nil, "apply mathematical expression to an image channel(s)")
	mogrifyCmd.Flags().StringSliceS("gamma", "gamma", nil, "level of gamma correction")
	mogrifyCmd.Flags().StringSliceS("gaussian-blur", "gaussian-blur", nil, "reduce image noise and reduce detail levels")
	mogrifyCmd.Flags().StringSliceS("geometry", "geometry", nil, "preferred size or location of the image")
	mogrifyCmd.Flags().StringSliceS("gravity", "gravity", nil, "horizontal and vertical text placement")
	mogrifyCmd.Flags().StringSliceS("grayscale", "grayscale", nil, "convert image to grayscale")
	mogrifyCmd.Flags().StringSliceS("green-primary", "green-primary", nil, "chromaticity green primary point")
	mogrifyCmd.Flags().BoolS("help", "help", false, "print program options")
	mogrifyCmd.Flags().StringSliceS("hough-lines", "hough-lines", nil, "identify lines in the image")
	mogrifyCmd.Flags().CountS("identify", "identify", "identify the format and characteristics of the image")
	mogrifyCmd.Flags().CountS("ifft", "ifft", "implements the inverse discrete Fourier transform (DFT)")
	mogrifyCmd.Flags().StringSliceS("illuminant", "illuminant", nil, "reference illuminant")
	mogrifyCmd.Flags().StringSliceS("implode", "implode", nil, "implode image pixels about the center")
	mogrifyCmd.Flags().StringSliceS("insert", "insert", nil, "insert last image into the image sequence")
	mogrifyCmd.Flags().CountS("integral", "integral", "Calculate the sum of values (pixel values) in the image")
	mogrifyCmd.Flags().StringSliceS("intensity", "intensity", nil, "method to generate an intensity value from a pixel")
	mogrifyCmd.Flags().StringSliceS("intent", "intent", nil, "type of rendering intent when managing the image color")
	mogrifyCmd.Flags().StringSliceS("interlace", "interlace", nil, "type of image interlacing scheme")
	mogrifyCmd.Flags().StringSliceS("interline-spacing", "interline-spacing", nil, "the space between two text lines")
	mogrifyCmd.Flags().StringSliceS("interpolate", "interpolate", nil, "pixel color interpolation method")
	mogrifyCmd.Flags().StringSliceS("interword-spacing", "interword-spacing", nil, "the space between two words")
	mogrifyCmd.Flags().StringSliceS("kerning", "kerning", nil, "the space between two characters")
	mogrifyCmd.Flags().StringSliceS("kmeans", "kmeans", nil, "K means color reduction")
	mogrifyCmd.Flags().StringSliceS("kuwahara", "kuwahara", nil, "edge preserving noise reduction filter")
	mogrifyCmd.Flags().StringSliceS("label", "label", nil, "assign a label to an image")
	mogrifyCmd.Flags().StringSliceS("lat", "lat", nil, "local adaptive thresholding")
	mogrifyCmd.Flags().StringSliceS("layers", "layers", nil, "optimize or compare image layers")
	mogrifyCmd.Flags().StringSliceS("level", "level", nil, "adjust the level of image contrast")
	mogrifyCmd.Flags().StringSliceS("limit", "limit", nil, "pixel cache resource limit")
	mogrifyCmd.Flags().StringSliceS("linear-stretch", "linear-stretch", nil, "linear with saturation histogram stretch")
	mogrifyCmd.Flags().StringSliceS("liquid-rescale", "liquid-rescale", nil, "rescale image with seam-carving")
	mogrifyCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	mogrifyCmd.Flags().StringSliceS("loop", "loop", nil, "add Netscape loop extension to your GIF animation")
	mogrifyCmd.Flags().StringSliceS("mattecolor", "mattecolor", nil, "frame color")
	mogrifyCmd.Flags().StringSliceS("mean-shift", "mean-shift", nil, "delineate arbitrarily shaped clusters in the image")
	mogrifyCmd.Flags().StringSliceS("median", "median", nil, "apply a median filter to the image")
	mogrifyCmd.Flags().StringSliceS("metric", "metric", nil, "measure differences between images with this metric")
	mogrifyCmd.Flags().StringSliceS("mode", "mode", nil, "make each pixel the 'predominant color' of the neighborhood")
	mogrifyCmd.Flags().StringSliceS("modulate", "modulate", nil, "vary the brightness, saturation, and hue")
	mogrifyCmd.Flags().CountS("monitor", "monitor", "monitor progress")
	mogrifyCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	mogrifyCmd.Flags().StringSliceS("morph", "morph", nil, "morph an image sequence")
	mogrifyCmd.Flags().StringSliceS("morphology", "morphology", nil, "apply a morphology method to the image")
	mogrifyCmd.Flags().StringSliceS("motion-blur", "motion-blur", nil, "simulate motion blur")
	mogrifyCmd.Flags().CountS("negate", "negate", "replace each pixel with its complementary color")
	mogrifyCmd.Flags().StringSliceS("noise", "noise", nil, "add or reduce noise in an image")
	mogrifyCmd.Flags().CountS("normalize", "normalize", "transform image to span the full range of colors")
	mogrifyCmd.Flags().StringSliceS("opaque", "opaque", nil, "change this color to the fill color")
	mogrifyCmd.Flags().StringSliceS("ordered-dither", "ordered-dither", nil, "ordered dither the image")
	mogrifyCmd.Flags().StringSliceS("orient", "orient", nil, "image orientation")
	mogrifyCmd.Flags().StringSliceS("page", "page", nil, "size and location of an image canvas (setting)")
	mogrifyCmd.Flags().StringSliceS("paint", "paint", nil, "simulate an oil painting")
	mogrifyCmd.Flags().StringSliceS("path", "path", nil, "write images to this path on disk")
	mogrifyCmd.Flags().CountS("perceptible", "perceptible", "set each pixel whose value is less than |")
	mogrifyCmd.Flags().CountS("ping", "ping", "efficiently determine image attributes")
	mogrifyCmd.Flags().StringSliceS("pointsize", "pointsize", nil, "font point size")
	mogrifyCmd.Flags().StringSliceS("polaroid", "polaroid", nil, "simulate a Polaroid picture")
	mogrifyCmd.Flags().StringSliceS("poly", "poly", nil, "build a polynomial from the image sequence and the corresponding terms (coefficients and degree pairs).")
	mogrifyCmd.Flags().StringSliceS("posterize", "posterize", nil, "reduce the image to a limited number of color levels")
	mogrifyCmd.Flags().StringSliceS("precision", "precision", nil, "set the maximum number of significant digits to be printed")
	mogrifyCmd.Flags().StringSliceS("preview", "preview", nil, "image preview type")
	mogrifyCmd.Flags().StringSliceS("print", "print", nil, "interpret string and print to console")
	mogrifyCmd.Flags().StringSliceS("process", "process", nil, "process the image with a custom image filter")
	mogrifyCmd.Flags().StringSliceS("profile", "profile", nil, "add, delete, or apply an image profile")
	mogrifyCmd.Flags().StringSliceS("quality", "quality", nil, "JPEG/MIFF/PNG compression level")
	mogrifyCmd.Flags().StringSliceS("quantize", "quantize", nil, "reduce image colors in this colorspace")
	mogrifyCmd.Flags().CountS("quiet", "quiet", "suppress all warning messages")
	mogrifyCmd.Flags().StringSliceS("raise", "raise", nil, "lighten/darken image edges to create a 3-D effect")
	mogrifyCmd.Flags().StringSliceS("random-threshold", "random-threshold", nil, "random threshold the image")
	mogrifyCmd.Flags().StringSliceS("range-threshold", "range-threshold", nil, "perform either hard or soft thresholding within some range of values in an image")
	mogrifyCmd.Flags().StringSliceS("read-mask", "read-mask", nil, "associate a read mask with the image")
	mogrifyCmd.Flags().StringSliceS("red-primary", "red-primary", nil, "chromaticity red primary point")
	mogrifyCmd.Flags().CountS("regard-warnings", "regard-warnings", "pay attention to warning messages.")
	mogrifyCmd.Flags().StringSliceS("region", "region", nil, "apply options to a portion of the image")
	mogrifyCmd.Flags().StringSliceS("remap", "remap", nil, "transform image colors to match this set of colors")
	mogrifyCmd.Flags().CountS("render", "render", "render vector graphics")
	mogrifyCmd.Flags().StringSliceS("repage", "repage", nil, "size and location of an image canvas")
	mogrifyCmd.Flags().StringSliceS("resample", "resample", nil, "change the resolution of an image")
	mogrifyCmd.Flags().StringSliceS("reshape", "reshape", nil, "reshape the image")
	mogrifyCmd.Flags().StringSliceS("resize", "resize", nil, "resize the image")
	mogrifyCmd.Flags().CountS("respect-parentheses", "respect-parentheses", "settings remain in effect until parenthesis boundary.")
	mogrifyCmd.Flags().StringSliceS("roll", "roll", nil, "roll an image vertically or horizontally")
	mogrifyCmd.Flags().StringSliceS("rotate", "rotate", nil, "apply Paeth rotation to the image")
	mogrifyCmd.Flags().StringSliceS("rotational-blur", "rotational-blur", nil, "radial blur the image")
	mogrifyCmd.Flags().StringSliceS("sample", "sample", nil, "scale image with pixel sampling")
	mogrifyCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factor")
	mogrifyCmd.Flags().StringSliceS("scale", "scale", nil, "scale the image")
	mogrifyCmd.Flags().StringSliceS("scene", "scene", nil, "image scene number")
	mogrifyCmd.Flags().StringSliceS("seed", "seed", nil, "seed a new sequence of pseudo-random numbers")
	mogrifyCmd.Flags().StringSliceS("segment", "segment", nil, "segment an image")
	mogrifyCmd.Flags().StringSliceS("selective-blur", "selective-blur", nil, "selectively blur pixels within a contrast threshold")
	mogrifyCmd.Flags().CountS("separate", "separate", "separate an image channel into a grayscale image")
	mogrifyCmd.Flags().StringSliceS("sepia-tone", "sepia-tone", nil, "simulate a sepia-toned photo")
	mogrifyCmd.Flags().StringSliceS("set", "set", nil, "set an image attribute")
	mogrifyCmd.Flags().StringSliceS("shade", "shade", nil, "shade the image using a distant light source")
	mogrifyCmd.Flags().StringSliceS("shadow", "shadow", nil, "simulate an image shadow")
	mogrifyCmd.Flags().StringSliceS("sharpen", "sharpen", nil, "sharpen the image")
	mogrifyCmd.Flags().StringSliceS("shave", "shave", nil, "shave pixels from the image edges")
	mogrifyCmd.Flags().StringSliceS("shear", "shear", nil, "slide one edge of the image along the X or Y axis")
	mogrifyCmd.Flags().StringSliceS("sigmoidal-contrast", "sigmoidal-contrast", nil, "increase the contrast without saturating highlights or shadows")
	mogrifyCmd.Flags().StringSliceS("size", "size", nil, "width and height of image")
	mogrifyCmd.Flags().StringSliceS("sketch", "sketch", nil, "simulate a pencil sketch")
	mogrifyCmd.Flags().StringSliceS("smush", "smush", nil, "smush an image sequence together")
	mogrifyCmd.Flags().StringSliceS("solarize", "solarize", nil, "negate all pixels above the threshold level")
	mogrifyCmd.Flags().CountS("sort-pixels", "sort-pixels", "sorts pixels within each scanline in ascending order of intensity")
	mogrifyCmd.Flags().StringSliceS("splice", "splice", nil, "splice the background color into the image")
	mogrifyCmd.Flags().StringSliceS("spread", "spread", nil, "displace image pixels by a random amount")
	mogrifyCmd.Flags().StringSliceS("statistic", "statistic", nil, "replace each pixel with corresponding statistic from the neighborhood")
	mogrifyCmd.Flags().StringSliceS("stretch", "stretch", nil, "render text with this font stretch")
	mogrifyCmd.Flags().CountS("strip", "strip", "strip image of all profiles and comments")
	mogrifyCmd.Flags().StringSliceS("stroke", "stroke", nil, "graphic primitive stroke color")
	mogrifyCmd.Flags().StringSliceS("strokewidth", "strokewidth", nil, "graphic primitive stroke width")
	mogrifyCmd.Flags().StringSliceS("style", "style", nil, "render text with this font style")
	mogrifyCmd.Flags().StringSliceS("swap", "swap", nil, "swap two images in the image sequence")
	mogrifyCmd.Flags().StringSliceS("swirl", "swirl", nil, "swirl image pixels about the center")
	mogrifyCmd.Flags().CountS("synchronize", "synchronize", "synchronize image to storage device")
	mogrifyCmd.Flags().StringSliceS("texture", "texture", nil, "name of texture to tile onto the image background")
	mogrifyCmd.Flags().StringSliceS("threshold", "threshold", nil, "threshold the image")
	mogrifyCmd.Flags().StringSliceS("thumbnail", "thumbnail", nil, "create a thumbnail of the image")
	mogrifyCmd.Flags().StringSliceS("tile", "tile", nil, "tile image when filling a graphic primitive")
	mogrifyCmd.Flags().StringSliceS("tile-offset", "tile-offset", nil, "set the image tile offset")
	mogrifyCmd.Flags().StringSliceS("tint", "tint", nil, "tint the image with the fill color")
	mogrifyCmd.Flags().CountS("transform", "transform", "affine transform image")
	mogrifyCmd.Flags().StringSliceS("transparent", "transparent", nil, "make this color transparent within the image")
	mogrifyCmd.Flags().StringSliceS("transparent-color", "transparent-color", nil, "transparent color")
	mogrifyCmd.Flags().CountS("transpose", "transpose", "flip image in the vertical direction and rotate 90 degrees")
	mogrifyCmd.Flags().CountS("transverse", "transverse", "flop image in the horizontal direction and rotate 270 degrees")
	mogrifyCmd.Flags().StringSliceS("treedepth", "treedepth", nil, "color tree depth")
	mogrifyCmd.Flags().CountS("trim", "trim", "trim image edges")
	mogrifyCmd.Flags().StringSliceS("type", "type", nil, "image type")
	mogrifyCmd.Flags().StringSliceS("undercolor", "undercolor", nil, "annotation bounding box color")
	mogrifyCmd.Flags().CountS("unique-colors", "unique-colors", "discard all but one of any pixel color.")
	mogrifyCmd.Flags().StringSliceS("units", "units", nil, "the units of image resolution")
	mogrifyCmd.Flags().StringSliceS("unsharp", "unsharp", nil, "sharpen the image")
	mogrifyCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	mogrifyCmd.Flags().BoolS("version", "version", false, "print version information")
	mogrifyCmd.Flags().CountS("view", "view", "FlashPix viewing transforms")
	mogrifyCmd.Flags().StringSliceS("vignette", "vignette", nil, "soften the edges of the image in vignette style")
	mogrifyCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", nil, "access method for pixels outside the boundaries of the image")
	mogrifyCmd.Flags().StringSliceS("wave", "wave", nil, "alter an image along a sine wave")
	mogrifyCmd.Flags().StringSliceS("wavelet-denoise", "wavelet-denoise", nil, "removes noise from the image using a wavelet transform")
	mogrifyCmd.Flags().StringSliceS("weight", "weight", nil, "render text with this font weight")
	mogrifyCmd.Flags().StringSliceS("white-point", "white-point", nil, "chromaticity white point")
	mogrifyCmd.Flags().StringSliceS("white-threshold", "white-threshold", nil, "force all pixels above the threshold into white")
	mogrifyCmd.Flags().StringSliceS("word-break", "word-break", nil, "sets whether line breaks appear wherever the text would otherwise overflow its content box. Choose from normal, the default, or break-word.")
	mogrifyCmd.Flags().StringSliceS("write", "write", nil, "write images to this file")
	mogrifyCmd.Flags().StringSliceS("write-mask", "write-mask", nil, "associate a write mask with the image")
	rootCmd.AddCommand(mogrifyCmd)

	carapace.Gen(mogrifyCmd).FlagCompletion(carapace.ActionMap{
		"auto-threshold":    carapace.ActionValues("Undefined", "Kapur", "OTSU", "Triangle"),
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
		"mattecolor":        action.ActionColors(),
		"metric":            action.ActionCompleteOption("Metric"),
		"morphology":        action.ActionCompleteOption("Morphology"),
		"opaque":            action.ActionColors(),
		"orient":            action.ActionCompleteOption("Orientation"),
		"path":              carapace.ActionDirectories(),
		"preview":           action.ActionCompleteOption("Preview"),
		"process":           action.ActionCompleteOption("Filter"),
		"profile":           carapace.ActionFiles(),
		"quantize":          action.ActionCompleteOption("Colorspace"),
		"read-mask":         carapace.ActionFiles(),
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
		"write-mask":        carapace.ActionFiles(),
	})

	carapace.Gen(mogrifyCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
