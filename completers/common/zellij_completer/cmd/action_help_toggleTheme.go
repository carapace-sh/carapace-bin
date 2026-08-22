package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_toggleThemeCmd = &cobra.Command{
	Use:   "toggle-theme",
	Short: "Toggle between dark and light themes (used configured `theme_dark` and `theme_light`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_toggleThemeCmd).Standalone()

	action_helpCmd.AddCommand(action_help_toggleThemeCmd)
}
