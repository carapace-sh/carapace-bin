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
