package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getwindowfocusCmd = &cobra.Command{
	Use:   "getwindowfocus",
	Short: "Prints the window id of the currently focused window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getwindowfocusCmd).Standalone()
	getwindowfocusCmd.Flags().BoolS("f", "f", false, "Report the window with focus even if it is not a top-level window")

	rootCmd.AddCommand(getwindowfocusCmd)
}
