package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_setSpacingCmd = &cobra.Command{
	Use:   "set-spacing",
	Short: "Set window paddings and margins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_setSpacingCmd).Standalone()

	at_setSpacingCmd.Flags().BoolP("all", "a", false, "By default, settings are only changed for the currently active window")
	at_setSpacingCmd.Flags().BoolP("configured", "c", false, "Also change the configured paddings and margins")
	at_setSpacingCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_setSpacingCmd.Flags().StringP("match", "m", "", "The window to match")
	at_setSpacingCmd.Flags().StringP("match-tab", "t", "", "The tab to match")
	atCmd.AddCommand(at_setSpacingCmd)

}
