package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_help_session_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List captured threads, newest first",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_help_session_listCmd).Standalone()

	ai_help_sessionCmd.AddCommand(ai_help_session_listCmd)
}
