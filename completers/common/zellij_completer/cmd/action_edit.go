package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open the specified file in a new zellij pane with your default EDITOR Returns: Created pane ID (format: terminal_<id>)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_editCmd).Standalone()

	action_editCmd.Flags().StringP("borderless", "b", "", "start this pane without a border (warning: will make it impossible to move with the mouse)")
	action_editCmd.Flags().Bool("close-replaced-pane", false, "Close the replaced pane instead of suspending it (only effective with --in-place)")
	action_editCmd.Flags().String("cwd", "", "Change the working directory of the editor")
	action_editCmd.Flags().StringP("direction", "d", "", "Direction to open the new pane in")
	action_editCmd.Flags().BoolP("floating", "f", false, "Open the new pane in floating mode")
	action_editCmd.Flags().String("height", "", "The height if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	action_editCmd.Flags().BoolP("help", "h", false, "Print help")
	action_editCmd.Flags().BoolP("in-place", "i", false, "Open the new pane in place of the current pane, temporarily suspending it")
	action_editCmd.Flags().StringP("line-number", "l", "", "Open the file in the specified line number")
	action_editCmd.Flags().Bool("near-current-pane", false, "if set, will open the pane near the current one rather than following the user's focus")
	action_editCmd.Flags().Bool("no-focus", false, "if set, will open the pane without changing the focus of any client, placing it relative to the pane the command was issued from")
	action_editCmd.Flags().String("pinned", "", "Whether to pin a floating pane so that it is always on top")
	action_editCmd.Flags().String("tab-id", "", "Target a specific tab by ID")
	action_editCmd.Flags().String("width", "", "The width if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	action_editCmd.Flags().StringP("x", "x", "", "The x coordinates if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	action_editCmd.Flags().StringP("y", "y", "", "The y coordinates if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	actionCmd.AddCommand(action_editCmd)

	carapace.Gen(action_editCmd).FlagCompletion(carapace.ActionMap{
		"borderless": carapace.ActionValues("true", "false"),
		"cwd":        carapace.ActionFiles(),
		"direction":  zellij.ActionDirections(),
		"pinned":     carapace.ActionValues("true", "false"),
		"tab-id":     zellij.ActionTabs(),
	})

	carapace.Gen(action_editCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
