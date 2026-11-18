package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_session_help_resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Pick a captured thread and resume it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_session_help_resumeCmd).Standalone()

	ai_session_helpCmd.AddCommand(ai_session_help_resumeCmd)
}
