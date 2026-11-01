package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_saveSessionCmd = &cobra.Command{
	Use:   "save-session",
	Short: "Save the current session state to disk immediately",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_saveSessionCmd).Standalone()

	help_actionCmd.AddCommand(help_action_saveSessionCmd)
}
