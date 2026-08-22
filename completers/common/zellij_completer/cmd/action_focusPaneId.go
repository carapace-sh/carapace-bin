package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_focusPaneIdCmd = &cobra.Command{
	Use:   "focus-pane-id",
	Short: "Focus a specific pane by its ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_focusPaneIdCmd).Standalone()

	action_focusPaneIdCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_focusPaneIdCmd)

	carapace.Gen(action_focusPaneIdCmd).PositionalCompletion(
		zellij.ActionSelectablePanes(),
	)
}
