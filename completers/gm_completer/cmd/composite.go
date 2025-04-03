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

	compositeCmd.Flags().StringSliceS("affine", "affine", []string{}, "affine transform matrix")
	compositeCmd.Flags().StringSliceS("authenticate", "authenticate", []string{}, "decrypt image with this password")
	compositeCmd.Flags().StringSliceS("blue-primary", "blue-primary", []string{}, "chomaticity blue primary point")
	compositeCmd.Flags().CountS("colors", "colors", "preferred number of colors in the image")
	compositeCmd.Flags().StringSliceS("colorspace", "colorspace", []string{}, "alternate image colorspace")
	compositeCmd.Flags().StringSliceS("comment", "comment", []string{}, "annotate image with comment")
	compositeCmd.Flags().StringSliceS("compose", "compose", []string{}, "composite operator")
	compositeCmd.Flags().StringSliceS("compress", "compress", []string{}, "image compression type")
	compositeCmd.Flags().StringSliceS("debug", "debug", []string{}, "display copious debugging information")
	compositeCmd.Flags().StringSliceS("define", "define", []string{}, "Coder/decoder specific options")
	compositeCmd.Flags().StringSliceS("density", "density", []string{}, "horizontal and vertical density of the image")
	compositeCmd.Flags().StringSliceS("depth", "depth", []string{}, "image depth")
	compositeCmd.Flags().StringSliceS("displace", "displace", []string{}, "shift image pixels defined by a displacement map")
	compositeCmd.Flags().StringSliceS("display", "display", []string{}, "get image or font from this X server")
	compositeCmd.Flags().StringSliceS("dispose", "dispose", []string{}, "Undefined, None, Background, Previous")
	compositeCmd.Flags().StringSliceS("dissolve", "dissolve", []string{}, "dissolve the two images a given percent")
	compositeCmd.Flags().CountS("dither", "dither", "apply Floyd/Steinberg error diffusion to image")
	compositeCmd.Flags().StringSliceS("encoding", "encoding", []string{}, "text encoding type")
	compositeCmd.Flags().StringSliceS("endian", "endian", []string{}, "multibyte word order (LSB, MSB, or Native)")
	compositeCmd.Flags().StringSliceS("filter", "filter", []string{}, "use this filter when resizing an image")
	compositeCmd.Flags().CountS("flag", "flag", "affine transform image")
	compositeCmd.Flags().StringSliceS("font", "font", []string{}, "render text with this font")
	compositeCmd.Flags().StringSliceS("geometry", "geometry", []string{}, "location of the composite image")
	compositeCmd.Flags().StringSliceS("gravity", "gravity", []string{}, "which direction to gravitate towards")
	compositeCmd.Flags().StringSliceS("green-primary", "green-primary", []string{}, "chomaticity green primary point")
	compositeCmd.Flags().BoolS("help", "help", false, "print program options")
	compositeCmd.Flags().StringSliceS("interlace", "interlace", []string{}, "None, Line, Plane, or Partition")
	compositeCmd.Flags().StringSliceS("label", "label", []string{}, "ssign a label to an image")
	compositeCmd.Flags().StringSliceS("limit", "limit", []string{}, "Disk, File, Map, Memory, Pixels, Width, Height, Threads, Read, or Write resource limit")
	compositeCmd.Flags().StringSliceS("log", "log", []string{}, "format of debugging information")
	compositeCmd.Flags().CountS("matte", "matte", "store matte channel if the image has one")
	compositeCmd.Flags().CountS("monitor", "monitor", "show progress indication")
	compositeCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	compositeCmd.Flags().CountS("negate", "negate", "replace every pixel with its complementary color ")
	compositeCmd.Flags().StringSliceS("page", "page", []string{}, "size and location of an image canvas")
	compositeCmd.Flags().StringSliceS("profile", "profile", []string{}, "add ICM or IPTC information profile to image")
	compositeCmd.Flags().StringSliceS("quality", "quality", []string{}, "JPEG/MIFF/PNG compression level")
	compositeCmd.Flags().StringSliceS("recolor", "recolor", []string{}, "apply a color translation matrix to image channels")
	compositeCmd.Flags().StringSliceS("red-primary", "red-primary", []string{}, "chomaticity red primary point")
	compositeCmd.Flags().StringSliceS("repage", "repage", []string{}, "adjust current page offsets by geometry")
	compositeCmd.Flags().StringSliceS("resize", "resize", []string{}, "resize the image")
	compositeCmd.Flags().StringSliceS("rotate", "rotate", []string{}, "apply Paeth rotation to the image")
	compositeCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", []string{}, "horizontal and vertical sampling factors")
	compositeCmd.Flags().StringSliceS("scene", "scene", []string{}, "image scene number")
	compositeCmd.Flags().StringSliceS("set", "set", []string{}, "set image attribute")
	compositeCmd.Flags().StringSliceS("sharpen", "sharpen", []string{}, "sharpen the image")
	compositeCmd.Flags().StringSliceS("size", "size", []string{}, "width and height of image")
	compositeCmd.Flags().StringSliceS("stegano", "stegano", []string{}, "hide watermark within an image")
	compositeCmd.Flags().CountS("stereo", "stereo", "combine two image to create a stereo anaglyph")
	compositeCmd.Flags().CountS("strip", "strip", "strip all profiles and text attributes from image")
	compositeCmd.Flags().StringSliceS("thumbnail", "thumbnail", []string{}, "resize the image (optimized for thumbnails)")
	compositeCmd.Flags().CountS("tile", "tile", "repeat composite operation across image")
	compositeCmd.Flags().StringSliceS("treedepth", "treedepth", []string{}, "color tree depth")
	compositeCmd.Flags().StringSliceS("type", "type", []string{}, "image type")
	compositeCmd.Flags().StringSliceS("units", "units", []string{}, "PixelsPerInch, PixelsPerCentimeter, or Undefined")
	compositeCmd.Flags().StringSliceS("unsharp", "unsharp", []string{}, "sharpen the image")
	compositeCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	compositeCmd.Flags().BoolS("version", "version", false, "print version information")
	compositeCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", []string{}, "Constant, Edge, Mirror, or Tile")
	compositeCmd.Flags().StringSliceS("watermark", "watermark", []string{}, "percent brightness and saturation of a watermark")
	compositeCmd.Flags().StringSliceS("white-point", "white-point", []string{}, "chomaticity white point")
	compositeCmd.Flags().StringSliceS("write", "write", []string{}, "write image to this file")
	rootCmd.AddCommand(compositeCmd)

	carapace.Gen(compositeCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
