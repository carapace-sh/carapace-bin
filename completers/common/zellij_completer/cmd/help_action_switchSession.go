package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_switchSessionCmd = &cobra.Command{
	Use:   "switch-session",
	Short: "Switch to a different session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_switchSessionCmd).Standalone()

	help_actionCmd.AddCommand(help_action_switchSessionCmd)
}
