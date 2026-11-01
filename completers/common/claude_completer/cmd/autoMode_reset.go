package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoMode_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset auto mode configuration to the shipped defaults by removing the autoMode section from your…",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoMode_resetCmd).Standalone()

	autoMode_resetCmd.Flags().BoolP("help", "h", false, "Display help for command")
	autoMode_resetCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	autoModeCmd.AddCommand(autoMode_resetCmd)
}
