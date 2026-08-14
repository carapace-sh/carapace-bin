package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_newPaneCmd = &cobra.Command{
	Use:   "new-pane",
	Short: "Open a new pane in the specified direction [right|down] If no direction is specified, will try to use the biggest available space. Returns: Created pane ID (format: terminal_<id> or plugin_<id>)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_newPaneCmd).Standalone()

	action_newPaneCmd.Flags().Bool("block-until-exit", false, "Block until the command exits (regardless of exit status) OR its pane has been closed")
	action_newPaneCmd.Flags().Bool("block-until-exit-failure", false, "Block until the command exits with failure (non-zero exit status) OR its pane has been closed")
	action_newPaneCmd.Flags().Bool("block-until-exit-success", false, "Block until the command exits successfully (exit status 0) OR its pane has been closed")
	action_newPaneCmd.Flags().BoolP("blocking", "b", false, "Block until the command has finished and its pane has been closed")
	action_newPaneCmd.Flags().String("borderless", "", "start this pane without a border (warning: will make it impossible to move with the mouse)")
	action_newPaneCmd.Flags().BoolP("close-on-exit", "c", false, "Close the pane immediately when its command exits")
	action_newPaneCmd.Flags().Bool("close-replaced-pane", false, "Close the replaced pane instead of suspending it (only effective with --in-place)")
	action_newPaneCmd.Flags().String("configuration", "", "")
	action_newPaneCmd.Flags().String("cwd", "", "Change the working directory of the new pane")
	action_newPaneCmd.Flags().StringP("direction", "d", "", "Direction to open the new pane in")
	action_newPaneCmd.Flags().BoolP("floating", "f", false, "Open the new pane in floating mode")
	action_newPaneCmd.Flags().String("height", "", "The height if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	action_newPaneCmd.Flags().BoolP("help", "h", false, "Print help")
	action_newPaneCmd.Flags().BoolP("in-place", "i", false, "Open the new pane in place of the current pane, temporarily suspending it")
	action_newPaneCmd.Flags().StringP("name", "n", "", "Name of the new pane")
	action_newPaneCmd.Flags().Bool("near-current-pane", false, "if set, will open the pane near the current one rather than following the user's focus")
	action_newPaneCmd.Flags().Bool("no-focus", false, "if set, will open the pane without changing the focus of any client, placing it relative to the pane the command was issued from")
	action_newPaneCmd.Flags().String("pane-id", "", "The pane to replace when opening in place, eg. terminal_1, plugin_2 or 3 (only effective with --in-place; defaults to the focused pane)")
	action_newPaneCmd.Flags().String("pinned", "", "Whether to pin a floating pane so that it is always on top")
	action_newPaneCmd.Flags().StringP("plugin", "p", "", "")
	action_newPaneCmd.Flags().Bool("skip-plugin-cache", false, "")
	action_newPaneCmd.Flags().Bool("stacked", false, "")
	action_newPaneCmd.Flags().BoolP("start-suspended", "s", false, "Start the command suspended, only running it after the you first press ENTER")
	action_newPaneCmd.Flags().String("tab-id", "", "Target a specific tab by ID")
	action_newPaneCmd.Flags().String("width", "", "The width if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	action_newPaneCmd.Flags().StringP("x", "x", "", "The x coordinates if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	action_newPaneCmd.Flags().StringP("y", "y", "", "The y coordinates if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	actionCmd.AddCommand(action_newPaneCmd)

	carapace.Gen(action_newPaneCmd).FlagCompletion(carapace.ActionMap{
		"borderless": carapace.ActionValues("true", "false"),
		"cwd":        carapace.ActionFiles(),
		"direction":  actionDirections(),
		"pinned":     carapace.ActionValues("true", "false"),
	})
}
