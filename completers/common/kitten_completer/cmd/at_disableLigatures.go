package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_disableLigaturesCmd = &cobra.Command{
	Use:   "disable-ligatures",
	Short: "Control ligature rendering",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_disableLigaturesCmd).Standalone()

	at_disableLigaturesCmd.Flags().BoolP("all", "a", false, "By default, ligatures are only affected in the active window")
	at_disableLigaturesCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_disableLigaturesCmd.Flags().StringP("match", "m", "", "The window to match")
	at_disableLigaturesCmd.Flags().StringP("match-tab", "t", "", "The tab to match")
	atCmd.AddCommand(at_disableLigaturesCmd)

}
