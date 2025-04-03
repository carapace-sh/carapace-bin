package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var montageCmd = &cobra.Command{
	Use:   "montage",
	Short: "create a composite image (in a grid) from separate images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(montageCmd).Standalone()

	montageCmd.Flags().CountS("adjoin", "adjoin", "join images into a single multi-image file")
	montageCmd.Flags().StringSliceS("affine", "affine", []string{}, "affine transform matrix")
	montageCmd.Flags().StringSliceS("authenticate", "authenticate", []string{}, "decrypt image with this password")
	montageCmd.Flags().StringSliceS("background", "background", []string{}, "background color")
	montageCmd.Flags().StringSliceS("blue-primary", "blue-primary", []string{}, "chomaticity blue primary point")
	montageCmd.Flags().StringSliceS("blur", "blur", []string{}, "apply a filter to blur the image")
	montageCmd.Flags().StringSliceS("bordercolor", "bordercolor", []string{}, "border color")
	montageCmd.Flags().StringSliceS("borderwidth", "borderwidth", []string{}, "border width")
	montageCmd.Flags().CountS("colors", "colors", "preferred number of colors in the image")
	montageCmd.Flags().StringSliceS("colorspace", "colorspace", []string{}, "alternate image colorsapce")
	montageCmd.Flags().StringSliceS("comment", "comment", []string{}, "annotate image with comment")
	montageCmd.Flags().StringSliceS("compose", "compose", []string{}, "composite operator")
	montageCmd.Flags().StringSliceS("compress", "compress", []string{}, "image compression type")
	montageCmd.Flags().StringSliceS("crop", "crop", []string{}, "preferred size and location of the cropped image")
	montageCmd.Flags().StringSliceS("debug", "debug", []string{}, "display copious debugging information")
	montageCmd.Flags().StringSliceS("define", "define", []string{}, "Coder/decoder specific options")
	montageCmd.Flags().StringSliceS("density", "density", []string{}, "horizontal and vertical density of the image")
	montageCmd.Flags().StringSliceS("depth", "depth", []string{}, "image depth")
	montageCmd.Flags().StringSliceS("display", "display", []string{}, "query font from this X server")
	montageCmd.Flags().StringSliceS("dispose", "dispose", []string{}, "Undefined, None, Background, Previous")
	montageCmd.Flags().CountS("dither", "dither", "apply Floyd/Steinberg error diffusion to image")
	montageCmd.Flags().StringSliceS("draw", "draw", []string{}, "annotate the image with a graphic primitive")
	montageCmd.Flags().StringSliceS("encoding", "encoding", []string{}, "text encoding type")
	montageCmd.Flags().StringSliceS("endian", "endian", []string{}, "multibyte word order (LSB, MSB, or Native)")
	montageCmd.Flags().StringSliceS("fill", "fill", []string{}, "color to use when filling a graphic primitive")
	montageCmd.Flags().StringSliceS("filter", "filter", []string{}, "use this filter when resizing an image")
	montageCmd.Flags().CountS("flip", "flip", "flip image in the vertical direction")
	montageCmd.Flags().CountS("flop", "flop", "flop image in the horizontal direction")
	montageCmd.Flags().StringSliceS("font", "font", []string{}, "font to use when annotating with text")
	montageCmd.Flags().StringSliceS("format", "format", []string{}, "output formatted image characteristics")
	montageCmd.Flags().StringSliceS("frame", "frame", []string{}, "surround image with an ornamental border")
	montageCmd.Flags().StringSliceS("gamma", "gamma", []string{}, "level of gamma correction")
	montageCmd.Flags().StringSliceS("geometry", "geometry", []string{}, "preferred tile and border sizes")
	montageCmd.Flags().StringSliceS("gravity", "gravity", []string{}, "which direction to gravitate towards")
	montageCmd.Flags().StringSliceS("green-primary", "green-primary", []string{}, "chomaticity green primary point")
	montageCmd.Flags().CountS("help", "help", "print program options")
	montageCmd.Flags().StringSliceS("interlace", "interlace", []string{}, "None, Line, Plane, or Partition")
	montageCmd.Flags().StringSliceS("label", "label", []string{}, "assign a label to an image")
	montageCmd.Flags().StringSliceS("limit", "limit", []string{}, "Disk, File, Map, Memory, Pixels, Width, Height or Threads resource limit")
	montageCmd.Flags().StringSliceS("log", "log", []string{}, "format of debugging information")
	montageCmd.Flags().CountS("matte", "matte", "store matte channel if the image has one")
	montageCmd.Flags().StringSliceS("mattecolor", "mattecolor", []string{}, "color to be used with the -frame option")
	montageCmd.Flags().StringSliceS("mode", "mode", []string{}, "Frame, Unframe, or Concatenate")
	montageCmd.Flags().CountS("monitor", "monitor", "show progress indication")
	montageCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	montageCmd.Flags().CountS("noop", "noop", "do not apply options to image")
	montageCmd.Flags().StringSliceS("page", "page", []string{}, "size and location of an image canvas")
	montageCmd.Flags().StringSliceS("pointsize", "pointsize", []string{}, "font point size")
	montageCmd.Flags().StringSliceS("quality", "quality", []string{}, "JPEG/MIFF/PNG compression level")
	montageCmd.Flags().StringSliceS("red-primary", "red-primary", []string{}, "chomaticity red primary point")
	montageCmd.Flags().StringSliceS("repage", "repage", []string{}, "adjust current page offsets by geometry")
	montageCmd.Flags().StringSliceS("resize", "resize", []string{}, "resize the image")
	montageCmd.Flags().StringSliceS("rotate", "rotate", []string{}, "apply Paeth rotation to the image")
	montageCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", []string{}, "horizontal and vertical sampling factors")
	montageCmd.Flags().StringSliceS("scenes", "scenes", []string{}, "image scene range")
	montageCmd.Flags().StringSliceS("set", "set", []string{}, "set image attribute")
	montageCmd.Flags().CountS("shadow", "shadow", "add a shadow beneath a tile to simulate depth")
	montageCmd.Flags().StringSliceS("sharpen", "sharpen", []string{}, "sharpen the image")
	montageCmd.Flags().StringSliceS("size", "size", []string{}, "width and height of image")
	montageCmd.Flags().CountS("strip", "strip", "strip all profiles and text attributes from image")
	montageCmd.Flags().StringSliceS("stroke", "stroke", []string{}, "color to use when stroking a graphic primitive")
	montageCmd.Flags().StringSliceS("strokewidth", "strokewidth", []string{}, "stroke (line) width")
	montageCmd.Flags().StringSliceS("texture", "texture", []string{}, "name of texture to tile onto the image background")
	montageCmd.Flags().StringSliceS("thumbnail", "thumbnail", []string{}, "resize the image (optimized for thumbnails)")
	montageCmd.Flags().StringSliceS("tile", "tile", []string{}, "number of tiles per row and column")
	montageCmd.Flags().StringSliceS("title", "title", []string{}, "thumbnail title")
	montageCmd.Flags().CountS("transform", "transform", "affine transform image")
	montageCmd.Flags().StringSliceS("transparent", "transparent", []string{}, "make this color transparent within the image")
	montageCmd.Flags().StringSliceS("treedepth", "treedepth", []string{}, "color tree depth")
	montageCmd.Flags().CountS("trim", "trim", "trim image edges")
	montageCmd.Flags().StringSliceS("type", "type", []string{}, "image type")
	montageCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	montageCmd.Flags().BoolS("version", "version", false, "print version information")
	montageCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", []string{}, "Constant, Edge, Mirror, or Tile")
	montageCmd.Flags().StringSliceS("white-point", "white-point", []string{}, "chomaticity white point")
	rootCmd.AddCommand(montageCmd)

	carapace.Gen(montageCmd).FlagCompletion(carapace.ActionMap{
		"fill": action.ActionColor(),
	})

	carapace.Gen(montageCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
