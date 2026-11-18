package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flagCellarCmd = &cobra.Command{
	Use:   "cellar",
	Short: "Display Homebrew's Cellar path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flagCellarCmd).Standalone()

	flagCellarCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	flagCellarCmd.Flags().BoolP("help", "h", false, "Show this message.")
	flagCellarCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	flagCellarCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
}
