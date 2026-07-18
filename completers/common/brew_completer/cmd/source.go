package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sourceCmd = &cobra.Command{
	Use:     "source",
	Short:   "Open a <formula>'s source repository in a browser, or open Homebrew's own repository if no argument is provided",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sourceCmd).Standalone()

	sourceCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	sourceCmd.Flags().Bool("debug", false, "Display any debugging information.")
	sourceCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	sourceCmd.Flags().Bool("help", false, "Show this message.")
	sourceCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	sourceCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(sourceCmd)
}
