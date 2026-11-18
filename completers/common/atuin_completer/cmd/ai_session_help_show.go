package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_session_help_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a thread and its messages. Accepts an id, a harness id or `latest`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_session_help_showCmd).Standalone()

	ai_session_helpCmd.AddCommand(ai_session_help_showCmd)
}
