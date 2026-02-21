package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var enablePortalCmd = &cobra.Command{
	Use:   "enable-portal",
	Short: "Create or edit files needed so that the portal from this kitten is used by xdg-desktop-portal",
}

func init() {
	enablePortalCmd.AddCommand(desktopUiCmd)
	carapace.Gen(enablePortalCmd).Standalone()

	enablePortalCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(enablePortalCmd).FlagCompletion(carapace.ActionMap{})
}
