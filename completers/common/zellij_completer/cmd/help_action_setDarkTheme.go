package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_setDarkThemeCmd = &cobra.Command{
	Use:   "set-dark-theme",
	Short: "Switch the theme to dark (uses configured `theme_dark`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_setDarkThemeCmd).Standalone()

	help_actionCmd.AddCommand(help_action_setDarkThemeCmd)
}
