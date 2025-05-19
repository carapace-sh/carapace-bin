package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var homepageCmd = &cobra.Command{
	Use:   "homepage",
	Short: "Open a <formula> or <cask>'s homepage in a browser, or open Homebrew's own homepage if no argument is provided",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(homepageCmd).Standalone()

	homepageCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	homepageCmd.Flags().Bool("debug", false, "Display any debugging information.")
	homepageCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	homepageCmd.Flags().Bool("help", false, "Show this message.")
	homepageCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	homepageCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(homepageCmd)
}
