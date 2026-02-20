package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resizeOsWindowCmd = &cobra.Command{
	Use:   "resize-os-window",
	Short: "Resize/show/hide/etc. the specified OS Windows",
}

func init() {
	resizeOsWindowCmd.AddCommand(atCmd)
	carapace.Gen(resizeOsWindowCmd).Standalone()

	resizeOsWindowCmd.Flags().String("action", "resize", "The action to perform")
	resizeOsWindowCmd.Flags().String("height", "0", "Change the height of the window")
	resizeOsWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	resizeOsWindowCmd.Flags().Bool("incremental", false, "Treat the specified sizes as increments on the existing window size")
	resizeOsWindowCmd.Flags().StringP("match", "m", "", "The window to match")
	resizeOsWindowCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")
	resizeOsWindowCmd.Flags().Bool("self", false, "Resize the window this command is run in, rather than the active window")
	resizeOsWindowCmd.Flags().String("unit", "cells", "The unit in which to interpret specified sizes")
	resizeOsWindowCmd.Flags().String("width", "0", "Change the width of the window")

	carapace.Gen(resizeOsWindowCmd).FlagCompletion(carapace.ActionMap{})
}