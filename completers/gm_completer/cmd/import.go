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
	importCmd.Flags().StringSliceS("colorspace", "colorspace", []string{}, "alternate image colorspace")
	importCmd.Flags().StringSliceS("comment", "comment", []string{}, "annotate image with comment")
	importCmd.Flags().StringSliceS("compress", "compress", []string{}, "image compression type")
	importCmd.Flags().StringSliceS("crop", "crop", []string{}, "preferred size and location of the cropped image")
	importCmd.Flags().StringSliceS("debug", "debug", []string{}, "display copious debugging information")
	importCmd.Flags().StringSliceS("define", "define", []string{}, "Coder/decoder specific options")
	importCmd.Flags().StringSliceS("delay", "delay", []string{}, "display the next image after pausing")
	importCmd.Flags().StringSliceS("density", "density", []string{}, "horizontal and vertical density of the image")
	importCmd.Flags().StringSliceS("depth", "depth", []string{}, "image depth")
	importCmd.Flags().CountS("descend", "descend", "obtain image by descending window hierarchy")
	importCmd.Flags().StringSliceS("display", "display", []string{}, "X server to contact")
	importCmd.Flags().StringSliceS("dispose", "dispose", []string{}, "Undefined, None, Background, Previous")
	importCmd.Flags().CountS("dither", "dither", "apply Floyd/Steinberg error diffusion to image")
	importCmd.Flags().StringSliceS("encoding", "encoding", []string{}, "text encoding type")
	importCmd.Flags().StringSliceS("endian", "endian", []string{}, "multibyte word order (LSB, MSB, or Native)")
	importCmd.Flags().CountS("frame", "frame", "include window manager frame")
	importCmd.Flags().StringSliceS("geometry", "geometry", []string{}, "preferred size or location of the image")
	importCmd.Flags().BoolS("help", "help", false, "program options")
	importCmd.Flags().StringSliceS("interlace", "interlace", []string{}, "None, Line, Plane, or Partition")
	importCmd.Flags().StringSliceS("label", "label", []string{}, "assign a label to an image")
	importCmd.Flags().StringSliceS("limit", "limit", []string{}, "Disk, File, Map, Memory, Pixels, Width, Height or Threads resource limit")
	importCmd.Flags().StringSliceS("log", "log", []string{}, "format of debugging information")
	importCmd.Flags().CountS("monitor", "monitor", "show progress indication")
	importCmd.Flags().StringSliceS("monochrome", "monochrome", []string{}, "image to black and white")
	importCmd.Flags().CountS("negate", "negate", "replace every pixel with its complementary color ")
	importCmd.Flags().StringSliceS("page", "page", []string{}, "size and location of an image canvas")
	importCmd.Flags().StringSliceS("pause", "pause", []string{}, "seconds delay between snapshots")
	importCmd.Flags().StringSliceS("pointsize", "pointsize", []string{}, "font point size")
	importCmd.Flags().StringSliceS("quality", "quality", []string{}, "JPEG/MIFF/PNG compression level")
	importCmd.Flags().StringSliceS("resize", "resize", []string{}, "resize the image")
	importCmd.Flags().StringSliceS("rotate", "rotate", []string{}, "apply Paeth rotation to the image")
	importCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", []string{}, "horizontal and vertical sampling factors")
	importCmd.Flags().StringSliceS("scene", "scene", []string{}, "image scene number")
	importCmd.Flags().CountS("screen", "screen", "select image from root window")
	importCmd.Flags().StringSliceS("set", "set", []string{}, "set image attribute")
	importCmd.Flags().CountS("silent", "silent", "operate silently, i.e. don't ring any bells ")
	importCmd.Flags().StringSliceS("snaps", "snaps", []string{}, "number of screen snapshots")
	importCmd.Flags().StringSliceS("thumbnail", "thumbnail", []string{}, "resize the image (optimized for thumbnails)")
	importCmd.Flags().StringSliceS("transparent", "transparent", []string{}, "make this color transparent within the image")
	importCmd.Flags().StringSliceS("treedepth", "treedepth", []string{}, "color tree depth")
	importCmd.Flags().CountS("trim", "trim", "trim image edges")
	importCmd.Flags().StringSliceS("type", "type", []string{}, "image type")
	importCmd.Flags().BoolS("verbose", "verbose", false, "detailed information about the image")
	importCmd.Flags().BoolS("version", "version", false, "version information")
	importCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", []string{}, "Constant, Edge, Mirror, or Tile")
	importCmd.Flags().StringSliceS("window", "window", []string{}, "select window with this id or name")
	rootCmd.AddCommand(importCmd)

	carapace.Gen(importCmd).PositionalCompletion(carapace.ActionFiles())
}
