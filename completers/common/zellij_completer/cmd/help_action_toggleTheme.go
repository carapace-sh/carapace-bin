package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_toggleThemeCmd = &cobra.Command{
	Use:   "toggle-theme",
	Short: "Toggle between dark and light themes (used configured `theme_dark` and `theme_light`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_toggleThemeCmd).Standalone()

	help_actionCmd.AddCommand(help_action_toggleThemeCmd)
}
