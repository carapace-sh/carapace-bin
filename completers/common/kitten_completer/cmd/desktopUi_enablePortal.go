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

// Usage: kitten desktop-ui enable-portal
//
// Once you run this command, add kitten desktop-ui run-server to your window
// manager startup sequence and reboot your computer (or logout and restart your
// session) and hopefully xdg-desktop-portal should now delegate to kitty for the
// portals implemented here.
//
// Options:
//   --help, -h [=no]
//     Show help for this command
