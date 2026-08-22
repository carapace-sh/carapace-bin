package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_newTabCmd = &cobra.Command{
	Use:   "new-tab",
	Short: "Create a new tab, optionally with a specified tab layout and name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_newTabCmd).Standalone()

	action_newTabCmd.Flags().Bool("block-until-exit", false, "Block until the command exits (regardless of exit status) OR its pane has been closed")
	action_newTabCmd.Flags().Bool("block-until-exit-failure", false, "Block until the command exits with failure (non-zero exit status) OR its pane has been closed")
	action_newTabCmd.Flags().Bool("block-until-exit-success", false, "Block until the command exits successfully (exit status 0) OR its pane has been closed")
	action_newTabCmd.Flags().Bool("close-on-exit", false, "Close the pane immediately when its command exits")
	action_newTabCmd.Flags().StringP("cwd", "c", "", "Change the working directory of the new tab")
	action_newTabCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	action_newTabCmd.Flags().String("initial-plugin", "", "Initial plugin to load in the new tab")
	action_newTabCmd.Flags().StringP("layout", "l", "", "Layout to use for the new tab")
	action_newTabCmd.Flags().String("layout-dir", "", "Default folder to look for layouts")
	action_newTabCmd.Flags().String("layout-string", "", "Raw KDL layout string to use directly (instead of a layout file path)")
	action_newTabCmd.Flags().StringP("name", "n", "", "Name of the new tab")
	action_newTabCmd.Flags().Bool("no-focus", false, "if set, will create the tab without changing the focus of any client")
	action_newTabCmd.Flags().Bool("start-suspended", false, "Start the command suspended, only running it after you first press ENTER")
	actionCmd.AddCommand(action_newTabCmd)

	carapace.Gen(action_newTabCmd).FlagCompletion(carapace.ActionMap{
		"cwd":        carapace.ActionDirectories(),
		"layout":     carapace.ActionFiles(),
		"layout-dir": carapace.ActionDirectories(),
	})
}
