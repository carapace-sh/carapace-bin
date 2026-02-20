package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var desktopUiCmd = &cobra.Command{
	Use:   "desktop-ui",
	Short: "Implement various desktop components for use with lightweight compositors/window managers on Linux",
}

func init() {
	rootCmd.AddCommand(desktopUiCmd)
	carapace.Gen(desktopUiCmd).Standalone()

	desktopUiCmd.Flags().BoolP("help", "h", false, "Show help for this command")
}

// TODO: subcommands
// Usage: kitten desktop-ui

// Implement various desktop components for use with lightweight compositors/window
// managers on Linux

// Commands:
//    run-server
//     Start the various servers used to integrate with the Linux desktop
//    enable-portal
//     This will create or edit the various files needed so that the portal from
//     this kitten is used by xdg-desktop-portal
//    set-color-scheme
//     Change the color scheme
//    set-accent-color
//     Change the accent color
//    set-contrast
//     Change the contrast. Can be high or normal.
//    set-setting
//     Change an arbitrary setting
//    show-settings
//     Print the current values of the desktop settings

// Get help for an individual command by running:
//     kitten desktop-ui command -h

// Options:
//   --help, -h [=no]
//     Show help for this command

// kitten desktop-ui 0.45.0 created by Kovid Goyal
