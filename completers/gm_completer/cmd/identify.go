package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identifyCmd = &cobra.Command{
	Use:   "identify",
	Short: "describe an image or image sequence",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identifyCmd).Standalone()

	identifyCmd.Flags().StringSliceS("debug", "debug", []string{}, "display copious debugging information")
	identifyCmd.Flags().StringSliceS("define", "define", []string{}, "Coder/decoder specific options")
	identifyCmd.Flags().StringSliceS("density", "density", []string{}, "horizontal and vertical density of the image")
	identifyCmd.Flags().StringSliceS("depth", "depth", []string{}, "image depth")
	identifyCmd.Flags().StringSliceS("format", "format", []string{}, "output formatted image characteristics")
	identifyCmd.Flags().BoolS("help", "help", false, "print program options")
	identifyCmd.Flags().StringSliceS("interlace", "interlace", []string{}, "None, Line, Plane, or Partition")
	identifyCmd.Flags().StringSliceS("log", "log", []string{}, "format of debugging information")
	identifyCmd.Flags().CountS("monitor", "monitor", "show progress indication")
	identifyCmd.Flags().CountS("ping", "ping", "efficiently determine image attributes")
	identifyCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", []string{}, "horizontal and vertical sampling factors")
	identifyCmd.Flags().StringSliceS("size", "size", []string{}, "width and height of image")
	identifyCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	identifyCmd.Flags().BoolS("version", "version", false, "print version information")
	identifyCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", []string{}, "Constant, Edge, Mirror, or Tile")
	rootCmd.AddCommand(identifyCmd)

	carapace.Gen(identifyCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
