package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_setLightThemeCmd = &cobra.Command{
	Use:   "set-light-theme",
	Short: "Switch the theme to light (uses configured `theme_light`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_setLightThemeCmd).Standalone()

	action_helpCmd.AddCommand(action_help_setLightThemeCmd)
}
