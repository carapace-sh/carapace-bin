package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var linkageCmd = &cobra.Command{
	Use:     "linkage",
	Short:   "Check the library links from the given <formula> kegs",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkageCmd).Standalone()

	linkageCmd.Flags().Bool("cached", false, "Print the cached linkage values stored in `HOMEBREW_CACHE`, set by a previous `brew linkage` run.")
	linkageCmd.Flags().Bool("debug", false, "Display any debugging information.")
	linkageCmd.Flags().Bool("help", false, "Show this message.")
	linkageCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	linkageCmd.Flags().Bool("reverse", false, "For every library that a keg references, print its dylib path followed by the binaries that link to it.")
	linkageCmd.Flags().Bool("strict", false, "Exit with a non-zero status if any undeclared dependencies with linkage are found.")
	linkageCmd.Flags().Bool("test", false, "Show only missing libraries and exit with a non-zero status if any missing libraries are found.")
	linkageCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(linkageCmd)
}
