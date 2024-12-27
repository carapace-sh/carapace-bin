package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listKeybindsCmd = &cobra.Command{
	Use:   "+list-keybinds",
	Short: "list all the available keybinds",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listKeybindsCmd).Standalone()

	listKeybindsCmd.Flags().Bool("default", false, "print out all the default keybinds")
	listKeybindsCmd.Flags().Bool("plain", false, "disable formatting")
	rootCmd.AddCommand(listKeybindsCmd)
}
