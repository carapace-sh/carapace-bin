package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_setTabTitleCmd = &cobra.Command{
	Use:   "set-tab-title",
	Short: "Set the tab title",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_setTabTitleCmd).Standalone()

	at_setTabTitleCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_setTabTitleCmd.Flags().StringP("match", "m", "", "The tab to match")
	atCmd.AddCommand(at_setTabTitleCmd)

}
