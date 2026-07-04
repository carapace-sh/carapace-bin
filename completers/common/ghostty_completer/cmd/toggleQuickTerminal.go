package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var toggleQuickTerminalCmd = &cobra.Command{
	Use:   "+toggle-quick-terminal",
	Short: "Toggle the visibility of the quick terminal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toggleQuickTerminalCmd).Standalone()

	toggleQuickTerminalCmd.Flags().String("class", "", "If set, connect to a custom instance of Ghostty")
	toggleQuickTerminalCmd.Flags().Bool("help", false, "show help")
	rootCmd.AddCommand(toggleQuickTerminalCmd)
}
