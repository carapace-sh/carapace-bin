package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rubyCmd = &cobra.Command{
	Use:     "ruby",
	Short:   "Run a Ruby instance with Homebrew's libraries loaded",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rubyCmd).Standalone()

	rubyCmd.Flags().Bool("debug", false, "Display any debugging information.")
	rubyCmd.Flags().BoolS("e", "e", false, "Execute the given text string as a script.")
	rubyCmd.Flags().Bool("help", false, "Show this message.")
	rubyCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	rubyCmd.Flags().BoolS("r", "r", false, "Load a library using `require`.")
	rubyCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(rubyCmd)
}
