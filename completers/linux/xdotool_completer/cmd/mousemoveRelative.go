package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mousemoveRelativeCmd = &cobra.Command{
	Use:   "mousemove_relative",
	Short: "Move the mouse x,y pixels relative to the current position of the mouse cursor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mousemoveRelativeCmd).Standalone()

	mousemoveRelativeCmd.Flags().BoolP("clearmodifiers", "c", false, "See CLEARMODIFIERS")
	mousemoveRelativeCmd.Flags().BoolP("polar", "p", false, "Use polar coordinates")
	mousemoveRelativeCmd.Flags().Bool("sync", false, "After sending the mouse move request, wait until the mouse is actually moved")
	rootCmd.AddCommand(mousemoveRelativeCmd)
}
