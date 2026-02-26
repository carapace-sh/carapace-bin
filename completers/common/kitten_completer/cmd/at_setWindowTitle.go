package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_setWindowTitleCmd = &cobra.Command{
	Use:   "set-window-title",
	Short: "Set the window title",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_setWindowTitleCmd).Standalone()

	at_setWindowTitleCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_setWindowTitleCmd.Flags().StringP("match", "m", "", "The window to match")
	at_setWindowTitleCmd.Flags().Bool("temporary", false, "By default, the title will be permanently changed and programs running in the window will not be able to change it again")
	atCmd.AddCommand(at_setWindowTitleCmd)

}
