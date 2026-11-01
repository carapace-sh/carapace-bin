package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_renamePaneCmd = &cobra.Command{
	Use:   "rename-pane",
	Short: "Renames the focused pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_renamePaneCmd).Standalone()

	action_renamePaneCmd.Flags().BoolP("help", "h", false, "Print help")
	action_renamePaneCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_renamePaneCmd)

	carapace.Gen(action_renamePaneCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": zellij.ActionSelectablePanes(),
	})
}
