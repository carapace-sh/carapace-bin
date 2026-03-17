package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_setFontSizeCmd = &cobra.Command{
	Use:   "set-font-size",
	Short: "Set the font size in the active top-level OS window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_setFontSizeCmd).Standalone()

	at_setFontSizeCmd.Flags().BoolP("all", "a", false, "By default, the font size is only changed in the active OS window, this option will cause it to be changed in all OS windows")
	at_setFontSizeCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	atCmd.AddCommand(at_setFontSizeCmd)

}
