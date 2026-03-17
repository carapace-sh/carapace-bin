package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var desktopUi_setColorSchemeCmd = &cobra.Command{
	Use:   "set-color-scheme",
	Short: "Change the color scheme",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(desktopUi_setColorSchemeCmd).Standalone()

	desktopUi_setColorSchemeCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	desktopUiCmd.AddCommand(desktopUi_setColorSchemeCmd)

	carapace.Gen(desktopUi_setColorSchemeCmd).PositionalCompletion(
		carapace.ActionValues("light", "dark", "no-preference", "toggle"),
	)
}
