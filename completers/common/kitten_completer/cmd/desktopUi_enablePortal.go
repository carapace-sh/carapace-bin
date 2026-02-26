package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var desktopUi_enablePortalCmd = &cobra.Command{
	Use:   "enable-portal",
	Short: "Create or edit files needed so that the portal from this kitten is used by xdg-desktop-portal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(desktopUi_enablePortalCmd).Standalone()

	desktopUi_enablePortalCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	desktopUiCmd.AddCommand(desktopUi_enablePortalCmd)

}
