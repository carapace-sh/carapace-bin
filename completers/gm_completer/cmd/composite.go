package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var compositeCmd = &cobra.Command{
	Use:   "composite",
	Short: "composite images together",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compositeCmd).Standalone()

	compositeCmd.Flags().StringSliceS("affine", "affine", nil, "affine transform matrix")
	compositeCmd.Flags().StringSliceS("authenticate", "authenticate", nil, "decrypt image with this password")
	compositeCmd.Flags().StringSliceS("blue-primary", "blue-primary", nil, "chomaticity blue primary point")
	compositeCmd.Flags().CountS("colors", "colors", "preferred number of colors in the image")
	compositeCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "alternate image colorspace")
	compositeCmd.Flags().StringSliceS("comment", "comment", nil, "annotate image with comment")
	compositeCmd.Flags().StringSliceS("compose", "compose", nil, "composite operator")
	compositeCmd.Flags().StringSliceS("compress", "compress", nil, "image compression type")
	compositeCmd.Flags().StringSliceS("debug", "debug", nil, "display copious debugging information")
	compositeCmd.Flags().StringSliceS("define", "define", nil, "Coder/decoder specific options")
	compositeCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	compositeCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	compositeCmd.Flags().StringSliceS("displace", "displace", nil, "shift image pixels defined by a displacement map")
	compositeCmd.Flags().StringSliceS("display", "display", nil, "get image or font from this X server")
	compositeCmd.Flags().StringSliceS("dispose", "dispose", nil, "Undefined, None, Background, Previous")
	compositeCmd.Flags().StringSliceS("dissolve", "dissolve", nil, "dissolve the two images a given percent")
	compositeCmd.Flags().CountS("dither", "dither", "apply Floyd/Steinberg error diffusion to image")
	compositeCmd.Flags().StringSliceS("encoding", "encoding", nil, "text encoding type")
	compositeCmd.Flags().StringSliceS("endian", "endian", nil, "multibyte word order (LSB, MSB, or Native)")
	compositeCmd.Flags().StringSliceS("filter", "filter", nil, "use this filter when resizing an image")
	compositeCmd.Flags().CountS("flag", "flag", "affine transform image")
	compositeCmd.Flags().StringSliceS("font", "font", nil, "render text with this font")
	compositeCmd.Flags().StringSliceS("geometry", "geometry", nil, "location of the composite image")
	compositeCmd.Flags().StringSliceS("gravity", "gravity", nil, "which direction to gravitate towards")
	compositeCmd.Flags().StringSliceS("green-primary", "green-primary", nil, "chomaticity green primary point")
	compositeCmd.Flags().BoolS("help", "help", false, "print program options")
	compositeCmd.Flags().StringSliceS("interlace", "interlace", nil, "None, Line, Plane, or Partition")
	compositeCmd.Flags().StringSliceS("label", "label", nil, "ssign a label to an image")
	compositeCmd.Flags().StringSliceS("limit", "limit", nil, "Disk, File, Map, Memory, Pixels, Width, Height, Threads, Read, or Write resource limit")
	compositeCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	compositeCmd.Flags().CountS("matte", "matte", "store matte channel if the image has one")
	compositeCmd.Flags().CountS("monitor", "monitor", "show progress indication")
	compositeCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	compositeCmd.Flags().CountS("negate", "negate", "replace every pixel with its complementary color ")
	compositeCmd.Flags().StringSliceS("page", "page", nil, "size and location of an image canvas")
	compositeCmd.Flags().StringSliceS("profile", "profile", nil, "add ICM or IPTC information profile to image")
	compositeCmd.Flags().StringSliceS("quality", "quality", nil, "JPEG/MIFF/PNG compression level")
	compositeCmd.Flags().StringSliceS("recolor", "recolor", nil, "apply a color translation matrix to image channels")
	compositeCmd.Flags().StringSliceS("red-primary", "red-primary", nil, "chomaticity red primary point")
	compositeCmd.Flags().StringSliceS("repage", "repage", nil, "adjust current page offsets by geometry")
	compositeCmd.Flags().StringSliceS("resize", "resize", nil, "resize the image")
	compositeCmd.Flags().StringSliceS("rotate", "rotate", nil, "apply Paeth rotation to the image")
	compositeCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factors")
	compositeCmd.Flags().StringSliceS("scene", "scene", nil, "image scene number")
	compositeCmd.Flags().StringSliceS("set", "set", nil, "set image attribute")
	compositeCmd.Flags().StringSliceS("sharpen", "sharpen", nil, "sharpen the image")
	compositeCmd.Flags().StringSliceS("size", "size", nil, "width and height of image")
	compositeCmd.Flags().StringSliceS("stegano", "stegano", nil, "hide watermark within an image")
	compositeCmd.Flags().CountS("stereo", "stereo", "combine two image to create a stereo anaglyph")
	compositeCmd.Flags().CountS("strip", "strip", "strip all profiles and text attributes from image")
	compositeCmd.Flags().StringSliceS("thumbnail", "thumbnail", nil, "resize the image (optimized for thumbnails)")
	compositeCmd.Flags().CountS("tile", "tile", "repeat composite operation across image")
	compositeCmd.Flags().StringSliceS("treedepth", "treedepth", nil, "color tree depth")
	compositeCmd.Flags().StringSliceS("type", "type", nil, "image type")
	compositeCmd.Flags().StringSliceS("units", "units", nil, "PixelsPerInch, PixelsPerCentimeter, or Undefined")
	compositeCmd.Flags().StringSliceS("unsharp", "unsharp", nil, "sharpen the image")
	compositeCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	compositeCmd.Flags().BoolS("version", "version", false, "print version information")
	compositeCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", nil, "Constant, Edge, Mirror, or Tile")
	compositeCmd.Flags().StringSliceS("watermark", "watermark", nil, "percent brightness and saturation of a watermark")
	compositeCmd.Flags().StringSliceS("white-point", "white-point", nil, "chomaticity white point")
	compositeCmd.Flags().StringSliceS("write", "write", nil, "write image to this file")
	rootCmd.AddCommand(compositeCmd)

	carapace.Gen(compositeCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
