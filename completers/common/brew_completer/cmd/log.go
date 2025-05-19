package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:     "log",
	Short:   "Show the `git log` for <formula> or <cask>, or show the log for the Homebrew repository if no formula or cask is provided",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()

	logCmd.Flags().BoolS("1", "1", false, "Print only one commit.")
	logCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	logCmd.Flags().Bool("debug", false, "Display any debugging information.")
	logCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	logCmd.Flags().Bool("help", false, "Show this message.")
	logCmd.Flags().Bool("max-count", false, "Print only a specified number of commits.")
	logCmd.Flags().Bool("oneline", false, "Print only one line per commit.")
	logCmd.Flags().Bool("patch", false, "Also print patch from commit.")
	logCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	logCmd.Flags().Bool("stat", false, "Also print diffstat from commit.")
	logCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(logCmd)
}
