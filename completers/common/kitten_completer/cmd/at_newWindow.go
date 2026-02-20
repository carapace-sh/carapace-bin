package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var newWindowCmd = &cobra.Command{
	Use:   "new-window",
	Short: "Open new window",
}

func init() {
	newWindowCmd.AddCommand(atCmd)
	carapace.Gen(newWindowCmd).Standalone()

	newWindowCmd.Flags().String("cwd", "", "The initial working directory for the new window")
	newWindowCmd.Flags().Bool("dont-take-focus", false, "Keep the current window focused instead of switching to the newly opened window")
	newWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	newWindowCmd.Flags().StringP("match", "m", "", "The tab to match")
	newWindowCmd.Flags().Bool("new-tab", false, "Open a new tab")
	newWindowCmd.Flags().Bool("no-response", false, "Don't wait for a response giving the id of the newly opened window")
	newWindowCmd.Flags().String("tab-title", "", "Set the title of the tab, when open a new tab")
	newWindowCmd.Flags().String("title", "", "The title for the new window")
	newWindowCmd.Flags().String("window-type", "kitty", "What kind of window to open")

	carapace.Gen(newWindowCmd).FlagCompletion(carapace.ActionMap{
		"window-type": carapace.ActionValues("kitty", "os"),
	})
}