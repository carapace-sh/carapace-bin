package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_togglePanePinnedCmd = &cobra.Command{
	Use:   "toggle-pane-pinned",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_togglePanePinnedCmd).Standalone()

	action_helpCmd.AddCommand(action_help_togglePanePinnedCmd)
}
