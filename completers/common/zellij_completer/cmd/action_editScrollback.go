package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_editScrollbackCmd = &cobra.Command{
	Use:   "edit-scrollback",
	Short: "Open the pane scrollback in your default editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_editScrollbackCmd).Standalone()

	action_editScrollbackCmd.Flags().BoolP("ansi", "a", false, "Preserve ANSI styling in the scrollback dump")
	action_editScrollbackCmd.Flags().BoolP("help", "h", false, "Print help")
	action_editScrollbackCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_editScrollbackCmd)

	carapace.Gen(action_editScrollbackCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": zellij.ActionSelectablePanes(),
	})
}
