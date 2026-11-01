package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_toggleThemeCmd = &cobra.Command{
	Use:   "toggle-theme",
	Short: "Toggle between dark and light themes (used configured `theme_dark` and `theme_light`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_toggleThemeCmd).Standalone()

	action_toggleThemeCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_toggleThemeCmd)
}
