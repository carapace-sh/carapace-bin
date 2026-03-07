package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_setTabColorCmd = &cobra.Command{
	Use:   "set-tab-color",
	Short: "Change the color of the specified tabs in the tab bar",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_setTabColorCmd).Standalone()

	at_setTabColorCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_setTabColorCmd.Flags().StringP("match", "m", "", "The tab to match")
	at_setTabColorCmd.Flags().Bool("self", false, "Close the tab this command is run in, rather than the active tab")
	atCmd.AddCommand(at_setTabColorCmd)

}
