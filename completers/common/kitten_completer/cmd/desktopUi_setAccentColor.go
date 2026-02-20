package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setAccentColorCmd = &cobra.Command{
	Use:   "set-accent-color",
	Short: "Change the accent color",
}

func init() {
	setAccentColorCmd.AddCommand(desktopUiCmd)
	carapace.Gen(setAccentColorCmd).Standalone()

	setAccentColorCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(setAccentColorCmd).FlagCompletion(carapace.ActionMap{})
}

// Usage: kitten desktop-ui set-accent-color color_as_hex_or_name
//
// Change the accent color
//
// Options:
//   --help, -h [=no]
//     Show help for this command
