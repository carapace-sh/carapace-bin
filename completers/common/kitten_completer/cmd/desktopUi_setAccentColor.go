package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var desktopUi_setAccentColorCmd = &cobra.Command{
	Use:   "set-accent-color",
	Short: "Change the accent color",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(desktopUi_setAccentColorCmd).Standalone()

	desktopUi_setAccentColorCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	desktopUiCmd.AddCommand(desktopUi_setAccentColorCmd)

}
