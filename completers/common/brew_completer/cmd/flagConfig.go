package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flagConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Show Homebrew and system configuration info useful for debugging",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flagConfigCmd).Standalone()

	flagConfigCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	flagConfigCmd.Flags().BoolP("help", "h", false, "Show this message.")
	flagConfigCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	flagConfigCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
}
