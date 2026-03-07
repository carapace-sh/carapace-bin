package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_setColorsCmd = &cobra.Command{
	Use:   "set-colors",
	Short: "Set terminal colors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_setColorsCmd).Standalone()

	at_setColorsCmd.Flags().BoolP("all", "a", false, "By default, colors are only changed for the currently active window")
	at_setColorsCmd.Flags().BoolP("configured", "c", false, "Also change the configured colors")
	at_setColorsCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_setColorsCmd.Flags().StringP("match", "m", "", "The window to match")
	at_setColorsCmd.Flags().StringP("match-tab", "t", "", "The tab to match")
	at_setColorsCmd.Flags().Bool("reset", false, "Restore all colors to the values they had at kitty startup")
	atCmd.AddCommand(at_setColorsCmd)

}
