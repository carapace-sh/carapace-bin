package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ai_session_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Full-text search across captured threads, most relevant first.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ai_session_searchCmd).Standalone()

	help_ai_sessionCmd.AddCommand(help_ai_session_searchCmd)
}
