package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_setDarkThemeCmd = &cobra.Command{
	Use:   "set-dark-theme",
	Short: "Switch the theme to dark (uses configured `theme_dark`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_setDarkThemeCmd).Standalone()

	action_setDarkThemeCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_setDarkThemeCmd)
}
