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

	flagCellarCmd.Flags().Bool("debug", false, "Display any debugging information.")
	flagCellarCmd.Flags().Bool("help", false, "Show this message.")
	flagCellarCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	flagCellarCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
}
