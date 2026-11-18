package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_help_session_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Capture messages an agent hands over, one JSON object per line on stdin, as capture does its transcript's. Messages already captured are skipped",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_help_session_pushCmd).Standalone()

	ai_help_sessionCmd.AddCommand(ai_help_session_pushCmd)
}
