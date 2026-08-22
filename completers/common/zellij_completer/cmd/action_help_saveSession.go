package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_saveSessionCmd = &cobra.Command{
	Use:   "save-session",
	Short: "Save the current session state to disk immediately",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_saveSessionCmd).Standalone()

	action_helpCmd.AddCommand(action_help_saveSessionCmd)
}
