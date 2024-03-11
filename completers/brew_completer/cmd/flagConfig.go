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

	flagConfigCmd.Flags().Bool("debug", false, "Display any debugging information.")
	flagConfigCmd.Flags().Bool("help", false, "Show this message.")
	flagConfigCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	flagConfigCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
}
