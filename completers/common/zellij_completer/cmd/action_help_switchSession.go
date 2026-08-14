package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_switchSessionCmd = &cobra.Command{
	Use:   "switch-session",
	Short: "Switch to a different session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_switchSessionCmd).Standalone()

	action_helpCmd.AddCommand(action_help_switchSessionCmd)
}
