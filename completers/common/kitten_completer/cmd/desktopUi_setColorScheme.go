package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setColorSchemeCmd = &cobra.Command{
	Use:   "set-color-scheme",
	Short: "Change the color scheme",
}

func init() {
	setColorSchemeCmd.AddCommand(desktopUiCmd)
	carapace.Gen(setColorSchemeCmd).Standalone()

	setColorSchemeCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(setColorSchemeCmd).FlagCompletion(carapace.ActionMap{})

	carapace.Gen(setColorSchemeCmd).PositionalCompletion(
		carapace.ActionValues("light", "dark", "no-preference", "toggle"),
	)
}

// Usage: kitten desktop-ui set-color-scheme light|dark|no-preference|toggle
//
// Change the color scheme
//
// Options:
//   --help, -h [=no]
//     Show help for this command
