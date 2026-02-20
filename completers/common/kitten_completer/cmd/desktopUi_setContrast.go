package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setContrastCmd = &cobra.Command{
	Use:   "set-contrast",
	Short: "Change the contrast",
}

func init() {
	setContrastCmd.AddCommand(desktopUiCmd)
	carapace.Gen(setContrastCmd).Standalone()

	setContrastCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(setContrastCmd).FlagCompletion(carapace.ActionMap{})

	carapace.Gen(setContrastCmd).PositionalCompletion(
		carapace.ActionValues("high", "normal"),
	)
}
