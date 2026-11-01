package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoModeCmd = &cobra.Command{
	Use:   "auto-mode",
	Short: "Inspect or reset auto mode classifier configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoModeCmd).Standalone()

	autoModeCmd.Flags().BoolP("help", "h", false, "Display help for command")
	rootCmd.AddCommand(autoModeCmd)
}
