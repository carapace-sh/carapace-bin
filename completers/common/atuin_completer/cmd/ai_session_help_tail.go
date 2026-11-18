package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_session_help_tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Follow threads and messages as they are recorded (until interrupted)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_session_help_tailCmd).Standalone()

	ai_session_helpCmd.AddCommand(ai_session_help_tailCmd)
}
