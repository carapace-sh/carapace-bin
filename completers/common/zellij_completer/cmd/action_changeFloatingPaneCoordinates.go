package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_changeFloatingPaneCoordinatesCmd = &cobra.Command{
	Use:   "change-floating-pane-coordinates",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_changeFloatingPaneCoordinatesCmd).Standalone()

	action_changeFloatingPaneCoordinatesCmd.Flags().StringP("borderless", "b", "", "change this pane to be with/without a border (warning: will make it impossible to move with the mouse if without a border)")
	action_changeFloatingPaneCoordinatesCmd.Flags().String("height", "", "The height if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	action_changeFloatingPaneCoordinatesCmd.Flags().BoolP("help", "h", false, "Print help")
	action_changeFloatingPaneCoordinatesCmd.Flags().StringP("pane-id", "p", "", "The pane_id of the floating pane, eg.  terminal_1, plugin_2 or 3 (equivalent to terminal_3)")
	action_changeFloatingPaneCoordinatesCmd.Flags().String("pinned", "", "Whether to pin a floating pane so that it is always on top")
	action_changeFloatingPaneCoordinatesCmd.Flags().String("width", "", "The width if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	action_changeFloatingPaneCoordinatesCmd.Flags().StringP("x", "x", "", "The x coordinates if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	action_changeFloatingPaneCoordinatesCmd.Flags().StringP("y", "y", "", "The y coordinates if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	action_changeFloatingPaneCoordinatesCmd.MarkFlagRequired("pane-id")
	actionCmd.AddCommand(action_changeFloatingPaneCoordinatesCmd)

	carapace.Gen(action_changeFloatingPaneCoordinatesCmd).FlagCompletion(carapace.ActionMap{
		"borderless": carapace.ActionValues("true", "false"),
		"pane-id":    zellij.ActionPanes(),
		"pinned":     carapace.ActionValues("true", "false"),
	})
}
