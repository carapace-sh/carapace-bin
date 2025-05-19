package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var animateCmd = &cobra.Command{
	Use:   "animate",
	Short: "animate a sequence of images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(animateCmd).Standalone()

	animateCmd.Flags().StringSliceS("authenticate", "authenticate", nil, "decrypt image with this password")
	animateCmd.Flags().CountS("backdrop", "backdrop", "display image centered on a backdrop")
	animateCmd.Flags().StringSliceS("background", "background", nil, "the background color")
	animateCmd.Flags().StringSliceS("bordercolor", "bordercolor", nil, "the border color")
	animateCmd.Flags().StringSliceS("borderwidth", "borderwidth", nil, "the border width")
	animateCmd.Flags().StringSliceS("colormap", "colormap", nil, "Shared or Private")
	animateCmd.Flags().CountS("colors", "colors", "preferred number of colors in the image")
	animateCmd.Flags().StringSliceS("colorspace", "colorspace", nil, "alternate image colorspace")
	animateCmd.Flags().StringSliceS("debug", "debug", nil, "display copious debugging information")
	animateCmd.Flags().StringSliceS("define", "define", nil, "Coder/decoder specific options")
	animateCmd.Flags().StringSliceS("delay", "delay", nil, "display the next image after pausing")
	animateCmd.Flags().StringSliceS("density", "density", nil, "horizontal and vertical density of the image")
	animateCmd.Flags().StringSliceS("depth", "depth", nil, "image depth")
	animateCmd.Flags().StringSliceS("display", "display", nil, "display image to this X server")
	animateCmd.Flags().CountS("dither", "dither", "apply Floyd/Steinberg error diffusion to image")
	animateCmd.Flags().StringSliceS("font", "font", nil, "use this font when annotating the image with text")
	animateCmd.Flags().StringSliceS("foreground", "foreground", nil, "define the foreground color")
	animateCmd.Flags().StringSliceS("gamma", "gamma", nil, "level of gamma correction")
	animateCmd.Flags().StringSliceS("geometry", "geometry", nil, "preferred size and location of the cropped image")
	animateCmd.Flags().BoolS("help", "help", false, "print program options")
	animateCmd.Flags().StringSliceS("iconGeometry", "iconGeometry", nil, "specify the icon geometry")
	animateCmd.Flags().StringSliceS("iconic", "iconic", nil, "iconic animation")
	animateCmd.Flags().StringSliceS("interlace", "interlace", nil, "None, Line, Plane, or Partition")
	animateCmd.Flags().StringSliceS("limit", "limit", nil, "Disk, File, Map, Memory, Pixels, Width, Height, Threads, Read, or Write resource limit")
	animateCmd.Flags().StringSliceS("log", "log", nil, "format of debugging information")
	animateCmd.Flags().StringSliceS("map", "map", nil, "display image using this Standard Colormap")
	animateCmd.Flags().CountS("matte", "matte", "store matte channel if the image has one")
	animateCmd.Flags().StringSliceS("mattecolor", "mattecolor", nil, "specify the color to be used with the -frame option")
	animateCmd.Flags().CountS("monitor", "monitor", "show progress indication")
	animateCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	animateCmd.Flags().StringSliceS("name", "name", nil, "name an image")
	animateCmd.Flags().CountS("noop", "noop", "do not apply options to image")
	animateCmd.Flags().CountS("pause", "pause", "seconds to pause before reanimating")
	animateCmd.Flags().StringSliceS("remote", "remote", nil, "execute a command in a remote display process")
	animateCmd.Flags().StringSliceS("rotate", "rotate", nil, "apply Paeth rotation to the image")
	animateCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", nil, "horizontal and vertical sampling factors")
	animateCmd.Flags().StringSliceS("scenes", "scenes", nil, "image scene range")
	animateCmd.Flags().CountS("shared-memory", "shared-memory", "use shared memory")
	animateCmd.Flags().StringSliceS("size", "size", nil, "width and height of image")
	animateCmd.Flags().StringSliceS("title", "title", nil, "assign title to displayed image")
	animateCmd.Flags().StringSliceS("treedepth", "treedepth", nil, "color tree depth")
	animateCmd.Flags().CountS("trim", "trim", "trim image edges")
	animateCmd.Flags().StringSliceS("type", "type", nil, "image type")
	animateCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	animateCmd.Flags().BoolS("version", "version", false, "print version information")
	animateCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", nil, "Constant, Edge, Mirror, or Tile")
	animateCmd.Flags().StringSliceS("visual", "visual", nil, "display image using this visual type")
	animateCmd.Flags().StringSliceS("window", "window", nil, "display image to background of this window")
	rootCmd.AddCommand(animateCmd)

	carapace.Gen(animateCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
