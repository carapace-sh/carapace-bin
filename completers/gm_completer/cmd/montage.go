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
	montageCmd.Flags().StringSliceS("affine", "affine", nil, "affine transform matrix")
	montageCmd.Flags().StringSliceS("authenticate", "authenticate", nil, "decrypt image with this password")
	montageCmd.Flags().StringSliceS("background", "background", nil, "background color")
	montageCmd.Flags().StringSliceS("blue-primary", "blue-primary", nil, "chomaticity blue primary point")
	montageCmd.Flags().StringSliceS("blur", "blur", nil, "apply a filter to blur the image")
	montageCmd.Flags().StringSliceS("bordercolor", "bordercolor", nil, "border color")
	montageCmd.Flags().StringSliceS("borderwidth", "borderwidth", nil, "border width")
	montageCmd.Flags().CountS("colors", "colors", "preferred number of colors in the image")
	montageCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "alternate image colorsapce")
	montageCmd.Flags().StringSliceS("comment", "comment", nil, "annotate image with comment")
	montageCmd.Flags().StringSliceS("compose", "compose", nil, "composite operator")
	montageCmd.Flags().StringSliceS("compress", "compress", nil, "image compression type")
	montageCmd.Flags().StringSliceS("crop", "crop", nil, "preferred size and location of the cropped image")
	montageCmd.Flags().StringSliceS("debug", "debug", nil, "display copious debugging information")
	montageCmd.Flags().StringSliceS("define", "define", nil, "Coder/decoder specific options")
	montageCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	montageCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	montageCmd.Flags().StringSliceS("display", "display", nil, "query font from this X server")
	montageCmd.Flags().StringSliceS("dispose", "dispose", nil, "Undefined, None, Background, Previous")
	montageCmd.Flags().CountS("dither", "dither", "apply Floyd/Steinberg error diffusion to image")
	montageCmd.Flags().StringSliceS("draw", "draw", nil, "annotate the image with a graphic primitive")
	montageCmd.Flags().StringSliceS("encoding", "encoding", nil, "text encoding type")
	montageCmd.Flags().StringSliceS("endian", "endian", nil, "multibyte word order (LSB, MSB, or Native)")
	montageCmd.Flags().StringSliceS("fill", "fill", nil, "color to use when filling a graphic primitive")
	montageCmd.Flags().StringSliceS("filter", "filter", nil, "use this filter when resizing an image")
	montageCmd.Flags().CountS("flip", "flip", "flip image in the vertical direction")
	montageCmd.Flags().CountS("flop", "flop", "flop image in the horizontal direction")
	montageCmd.Flags().StringSliceS("font", "font", nil, "font to use when annotating with text")
	montageCmd.Flags().StringSliceS("format", "format", nil, "output formatted image characteristics")
	montageCmd.Flags().StringSliceS("frame", "frame", nil, "surround image with an ornamental border")
	montageCmd.Flags().StringSliceS("gamma", "gamma", nil, "level of gamma correction")
	montageCmd.Flags().StringSliceS("geometry", "geometry", nil, "preferred tile and border sizes")
	montageCmd.Flags().StringSliceS("gravity", "gravity", nil, "which direction to gravitate towards")
	montageCmd.Flags().StringSliceS("green-primary", "green-primary", nil, "chomaticity green primary point")
	montageCmd.Flags().CountS("help", "help", "print program options")
	montageCmd.Flags().StringSliceS("interlace", "interlace", nil, "None, Line, Plane, or Partition")
	montageCmd.Flags().StringSliceS("label", "label", nil, "assign a label to an image")
	montageCmd.Flags().StringSliceS("limit", "limit", nil, "Disk, File, Map, Memory, Pixels, Width, Height or Threads resource limit")
	montageCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	montageCmd.Flags().CountS("matte", "matte", "store matte channel if the image has one")
	montageCmd.Flags().StringSliceS("mattecolor", "mattecolor", nil, "color to be used with the -frame option")
	montageCmd.Flags().StringSliceS("mode", "mode", nil, "Frame, Unframe, or Concatenate")
	montageCmd.Flags().CountS("monitor", "monitor", "show progress indication")
	montageCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	montageCmd.Flags().CountS("noop", "noop", "do not apply options to image")
	montageCmd.Flags().StringSliceS("page", "page", nil, "size and location of an image canvas")
	montageCmd.Flags().StringSliceS("pointsize", "pointsize", nil, "font point size")
	montageCmd.Flags().StringSliceS("quality", "quality", nil, "JPEG/MIFF/PNG compression level")
	montageCmd.Flags().StringSliceS("red-primary", "red-primary", nil, "chomaticity red primary point")
	montageCmd.Flags().StringSliceS("repage", "repage", nil, "adjust current page offsets by geometry")
	montageCmd.Flags().StringSliceS("resize", "resize", nil, "resize the image")
	montageCmd.Flags().StringSliceS("rotate", "rotate", nil, "apply Paeth rotation to the image")
	montageCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factors")
	montageCmd.Flags().StringSliceS("scenes", "scenes", nil, "image scene range")
	montageCmd.Flags().StringSliceS("set", "set", nil, "set image attribute")
	montageCmd.Flags().CountS("shadow", "shadow", "add a shadow beneath a tile to simulate depth")
	montageCmd.Flags().StringSliceS("sharpen", "sharpen", nil, "sharpen the image")
	montageCmd.Flags().StringSliceS("size", "size", nil, "width and height of image")
	montageCmd.Flags().CountS("strip", "strip", "strip all profiles and text attributes from image")
	montageCmd.Flags().StringSliceS("stroke", "stroke", nil, "color to use when stroking a graphic primitive")
	montageCmd.Flags().StringSliceS("strokewidth", "strokewidth", nil, "stroke (line) width")
	montageCmd.Flags().StringSliceS("texture", "texture", nil, "name of texture to tile onto the image background")
	montageCmd.Flags().StringSliceS("thumbnail", "thumbnail", nil, "resize the image (optimized for thumbnails)")
	montageCmd.Flags().StringSliceS("tile", "tile", nil, "number of tiles per row and column")
	montageCmd.Flags().StringSliceS("title", "title", nil, "thumbnail title")
	montageCmd.Flags().CountS("transform", "transform", "affine transform image")
	montageCmd.Flags().StringSliceS("transparent", "transparent", nil, "make this color transparent within the image")
	montageCmd.Flags().StringSliceS("treedepth", "treedepth", nil, "color tree depth")
	montageCmd.Flags().CountS("trim", "trim", "trim image edges")
	montageCmd.Flags().StringSliceS("type", "type", nil, "image type")
	montageCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	montageCmd.Flags().BoolS("version", "version", false, "print version information")
	montageCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", nil, "Constant, Edge, Mirror, or Tile")
	montageCmd.Flags().StringSliceS("white-point", "white-point", nil, "chomaticity white point")
	rootCmd.AddCommand(montageCmd)

	carapace.Gen(montageCmd).FlagCompletion(carapace.ActionMap{
		"fill": action.ActionColor(),
	})

	carapace.Gen(montageCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
