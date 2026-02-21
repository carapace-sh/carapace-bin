package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var disableLigaturesCmd = &cobra.Command{
	Use:   "disable-ligatures",
	Short: "Control ligature rendering",
}

func init() {
	disableLigaturesCmd.AddCommand(atCmd)
	carapace.Gen(disableLigaturesCmd).Standalone()

	disableLigaturesCmd.Flags().BoolP("all", "a", false, "By default, ligatures are only affected in the active window")
	disableLigaturesCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	disableLigaturesCmd.Flags().StringP("match", "m", "", "The window to match")
	disableLigaturesCmd.Flags().StringP("match-tab", "t", "", "The tab to match")

	carapace.Gen(disableLigaturesCmd).FlagCompletion(carapace.ActionMap{})
}
