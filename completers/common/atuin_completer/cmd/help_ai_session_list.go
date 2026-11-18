package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ai_session_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List captured threads, newest first",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ai_session_listCmd).Standalone()

	help_ai_sessionCmd.AddCommand(help_ai_session_listCmd)
}
