package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ai_session_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Capture messages an agent hands over, one JSON object per line on stdin, as capture does its transcript's. Messages already captured are skipped",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ai_session_pushCmd).Standalone()

	help_ai_sessionCmd.AddCommand(help_ai_session_pushCmd)
}
