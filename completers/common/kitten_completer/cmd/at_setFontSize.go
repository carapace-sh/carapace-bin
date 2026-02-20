package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setFontSizeCmd = &cobra.Command{
	Use:   "set-font-size",
	Short: "Set the font size in the active top-level OS window",
}

func init() {
	setFontSizeCmd.AddCommand(atCmd)
	carapace.Gen(setFontSizeCmd).Standalone()

	setFontSizeCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(setFontSizeCmd).FlagCompletion(carapace.ActionMap{})
}