package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_newWindowCmd = &cobra.Command{
	Use:   "new-window",
	Short: "Open new window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_newWindowCmd).Standalone()

	at_newWindowCmd.Flags().String("cwd", "", "The initial working directory for the new window")
	at_newWindowCmd.Flags().Bool("dont-take-focus", false, "Keep the current window focused instead of switching to the newly opened window")
	at_newWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_newWindowCmd.Flags().StringP("match", "m", "", "The tab to match")
	at_newWindowCmd.Flags().Bool("new-tab", false, "Open a new tab")
	at_newWindowCmd.Flags().Bool("no-response", false, "Don't wait for a response giving the id of the newly opened window")
	at_newWindowCmd.Flags().String("tab-title", "", "Set the title of the tab, when open a new tab")
	at_newWindowCmd.Flags().String("title", "", "The title for the new window")
	at_newWindowCmd.Flags().String("window-type", "kitty", "What kind of window to open")
	atCmd.AddCommand(at_newWindowCmd)

	carapace.Gen(at_newWindowCmd).FlagCompletion(carapace.ActionMap{
		"window-type": carapace.ActionValues("kitty", "os"),
	})
}
