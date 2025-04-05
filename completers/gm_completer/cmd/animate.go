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

	animateCmd.Flags().StringSliceS("authenticate", "authenticate", []string{}, "decrypt image with this password")
	animateCmd.Flags().CountS("backdrop", "backdrop", "display image centered on a backdrop")
	animateCmd.Flags().StringSliceS("background", "background", []string{}, "the background color")
	animateCmd.Flags().StringSliceS("bordercolor", "bordercolor", []string{}, "the border color")
	animateCmd.Flags().StringSliceS("borderwidth", "borderwidth", []string{}, "the border width")
	animateCmd.Flags().StringSliceS("colormap", "colormap", []string{}, "Shared or Private")
	animateCmd.Flags().CountS("colors", "colors", "preferred number of colors in the image")
	animateCmd.Flags().StringSliceS("colorspace", "colorspace", []string{}, "alternate image colorspace")
	animateCmd.Flags().StringSliceS("debug", "debug", []string{}, "display copious debugging information")
	animateCmd.Flags().StringSliceS("define", "define", []string{}, "Coder/decoder specific options")
	animateCmd.Flags().StringSliceS("delay", "delay", []string{}, "display the next image after pausing")
	animateCmd.Flags().StringSliceS("density", "density", []string{}, "horizontal and vertical density of the image")
	animateCmd.Flags().StringSliceS("depth", "depth", []string{}, "image depth")
	animateCmd.Flags().StringSliceS("display", "display", []string{}, "display image to this X server")
	animateCmd.Flags().CountS("dither", "dither", "apply Floyd/Steinberg error diffusion to image")
	animateCmd.Flags().StringSliceS("font", "font", []string{}, "use this font when annotating the image with text")
	animateCmd.Flags().StringSliceS("foreground", "foreground", []string{}, "define the foreground color")
	animateCmd.Flags().StringSliceS("gamma", "gamma", []string{}, "level of gamma correction")
	animateCmd.Flags().StringSliceS("geometry", "geometry", []string{}, "preferred size and location of the cropped image")
	animateCmd.Flags().BoolS("help", "help", false, "print program options")
	animateCmd.Flags().StringSliceS("iconGeometry", "iconGeometry", []string{}, "specify the icon geometry")
	animateCmd.Flags().StringSliceS("iconic", "iconic", []string{}, "iconic animation")
	animateCmd.Flags().StringSliceS("interlace", "interlace", []string{}, "None, Line, Plane, or Partition")
	animateCmd.Flags().StringSliceS("limit", "limit", []string{}, "Disk, File, Map, Memory, Pixels, Width, Height, Threads, Read, or Write resource limit")
	animateCmd.Flags().StringSliceS("log", "log", []string{}, "format of debugging information")
	animateCmd.Flags().StringSliceS("map", "map", []string{}, "display image using this Standard Colormap")
	animateCmd.Flags().CountS("matte", "matte", "store matte channel if the image has one")
	animateCmd.Flags().StringSliceS("mattecolor", "mattecolor", []string{}, "specify the color to be used with the -frame option")
	animateCmd.Flags().CountS("monitor", "monitor", "show progress indication")
	animateCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	animateCmd.Flags().StringSliceS("name", "name", []string{}, "name an image")
	animateCmd.Flags().CountS("noop", "noop", "do not apply options to image")
	animateCmd.Flags().CountS("pause", "pause", "seconds to pause before reanimating")
	animateCmd.Flags().StringSliceS("remote", "remote", []string{}, "execute a command in a remote display process")
	animateCmd.Flags().StringSliceS("rotate", "rotate", []string{}, "apply Paeth rotation to the image")
	animateCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", []string{}, "horizontal and vertical sampling factors")
	animateCmd.Flags().StringSliceS("scenes", "scenes", []string{}, "image scene range")
	animateCmd.Flags().CountS("shared-memory", "shared-memory", "use shared memory")
	animateCmd.Flags().StringSliceS("size", "size", []string{}, "width and height of image")
	animateCmd.Flags().StringSliceS("title", "title", []string{}, "assign title to displayed image")
	animateCmd.Flags().StringSliceS("treedepth", "treedepth", []string{}, "color tree depth")
	animateCmd.Flags().CountS("trim", "trim", "trim image edges")
	animateCmd.Flags().StringSliceS("type", "type", []string{}, "image type")
	animateCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	animateCmd.Flags().BoolS("version", "version", false, "print version information")
	animateCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", []string{}, "Constant, Edge, Mirror, or Tile")
	animateCmd.Flags().StringSliceS("visual", "visual", []string{}, "display image using this visual type")
	animateCmd.Flags().StringSliceS("window", "window", []string{}, "display image to background of this window")
	rootCmd.AddCommand(animateCmd)

	carapace.Gen(animateCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
