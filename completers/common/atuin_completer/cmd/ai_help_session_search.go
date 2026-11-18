package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_help_session_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Full-text search across captured threads, most relevant first.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_help_session_searchCmd).Standalone()

	ai_help_sessionCmd.AddCommand(ai_help_session_searchCmd)
}
