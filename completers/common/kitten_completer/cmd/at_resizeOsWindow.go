package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_resizeOsWindowCmd = &cobra.Command{
	Use:   "resize-os-window",
	Short: "Resize/show/hide/etc. the specified OS Windows",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_resizeOsWindowCmd).Standalone()

	at_resizeOsWindowCmd.Flags().String("action", "resize", "The action to perform")
	at_resizeOsWindowCmd.Flags().String("height", "0", "Change the height of the window")
	at_resizeOsWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_resizeOsWindowCmd.Flags().Bool("incremental", false, "Treat the specified sizes as increments on the existing window size")
	at_resizeOsWindowCmd.Flags().StringP("match", "m", "", "The window to match")
	at_resizeOsWindowCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")
	at_resizeOsWindowCmd.Flags().Bool("self", false, "Resize the window this command is run in, rather than the active window")
	at_resizeOsWindowCmd.Flags().String("unit", "cells", "The unit in which to interpret specified sizes")
	at_resizeOsWindowCmd.Flags().String("width", "0", "Change the width of the window")
	atCmd.AddCommand(at_resizeOsWindowCmd)

}
