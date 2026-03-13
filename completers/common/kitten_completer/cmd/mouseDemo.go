package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mouseDemoCmd = &cobra.Command{
	Use:   "mouse-demo",
	Short: "Demo the mouse handling kitty implements for terminal programs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mouseDemoCmd).Standalone()

	mouseDemoCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	rootCmd.AddCommand(mouseDemoCmd)
}
