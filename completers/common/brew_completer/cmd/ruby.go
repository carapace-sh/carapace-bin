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

	rubyCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	rubyCmd.Flags().StringS("e", "e", "", "Execute the given text string as a script.")
	rubyCmd.Flags().BoolP("help", "h", false, "Show this message.")
	rubyCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	rubyCmd.Flags().StringS("r", "r", "", "Load a library using `require`.")
	rubyCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(rubyCmd)
}
