package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var homeCmd = &cobra.Command{
	Use:     "home",
	Short:   "Open a <formula> or <cask>'s homepage in a browser, or open Homebrew's own homepage if no argument is provided",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(homeCmd).Standalone()

	homeCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	homeCmd.Flags().Bool("debug", false, "Display any debugging information.")
	homeCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	homeCmd.Flags().Bool("help", false, "Show this message.")
	homeCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	homeCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(homeCmd)
}
