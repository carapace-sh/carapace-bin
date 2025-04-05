package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "display an image on a workstation running X",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(displayCmd).Standalone()

	displayCmd.Flags().StringS("authenticate", "authenticate", "", "decrypt image with this password")
	displayCmd.Flags().CountS("backdrop", "backdrop", "display image centered on a backdrop")
	displayCmd.Flags().StringS("border", "border", "", "surround image with a border of color")
	displayCmd.Flags().StringS("colormap", "colormap", "", "Shared or Private")
	displayCmd.Flags().CountS("colors", "colors", "preferred number of colors in the image")
	displayCmd.Flags().StringS("colorspace", "colorspace", "", "alternate image colorspace")
	displayCmd.Flags().StringS("comment", "comment", "", "annotate image with comment")
	displayCmd.Flags().StringS("compress", "compress", "", "image compression type")
	displayCmd.Flags().CountS("contrast", "contrast", "enhance or reduce the image contrast")
	displayCmd.Flags().StringS("crop", "crop", "", "preferred size and location of the cropped image")
	displayCmd.Flags().StringS("debug", "debug", "", "display copious debugging information")
	displayCmd.Flags().StringS("define", "define", "", "Coder/decoder specific options")
	displayCmd.Flags().StringS("delay", "delay", "", "display the next image after pausing")
	displayCmd.Flags().StringS("density", "density", "", "horizontal and vertical density of the image")
	displayCmd.Flags().StringS("depth", "depth", "", "image depth")
	displayCmd.Flags().CountS("despeckle", "despeckle", "reduce the speckles within an image")
	displayCmd.Flags().StringS("display", "display", "", "display image to this X server")
	displayCmd.Flags().StringS("dispose", "dispose", "", "Undefined, None, Background, Previous")
	displayCmd.Flags().CountS("dither", "dither", "apply Floyd/Steinberg error diffusion to image")
	displayCmd.Flags().StringS("edge", "edge", "", "apply a filter to detect edges in the image")
	displayCmd.Flags().StringS("endian", "endian", "", "multibyte word order (LSB, MSB, or Native)")
	displayCmd.Flags().CountS("enhance", "enhance", "apply a digital filter to enhance a noisy image")
	displayCmd.Flags().StringS("filter", "filter", "", "use this filter when resizing an image")
	displayCmd.Flags().CountS("flip", "flip", "flip image in the vertical direction")
	displayCmd.Flags().CountS("flop", "flop", "flop image in the horizontal direction")
	displayCmd.Flags().StringS("frame", "frame", "", "surround image with an ornamental border")
	displayCmd.Flags().StringS("gamma", "gamma", "", "level of gamma correction")
	displayCmd.Flags().StringS("geometry", "geometry", "", "preferred size and location of the Image window")
	displayCmd.Flags().BoolS("help", "help", false, "print program options")
	displayCmd.Flags().CountS("immutable", "immutable", "displayed image cannot be modified")
	displayCmd.Flags().StringS("interlace", "interlace", "", "None, Line, Plane, or Partition")
	displayCmd.Flags().StringS("label", "label", "", "assign a label to an image")
	displayCmd.Flags().StringS("limit", "limit", "", "Disk, File, Map, Memory, Pixels, Width, Height or Threads resource limit")
	displayCmd.Flags().StringS("log", "log", "", "format of debugging information")
	displayCmd.Flags().StringS("map", "map", "", "display image using this Standard Colormap")
	displayCmd.Flags().CountS("matte", "matte", "store matte channel if the image has one")
	displayCmd.Flags().CountS("monitor", "monitor", "show progress indication")
	displayCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	displayCmd.Flags().CountS("negate", "negate", "replace every pixel with its complementary color")
	displayCmd.Flags().CountS("noop", "noop", "do not apply options to image")
	displayCmd.Flags().StringS("page", "page", "", "size and location of an image canvas")
	displayCmd.Flags().StringS("quality", "quality", "", "JPEG/MIFF/PNG compression level")
	displayCmd.Flags().StringS("raise", "raise", "", "lighten/darken image edges to create a 3-D effect")
	displayCmd.Flags().StringS("remote", "remote", "", "execute a command in an remote display process")
	displayCmd.Flags().StringS("roll", "roll", "", "roll an image vertically or horizontally")
	displayCmd.Flags().StringS("rotate", "rotate", "", "apply Paeth rotation to the image")
	displayCmd.Flags().StringS("sample", "sample", "", "scale image with pixel sampling")
	displayCmd.Flags().StringS("sampling-factor", "sampling-factor", "", "horizontal and vertical sampling factors")
	displayCmd.Flags().StringS("scenes", "scenes", "", "image scene range")
	displayCmd.Flags().StringS("segment", "segment", "", "segment an image")
	displayCmd.Flags().StringS("set", "set", "", "set image attribute")
	displayCmd.Flags().StringS("sharpen", "sharpen", "", "sharpen the image")
	displayCmd.Flags().StringS("size", "size", "", "width and height of image")
	displayCmd.Flags().StringS("texture", "texture", "", "name of texture to tile onto the image background")
	displayCmd.Flags().StringS("treedepth", "treedepth", "", "color tree depth")
	displayCmd.Flags().CountS("trim", "trim", "trim image edges")
	displayCmd.Flags().StringS("type", "type", "", "image type")
	displayCmd.Flags().StringS("update", "update", "", "detect when image file is modified and redisplay")
	displayCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	displayCmd.Flags().BoolS("version", "version", false, "print version information")
	displayCmd.Flags().StringS("virtual-pixel", "virtual-pixel", "", "Constant, Edge, Mirror, or Tile")
	displayCmd.Flags().StringS("window", "window", "", "display image to background of this window")
	displayCmd.Flags().StringS("window_group", "window_group", "", "exit program when this window id is destroyed")
	displayCmd.Flags().StringS("write", "write", "", "write image to a file")
	rootCmd.AddCommand(displayCmd)

	carapace.Gen(displayCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
