package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_setLightThemeCmd = &cobra.Command{
	Use:   "set-light-theme",
	Short: "Switch the theme to light (uses configured `theme_light`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_setLightThemeCmd).Standalone()

	action_setLightThemeCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_setLightThemeCmd)
}
