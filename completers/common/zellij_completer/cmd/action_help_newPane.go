package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_newPaneCmd = &cobra.Command{
	Use:   "new-pane",
	Short: "Open a new pane in the specified direction [right|down] If no direction is specified, will try to use the biggest available space. Returns: Created pane ID (format: terminal_<id> or plugin_<id>)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_newPaneCmd).Standalone()

	action_helpCmd.AddCommand(action_help_newPaneCmd)
}
