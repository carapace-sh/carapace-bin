package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "capture an application or X server screen",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().CountS("adjoin", "adjoin", "join images into a single multi-image file")
	importCmd.Flags().CountS("border", "border", "include image borders in the output image")
	importCmd.Flags().CountS("colors", "colors", "preferred number of colors in the image")
	importCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "alternate image colorspace")
	importCmd.Flags().StringSliceS("comment", "comment", nil, "annotate image with comment")
	importCmd.Flags().StringSliceS("compress", "compress", nil, "image compression type")
	importCmd.Flags().StringSliceS("crop", "crop", nil, "preferred size and location of the cropped image")
	importCmd.Flags().StringSliceS("debug", "debug", nil, "display copious debugging information")
	importCmd.Flags().StringSliceS("define", "define", nil, "Coder/decoder specific options")
	importCmd.Flags().StringSliceS("delay", "delay", nil, "display the next image after pausing")
	importCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	importCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	importCmd.Flags().CountS("descend", "descend", "obtain image by descending window hierarchy")
	importCmd.Flags().StringSliceS("display", "display", nil, "X server to contact")
	importCmd.Flags().StringSliceS("dispose", "dispose", nil, "Undefined, None, Background, Previous")
	importCmd.Flags().CountS("dither", "dither", "apply Floyd/Steinberg error diffusion to image")
	importCmd.Flags().StringSliceS("encoding", "encoding", nil, "text encoding type")
	importCmd.Flags().StringSliceS("endian", "endian", nil, "multibyte word order (LSB, MSB, or Native)")
	importCmd.Flags().CountS("frame", "frame", "include window manager frame")
	importCmd.Flags().StringSliceS("geometry", "geometry", nil, "preferred size or location of the image")
	importCmd.Flags().BoolS("help", "help", false, "program options")
	importCmd.Flags().StringSliceS("interlace", "interlace", nil, "None, Line, Plane, or Partition")
	importCmd.Flags().StringSliceS("label", "label", nil, "assign a label to an image")
	importCmd.Flags().StringSliceS("limit", "limit", nil, "Disk, File, Map, Memory, Pixels, Width, Height or Threads resource limit")
	importCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	importCmd.Flags().CountS("monitor", "monitor", "show progress indication")
	importCmd.Flags().StringSliceS("monochrome", "monochrome", nil, "image to black and white")
	importCmd.Flags().CountS("negate", "negate", "replace every pixel with its complementary color ")
	importCmd.Flags().StringSliceS("page", "page", nil, "size and location of an image canvas")
	importCmd.Flags().StringSliceS("pause", "pause", nil, "seconds delay between snapshots")
	importCmd.Flags().StringSliceS("pointsize", "pointsize", nil, "font point size")
	importCmd.Flags().StringSliceS("quality", "quality", nil, "JPEG/MIFF/PNG compression level")
	importCmd.Flags().StringSliceS("resize", "resize", nil, "resize the image")
	importCmd.Flags().StringSliceS("rotate", "rotate", nil, "apply Paeth rotation to the image")
	importCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factors")
	importCmd.Flags().StringSliceS("scene", "scene", nil, "image scene number")
	importCmd.Flags().CountS("screen", "screen", "select image from root window")
	importCmd.Flags().StringSliceS("set", "set", nil, "set image attribute")
	importCmd.Flags().CountS("silent", "silent", "operate silently, i.e. don't ring any bells ")
	importCmd.Flags().StringSliceS("snaps", "snaps", nil, "number of screen snapshots")
	importCmd.Flags().StringSliceS("thumbnail", "thumbnail", nil, "resize the image (optimized for thumbnails)")
	importCmd.Flags().StringSliceS("transparent", "transparent", nil, "make this color transparent within the image")
	importCmd.Flags().StringSliceS("treedepth", "treedepth", nil, "color tree depth")
	importCmd.Flags().CountS("trim", "trim", "trim image edges")
	importCmd.Flags().StringSliceS("type", "type", nil, "image type")
	importCmd.Flags().BoolS("verbose", "verbose", false, "detailed information about the image")
	importCmd.Flags().BoolS("version", "version", false, "version information")
	importCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", nil, "Constant, Edge, Mirror, or Tile")
	importCmd.Flags().StringSliceS("window", "window", nil, "select window with this id or name")
	rootCmd.AddCommand(importCmd)

	carapace.Gen(importCmd).PositionalCompletion(carapace.ActionFiles())
}
