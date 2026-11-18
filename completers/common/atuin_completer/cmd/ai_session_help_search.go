package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_session_help_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Full-text search across captured threads, most relevant first.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_session_help_searchCmd).Standalone()

	ai_session_helpCmd.AddCommand(ai_session_help_searchCmd)
}
