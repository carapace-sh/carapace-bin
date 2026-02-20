package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setWindowLogoCmd = &cobra.Command{
	Use:   "set-window-logo",
	Short: "Set the window logo",
}

func init() {
	setWindowLogoCmd.AddCommand(atCmd)
	carapace.Gen(setWindowLogoCmd).Standalone()

	setWindowLogoCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(setWindowLogoCmd).FlagCompletion(carapace.ActionMap{})
}