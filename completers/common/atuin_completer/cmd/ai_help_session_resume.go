package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_help_session_resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Pick a captured thread and resume it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_help_session_resumeCmd).Standalone()

	ai_help_sessionCmd.AddCommand(ai_help_session_resumeCmd)
}
