package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var desktopUi_setContrastCmd = &cobra.Command{
	Use:   "set-contrast",
	Short: "Change the contrast",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(desktopUi_setContrastCmd).Standalone()

	desktopUi_setContrastCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	desktopUiCmd.AddCommand(desktopUi_setContrastCmd)

	carapace.Gen(desktopUi_setContrastCmd).PositionalCompletion(
		carapace.ActionValues("high", "normal"),
	)
}
