package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mouseDemoCmd = &cobra.Command{
	Use:   "mouse-demo",
	Short: "Demo the mouse handling kitty implements for terminal programs",
}

func init() {
	rootCmd.AddCommand(mouseDemoCmd)
	carapace.Gen(mouseDemoCmd).Standalone()

	mouseDemoCmd.Flags().BoolP("help", "h", false, "Show help for this command")
}
