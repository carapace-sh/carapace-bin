package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flagPrefixCmd = &cobra.Command{
	Use:   "prefix",
	Short: "Display Homebrew's install path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flagPrefixCmd).Standalone()

	flagPrefixCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	flagPrefixCmd.Flags().BoolP("help", "h", false, "Show this message.")
	flagPrefixCmd.Flags().Bool("installed", false, "Outputs nothing and returns a failing status code if <formula> is not installed.")
	flagPrefixCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	flagPrefixCmd.Flags().Bool("unbrewed", false, "List files in Homebrew's prefix not installed by Homebrew.")
	flagPrefixCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
}
