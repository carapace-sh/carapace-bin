package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_setLightThemeCmd = &cobra.Command{
	Use:   "set-light-theme",
	Short: "Switch the theme to light (uses configured `theme_light`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_setLightThemeCmd).Standalone()

	help_actionCmd.AddCommand(help_action_setLightThemeCmd)
}
