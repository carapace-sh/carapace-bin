package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ai_session_resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Pick a captured thread and resume it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ai_session_resumeCmd).Standalone()

	help_ai_sessionCmd.AddCommand(help_ai_session_resumeCmd)
}
